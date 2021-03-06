// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package terraform

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/mattermost/mattermost-load-test-ng/deployment"
	"github.com/mattermost/mattermost-load-test-ng/deployment/terraform/assets"
	"github.com/mattermost/mattermost-load-test-ng/deployment/terraform/ssh"

	"github.com/mattermost/mattermost-server/v5/mlog"
	"github.com/mattermost/mattermost-server/v5/model"
)

const cmdExecTimeoutMinutes = 30

const (
	latestReleaseURL           = "https://latest.mattermost.com/mattermost-enterprise-linux"
	defaultLoadTestDownloadURL = "https://github.com/mattermost/mattermost-load-test-ng/releases/download/v1.2.0/mattermost-load-test-ng-v1.2.0-linux-amd64.tar.gz"
	filePrefix                 = "file://"
	minSupportedVersion        = 0.12
	maxSupportedVersion        = 0.13
)

// A global mutex used to make t.init() safe for concurrent use.
// This is needed to prevent a data race caused by the "terraform init"
// command which can modify common files in the .terraform directory.
// Making this a global variable to avoid exporting more methods and
// having the user of this package deal with this special case.
var initMut sync.Mutex

// Terraform manages all operations related to interacting with
// an AWS environment using Terraform.
type Terraform struct {
	id          string
	config      *deployment.Config
	output      *Output
	dir         string
	initialized bool
}

// New returns a new Terraform instance.
func New(id string, cfg *deployment.Config) *Terraform {
	return &Terraform{
		id:     id,
		config: cfg,
	}
}

// Create creates a new load test environment.
func (t *Terraform) Create(initData bool) error {
	if err := t.preFlightCheck(); err != nil {
		return err
	}

	if err := validateLicense(t.config.MattermostLicenseFile); err != nil {
		return fmt.Errorf("license validation failed: %w", err)
	}

	extAgent, err := ssh.NewAgent()
	if err != nil {
		return err
	}

	var uploadBinary bool
	var binaryPath string
	if strings.HasPrefix(t.config.MattermostDownloadURL, filePrefix) {
		binaryPath = strings.TrimPrefix(t.config.MattermostDownloadURL, filePrefix)
		info, err := os.Stat(binaryPath)
		if err != nil {
			return err
		}

		// We make sure the file is executable by both the owner and group.
		if info.Mode()&0110 != 0110 {
			return fmt.Errorf("file %s has to be an executable", binaryPath)
		}

		if !info.Mode().IsRegular() {
			return fmt.Errorf("binary path %s has to be a regular file", binaryPath)
		}

		t.config.MattermostDownloadURL = latestReleaseURL
		uploadBinary = true
	}

	loadTestDownloadURL := t.config.LoadTestDownloadURL
	if strings.HasPrefix(t.config.LoadTestDownloadURL, filePrefix) {
		loadTestDownloadURL = defaultLoadTestDownloadURL
	}

	err = t.runCommand(nil, "apply",
		"-var", fmt.Sprintf("cluster_name=%s", t.config.ClusterName),
		"-var", fmt.Sprintf("cluster_vpc_id=%s", t.config.ClusterVpcID),
		"-var", fmt.Sprintf("cluster_subnet_id=%s", t.config.ClusterSubnetID),
		"-var", fmt.Sprintf("app_instance_count=%d", t.config.AppInstanceCount),
		"-var", fmt.Sprintf("app_instance_type=%s", t.config.AppInstanceType),
		"-var", fmt.Sprintf("agent_instance_count=%d", t.config.AgentInstanceCount),
		"-var", fmt.Sprintf("agent_instance_type=%s", t.config.AgentInstanceType),
		"-var", fmt.Sprintf("es_instance_count=%d", t.config.ElasticSearchSettings.InstanceCount),
		"-var", fmt.Sprintf("es_instance_type=%s", t.config.ElasticSearchSettings.InstanceType),
		"-var", fmt.Sprintf("es_version=%.1f", t.config.ElasticSearchSettings.Version),
		"-var", fmt.Sprintf("es_vpc=%s", t.config.ElasticSearchSettings.VpcID),
		"-var", fmt.Sprintf("es_create_role=%t", t.config.ElasticSearchSettings.CreateRole),
		"-var", fmt.Sprintf("proxy_instance_type=%s", t.config.ProxyInstanceType),
		"-var", fmt.Sprintf("ssh_public_key=%s", t.config.SSHPublicKey),
		"-var", fmt.Sprintf("db_instance_count=%d", t.config.TerraformDBSettings.InstanceCount),
		"-var", fmt.Sprintf("db_instance_engine=%s", t.config.TerraformDBSettings.InstanceEngine),
		"-var", fmt.Sprintf("db_instance_class=%s", t.config.TerraformDBSettings.InstanceType),
		"-var", fmt.Sprintf("db_username=%s", t.config.TerraformDBSettings.UserName),
		"-var", fmt.Sprintf("db_password=%s", t.config.TerraformDBSettings.Password),
		"-var", fmt.Sprintf("mattermost_download_url=%s", t.config.MattermostDownloadURL),
		"-var", fmt.Sprintf("mattermost_license_file=%s", t.config.MattermostLicenseFile),
		"-var", fmt.Sprintf("load_test_download_url=%s", loadTestDownloadURL),
		"-var", fmt.Sprintf("job_server_instance_count=%d", t.config.JobServerSettings.InstanceCount),
		"-var", fmt.Sprintf("job_server_instance_type=%s", t.config.JobServerSettings.InstanceType),
		"-auto-approve",
		"-input=false",
		"-state="+t.getStatePath(),
		t.dir,
	)
	if err != nil {
		return err
	}

	if err := t.loadOutput(); err != nil {
		return err
	}

	if t.output.HasMetrics() {
		// Setting up metrics server.
		if err := t.setupMetrics(extAgent); err != nil {
			return fmt.Errorf("error setting up metrics server: %w", err)
		}
	}

	if t.output.HasAppServers() {
		url := t.output.Instances[0].PublicDNS + ":8065"

		// Updating the config.json for each instance of app server
		t.setupAppServers(extAgent, uploadBinary, binaryPath)
		if t.output.HasProxy() {
			// Updating the nginx config on proxy server
			t.setupProxyServer(extAgent)
			url = t.output.Proxy.PublicDNS
		}

		if err := pingServer("http://" + url); err != nil {
			return fmt.Errorf("error whiling pinging server: %w", err)
		}

		if initData {
			if err := t.createAdminUser(extAgent); err != nil {
				return fmt.Errorf("could not create admin user: %w", err)
			}
		}
	}

	if err := t.setupLoadtestAgents(extAgent, initData); err != nil {
		return fmt.Errorf("error setting up loadtest agents: %w", err)
	}

	mlog.Info("Deployment complete.")
	displayInfo(t.output)
	runcmd := "go run ./cmd/ltctl"
	if strings.HasPrefix(os.Args[0], "ltctl") {
		runcmd = "ltctl"
	}
	fmt.Printf("To start coordinator, you can use %q command.\n", runcmd+" loadtest start")
	return nil
}

func (t *Terraform) setupAppServers(extAgent *ssh.ExtAgent, uploadBinary bool, binaryPath string) {
	for _, val := range t.output.Instances {
		err := t.setupMMServer(extAgent, val.PublicIP, uploadBinary, binaryPath)
		if err != nil {
			mlog.Error("error while setting up app server", mlog.Err(err))
		}
	}

	for _, val := range t.output.JobServers {
		err := t.setupJobServer(extAgent, val.PublicIP, uploadBinary, binaryPath)
		if err != nil {
			mlog.Error("error while setting up job server", mlog.Err(err))
		}
	}
}

func (t *Terraform) setupMMServer(extAgent *ssh.ExtAgent, ip string, uploadBinary bool, binaryPath string) error {
	return t.setupAppServer(extAgent, ip, mattermostServiceFile, uploadBinary, binaryPath, !t.output.HasJobServer())
}

func (t *Terraform) setupJobServer(extAgent *ssh.ExtAgent, ip string, uploadBinary bool, binaryPath string) error {
	return t.setupAppServer(extAgent, ip, jobServerServiceFile, uploadBinary, binaryPath, true)
}

func (t *Terraform) setupAppServer(extAgent *ssh.ExtAgent, ip, serviceFile string, uploadBinary bool, binaryPath string, jobServerEnabled bool) error {
	sshc, err := extAgent.NewClient(ip)
	if err != nil {
		return fmt.Errorf("error in getting ssh connection to %q: %w", ip, err)
	}
	defer func() {
		err := sshc.Close()
		if err != nil {
			mlog.Error("error closing ssh connection", mlog.Err(err))
		}
	}()

	// Upload files
	batch := []uploadInfo{
		{srcData: strings.TrimSpace(serverSysctlConfig), dstPath: "/etc/sysctl.conf"},
		{srcData: strings.TrimSpace(serviceFile), dstPath: "/lib/systemd/system/mattermost.service"},
		{srcData: strings.TrimPrefix(limitsConfig, "\n"), dstPath: "/etc/security/limits.conf"},
	}
	if err := uploadBatch(sshc, batch); err != nil {
		return fmt.Errorf("batch upload failed: %w", err)
	}

	cmd := "sudo systemctl daemon-reload && sudo service mattermost stop"
	if out, err := sshc.RunCommand(cmd); err != nil {
		return fmt.Errorf("error running ssh command %q, ourput: %q: %w", cmd, string(out), err)
	}

	mlog.Info("Updating config", mlog.String("host", ip))
	if err := t.updateAppConfig(ip, sshc, jobServerEnabled); err != nil {
		return fmt.Errorf("error updating config: %w", err)
	}

	// Upload binary if needed.
	if uploadBinary {
		mlog.Info("Uploading binary", mlog.String("host", ip))

		if out, err := sshc.UploadFile(binaryPath, "/opt/mattermost/bin/mattermost", false); err != nil {
			return fmt.Errorf("error uploading file %q, output: %q: %w", binaryPath, string(out), err)
		}
	}

	// Starting mattermost.
	mlog.Info("Applying kernel settings and starting mattermost", mlog.String("host", ip))
	cmd = "sudo sysctl -p && sudo systemctl daemon-reload && sudo service mattermost restart"
	if out, err := sshc.RunCommand(cmd); err != nil {
		return fmt.Errorf("error running ssh command %q, output: %q: %w", cmd, string(out), err)
	}

	return nil
}

func (t *Terraform) setupLoadtestAgents(extAgent *ssh.ExtAgent, initData bool) error {
	if err := t.configureAndRunAgents(extAgent); err != nil {
		return fmt.Errorf("error while setting up an agents: %w", err)
	}

	if !t.output.HasAppServers() {
		return nil
	}

	if err := t.initLoadtest(extAgent, initData); err != nil {
		return err
	}

	return nil
}

func (t *Terraform) setupProxyServer(extAgent *ssh.ExtAgent) {
	ip := t.output.Proxy.PublicDNS

	sshc, err := extAgent.NewClient(ip)
	if err != nil {
		mlog.Error("error in getting ssh connection", mlog.String("ip", ip), mlog.Err(err))
		return
	}
	func() {
		defer func() {
			err := sshc.Close()
			if err != nil {
				mlog.Error("error closing ssh connection", mlog.Err(err))
			}
		}()

		// Upload service file
		mlog.Info("Uploading nginx config", mlog.String("host", ip))

		backends := ""
		for _, addr := range t.output.Instances {
			backends += "server " + addr.PrivateIP + ":8065 max_fails=3;\n"
		}

		batch := []uploadInfo{
			{srcData: strings.TrimSpace(fmt.Sprintf(nginxSiteConfig, backends)), dstPath: "/etc/nginx/sites-available/mattermost"},
			{srcData: strings.TrimSpace(serverSysctlConfig), dstPath: "/etc/sysctl.conf"},
			{srcData: strings.TrimSpace(nginxConfig), dstPath: "/etc/nginx/nginx.conf"},
			{srcData: strings.TrimSpace(limitsConfig), dstPath: "/etc/security/limits.conf"},
		}
		if err := uploadBatch(sshc, batch); err != nil {
			mlog.Error("batch upload failed", mlog.Err(err))
			return
		}

		cmd := "sudo sysctl -p && sudo service nginx reload"
		if out, err := sshc.RunCommand(cmd); err != nil {
			mlog.Error("error running ssh command", mlog.String("output", string(out)), mlog.String("cmd", cmd), mlog.Err(err))
			return
		}

	}()
}

func (t *Terraform) createAdminUser(extAgent *ssh.ExtAgent) error {
	cmd := fmt.Sprintf("/opt/mattermost/bin/mattermost user create --email %s --username %s --password %s --system_admin",
		t.config.AdminEmail,
		t.config.AdminUsername,
		t.config.AdminPassword,
	)
	mlog.Info("Creating admin user:", mlog.String("cmd", cmd))
	sshc, err := extAgent.NewClient(t.output.Instances[0].PublicIP)
	if err != nil {
		return err
	}
	if out, err := sshc.RunCommand(cmd); err != nil {
		if strings.Contains(string(out), "account with that username already exists") {
			return nil
		}
		return fmt.Errorf("error running ssh command: %s, output: %s, error: %w", cmd, out, err)
	}

	return nil
}

func (t *Terraform) updateAppConfig(ip string, sshc *ssh.Client, jobServerEnabled bool) error {
	var clusterDSN, driverName string
	var readerDSN []string

	clusterDSN = t.config.ExternalDBSettings.DataSource
	readerDSN = t.config.ExternalDBSettings.DataSourceReplicas
	driverName = t.config.ExternalDBSettings.DriverName

	if t.output.HasDB() {
		switch t.config.TerraformDBSettings.InstanceEngine {
		case "aurora-postgresql":
			clusterDSN = "postgres://" + t.config.TerraformDBSettings.UserName + ":" + t.config.TerraformDBSettings.Password + "@" + t.output.DBCluster.ClusterEndpoint + "/" + t.config.DBName() + "?sslmode=disable"
			readerDSN = []string{"postgres://" + t.config.TerraformDBSettings.UserName + ":" + t.config.TerraformDBSettings.Password + "@" + t.output.DBCluster.ReaderEndpoint + "/" + t.config.DBName() + "?sslmode=disable"}
			driverName = "postgres"
		case "aurora-mysql":
			clusterDSN = t.config.TerraformDBSettings.UserName + ":" + t.config.TerraformDBSettings.Password + "@tcp(" + t.output.DBCluster.ClusterEndpoint + ")/" + t.config.DBName() + "?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s"
			readerDSN = []string{t.config.TerraformDBSettings.UserName + ":" + t.config.TerraformDBSettings.Password + "@tcp(" + t.output.DBCluster.ReaderEndpoint + ")/" + t.config.DBName() + "?charset=utf8mb4,utf8\u0026readTimeout=30s\u0026writeTimeout=30s"}
			driverName = "mysql"
		}
	}

	cfg := &model.Config{}
	cfg.SetDefaults()
	cfg.ServiceSettings.ListenAddress = model.NewString(":8065")
	cfg.ServiceSettings.LicenseFileLocation = model.NewString("/home/ubuntu/mattermost.mattermost-license")
	cfg.ServiceSettings.SiteURL = model.NewString("http://" + ip + ":8065")
	cfg.ServiceSettings.ReadTimeout = model.NewInt(60)
	cfg.ServiceSettings.WriteTimeout = model.NewInt(60)
	cfg.ServiceSettings.IdleTimeout = model.NewInt(90)
	cfg.ServiceSettings.CollapsedThreads = model.NewString(model.COLLAPSED_THREADS_DEFAULT_OFF)
	cfg.EmailSettings.SMTPServer = model.NewString(t.output.MetricsServer.PrivateIP)
	cfg.EmailSettings.SMTPPort = model.NewString("2500")

	if t.output.HasProxy() && t.output.HasS3Key() && t.output.HasS3Bucket() {
		cfg.FileSettings.DriverName = model.NewString("amazons3")
		cfg.FileSettings.AmazonS3AccessKeyId = model.NewString(t.output.S3Key.Id)
		cfg.FileSettings.AmazonS3SecretAccessKey = model.NewString(t.output.S3Key.Secret)
		cfg.FileSettings.AmazonS3Bucket = model.NewString(t.output.S3Bucket.Id)
		cfg.FileSettings.AmazonS3Region = model.NewString(t.output.S3Bucket.Region)
	}

	cfg.LogSettings.EnableConsole = model.NewBool(true)
	cfg.LogSettings.ConsoleLevel = model.NewString("ERROR")
	cfg.LogSettings.EnableFile = model.NewBool(true)
	cfg.LogSettings.FileLevel = model.NewString("WARN")

	cfg.SqlSettings.DriverName = model.NewString(driverName)
	cfg.SqlSettings.DataSource = model.NewString(clusterDSN)
	cfg.SqlSettings.DataSourceReplicas = readerDSN
	cfg.SqlSettings.MaxIdleConns = model.NewInt(100)
	cfg.SqlSettings.MaxOpenConns = model.NewInt(512)

	cfg.TeamSettings.MaxUsersPerTeam = model.NewInt(50000)
	cfg.TeamSettings.EnableOpenServer = model.NewBool(true)

	cfg.ClusterSettings.GossipPort = model.NewInt(8074)
	cfg.ClusterSettings.StreamingPort = model.NewInt(8075)
	cfg.ClusterSettings.Enable = model.NewBool(true)
	cfg.ClusterSettings.ClusterName = model.NewString(t.config.ClusterName)
	cfg.ClusterSettings.ReadOnlyConfig = model.NewBool(false)

	cfg.MetricsSettings.Enable = model.NewBool(true)

	cfg.PluginSettings.Enable = model.NewBool(true)
	cfg.PluginSettings.EnableUploads = model.NewBool(true)

	cfg.JobSettings.RunJobs = model.NewBool(jobServerEnabled)

	if t.output.HasElasticSearch() {
		cfg.ElasticsearchSettings.ConnectionUrl = model.NewString("https://" + t.output.ElasticSearchServer.Endpoint)
		cfg.ElasticsearchSettings.Username = model.NewString("")
		cfg.ElasticsearchSettings.Password = model.NewString("")
		cfg.ElasticsearchSettings.Sniff = model.NewBool(false)
		cfg.ElasticsearchSettings.EnableIndexing = model.NewBool(true)
		cfg.ElasticsearchSettings.EnableAutocomplete = model.NewBool(true)
		cfg.ElasticsearchSettings.EnableSearching = model.NewBool(true)
	}

	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("error in marshalling config: %w", err)
	}

	if out, err := sshc.Upload(bytes.NewReader(b), "/opt/mattermost/config/config.json", false); err != nil {
		return fmt.Errorf("error uploading config.json: output: %s,  error: %w", out, err)
	}

	return nil
}

func (t *Terraform) preFlightCheck() error {
	if os.Getenv("SSH_AUTH_SOCK") == "" {
		return errors.New("ssh agent not running. Please run eval \"$(ssh-agent -s)\" and then ssh-add")
	}

	if err := checkTerraformVersion(); err != nil {
		return fmt.Errorf("failed when checking terraform version: %w", err)
	}

	if !t.initialized {
		if err := t.init(); err != nil {
			return err
		}
		if err := t.validate(); err != nil {
			return err
		}
	}

	t.initialized = true

	return nil
}

func (t *Terraform) init() error {
	dir, err := ioutil.TempDir("", "terraform")
	if err != nil {
		return err
	}
	t.dir = dir
	assets.RestoreAssets(dir, "outputs.tf")
	assets.RestoreAssets(dir, "variables.tf")
	assets.RestoreAssets(dir, "cluster.tf")
	assets.RestoreAssets(dir, "datasource.yaml")
	assets.RestoreAssets(dir, "dashboard.yaml")
	assets.RestoreAssets(dir, "dashboard_data.json")
	assets.RestoreAssets(dir, "es_dashboard_data.json")

	// We lock to make this call safe for concurrent use
	// since "terraform init" command can write to common files under
	// the .terraform directory.
	initMut.Lock()
	defer initMut.Unlock()
	return t.runCommand(nil, "init", t.dir)
}

func (t *Terraform) validate() error {
	return t.runCommand(nil, "validate", t.dir)
}

func pingServer(addr string) error {
	mlog.Info("Checking server status:", mlog.String("host", addr))
	client := model.NewAPIv4Client(addr)
	client.HttpClient.Timeout = 10 * time.Second
	timeout := time.After(30 * time.Second)

	for {
		select {
		case <-timeout:
			return errors.New("timeout after 30 seconds, server is not responding")
		case <-time.After(3 * time.Second):
			_, resp := client.GetPingWithServerStatus()
			if resp.Error != nil {
				mlog.Debug("got error", mlog.Err(resp.Error))
				mlog.Info("Waiting for the server...")
				continue
			}
			mlog.Info("Server status is OK")
			return nil
		}
	}
}
