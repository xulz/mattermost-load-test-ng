// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/cluster.tf (10.463kB)
// assets/dashboard.yaml (228B)
// assets/dashboard_data.json (26.716kB)
// assets/datasource.yaml (296B)
// assets/outputs.tf (327B)
// assets/variables.tf (779B)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clusterTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\xe3\x36\x12\xfe\xee\x5f\x31\xa7\x5d\xa0\xed\x61\x25\xd9\x49\x9a\x4d\x5a\xa4\x45\xef\x0d\x38\xa0\xbd\x16\xb8\x02\xf7\x61\x11\x08\x34\x39\x96\x08\x53\x24\x8f\xa4\x9c\x78\xf7\x72\xbf\xfd\x40\x4a\xb2\x65\x59\x72\xbc\xde\x4d\x77\x7b\xbb\xc9\x17\x63\x38\x33\x1c\x3e\x33\xf3\xf0\x45\xda\xa8\x15\x67\x68\x20\x22\x77\x36\x82\x37\x13\x00\x6d\xd4\x82\x0b\x84\x1b\x88\xca\x32\x16\x8a\x30\x87\xd6\x45\x13\x00\x83\x39\x57\x12\xfc\x48\x65\x63\x24\xd6\xc5\x67\x5e\xbe\x42\x63\xfd\xc0\x0d\x44\xff\xfd\x0e\xce\x92\x8b\x97\xd1\xe4\x61\x32\x31\x68\x55\x65\x28\x06\xdf\xd9\x12\xd7\x99\x26\xdc\x44\x10\x2d\x71\x5d\x4f\xe5\x65\x92\x94\x08\xc1\xe7\xf3\x37\x2b\x62\x12\x2a\x2a\xeb\xd0\x04\xf9\x43\xbc\xc4\x75\x30\xf2\x71\x55\x73\xc1\xa9\xf7\x03\x37\xe0\x23\xfc\xd2\xab\x5b\x5b\x64\xdb\x91\xaf\xf6\xe7\xe5\xd2\x3a\x22\x29\x46\x10\x11\xad\x33\x8b\x66\x85\xa6\x9e\xde\x91\xdc\xc2\x4d\xf8\x09\xf0\x0f\x1f\xc7\x48\x14\x44\xeb\xf8\xf9\x1b\xaa\x2a\xe9\x12\x2e\x19\xde\x3f\xf8\x80\x1e\x26\x13\x00\xaa\xa4\x44\xea\xfc\xf2\x6b\x3f\xcf\xe0\xd7\x02\x81\xe1\x82\x54\xc2\x41\x65\xd1\x84\x15\x2e\x94\x01\x55\x19\xf8\xe1\xa7\xbf\x07\x35\xb7\xd6\x61\x3a\x6b\x8b\x28\x08\xbc\x66\x40\x76\x5e\x49\x57\xd5\xb2\x42\x59\x07\x37\x60\x51\x2c\x92\x66\x91\x5c\xb7\x33\x93\x92\xc3\xf6\xef\x06\x22\x52\xf2\x78\xba\xa0\x67\x53\xc6\x66\x8c\x5c\x4c\x2f\x5f\x5e\x4d\xe7\x11\x3c\x83\xd9\x55\x32\xbd\x80\x1f\x7f\xfd\xe7\x04\xa0\x85\x23\x6b\x02\xf0\x8b\xf5\xb8\xec\xc8\x77\x53\x13\xbc\x77\x53\x98\x2c\x71\x9d\x70\x16\x56\x5f\x49\xd7\x89\x61\xcf\x5b\x50\xf0\x25\xa2\x69\x66\x91\x56\x86\xbb\x75\x96\x1b\x55\xe9\x8c\x33\x0f\xfe\xab\xb0\xd0\xe8\xf9\x1b\x3f\xc1\xae\x86\xf7\x94\x70\xf6\x10\xbd\x38\xac\x93\xe5\xca\x5a\x5e\xab\x4e\x00\x6e\x3d\x38\x6c\x2d\x49\xc9\x29\x44\x46\x29\x97\xcd\x85\xa2\xcb\x8c\xe1\x8a\xfb\x3a\xa8\xf3\xb4\x50\x26\x43\x42\x8b\x26\xea\x3d\xbd\xa0\x44\x95\x74\x28\x5d\x63\x02\xb0\x52\xa2\x2a\x31\xb3\xfc\xb5\xc7\x4e\x28\xb5\xac\xf4\x97\x7b\xa6\xc9\x8a\x88\x0a\x5f\x40\xd4\x51\x8f\x5e\x80\xac\x84\xf8\x6a\xd7\x51\x93\x84\x63\x1d\x79\xf5\x1d\x47\x0f\x6d\x31\x84\x3e\xf6\x4d\xe8\x5b\xd9\x37\x47\xbb\xcc\xa6\x15\x3a\xf9\x29\x89\x73\x68\x4a\x65\x5d\x26\x38\x45\x69\x31\xf3\x06\x41\x9b\xa1\x75\x5c\x12\xd7\x74\x73\x5a\xa8\x12\xd3\xba\x22\xd3\xad\x5d\xc7\x45\xdc\xb8\x88\x06\x03\x31\x58\x2a\x87\x31\xde\x23\x6d\xe3\xe1\x52\x70\x89\x9b\xc4\x03\x44\x77\x85\x67\x9b\x57\xf0\x07\x88\x17\x90\xae\x88\x49\x05\x9f\xa7\x54\xa8\x8a\xa5\x6d\x1d\xa5\x73\xa5\x5c\xbc\xe0\x92\xdb\x02\x19\xdc\x7e\x0b\x4c\x01\xd2\x42\xc1\x17\xff\x22\xdc\x71\x99\x87\x06\x0b\x46\x31\x97\xdc\x25\x49\xf2\xc5\xb7\x60\x05\xa2\x86\x99\xd7\x96\xd8\x94\x91\x9f\x31\x47\x07\x71\x2c\x55\x4c\x0b\xa4\xcb\x98\xa2\x71\x7c\xc1\x29\x71\x08\xf1\xbf\x7f\x86\x18\x0a\xe7\xb4\xfd\x26\x4d\xed\x79\x8c\x55\x7c\x87\xd6\xc5\xb3\x84\x94\xe4\xb5\x92\xe4\xce\x26\x54\x95\x29\xc3\x79\x62\xd4\xbc\xb2\x4e\xa3\xa1\xa8\x3d\x66\x09\x57\xe9\xc5\xec\xaf\x7f\xfb\xf3\xf5\xf5\x5f\x92\x5c\xe7\xf0\x1f\xb0\x15\x53\x40\xb4\xf3\x4c\x06\x84\x31\x88\xb7\x71\x6c\xc6\x42\x3c\x6b\xa8\x34\x23\x0e\x47\xc6\x03\x14\x42\x78\x3d\x6d\x54\x89\xae\xc0\xca\xc6\x52\x31\x8f\xaf\x56\xc6\xa1\xe9\xaf\xf0\x67\xe8\x24\x8a\x71\xeb\x12\x47\x4c\x92\xbf\x86\x9a\xe3\x3a\x85\xc0\xd4\x9d\xf4\x4c\x9f\x55\x46\x3c\x6c\xdd\x38\x62\xe0\xfe\xf5\x62\xc4\x4d\x2f\xd0\x72\xd5\xd1\x83\x54\x69\x97\xd6\x2c\x76\x1b\x6a\xe3\x10\x31\x97\xe8\x0c\xa7\xf6\x34\x72\x6e\x8c\xff\x6f\x18\x39\x72\x67\x89\x20\x26\x0f\x3d\x75\x0c\x0d\x9f\x4a\xaf\x0d\x70\x1b\x8a\xfd\x4c\x9c\x9f\xf9\xea\x89\xf9\xaa\xa7\x6b\xd7\xd6\x61\x49\x9d\x00\x94\x64\x2e\x70\x5c\x73\xc0\x2b\x61\x2c\x74\xa8\xe0\xf3\x85\x92\x8e\x2a\xb9\xe0\xf9\xac\x87\x5a\x0b\x0c\x13\x49\x6e\xc8\x82\x48\x12\xc0\x50\xd6\xa6\x06\x05\x12\x8b\x69\x23\xcf\x2e\x93\xcb\xe4\x2c\x23\x25\xbb\xbc\x48\x18\xce\x7b\x01\x30\xbd\xcc\x21\xe6\x70\x9c\xf6\x76\x61\x8c\x60\xa9\x64\x6c\xd0\xd3\xeb\x63\xcb\x6f\x9c\xc7\x0d\x0f\xf6\xb4\xd1\xf8\xd2\xee\x29\x81\x75\xc4\xb8\x63\x89\x56\x1b\x75\xbf\x3e\x8d\x66\x83\x69\x4d\xb2\x7d\xa6\xdb\xfd\x1b\xe5\xbd\x3d\xae\xeb\xdb\xf9\x49\xeb\x08\xfb\xa7\xd1\xdd\x73\xe6\x90\xdd\xfe\xa9\x13\xbe\x83\x19\x7c\x0f\x33\xf8\x06\xa6\x3e\x64\x6b\x15\xe5\xc4\x61\xb6\xe1\xed\x8c\x30\x66\xd0\xfa\xf5\x3b\x53\xe1\xa9\x3c\x1a\x42\xde\x9e\x3e\x3b\xa4\x3d\xca\xd7\x9f\x1a\xcb\x7e\xb8\x3d\xf9\xf7\xc0\xef\x4f\x77\x0e\x1c\xb3\x94\x39\x97\xf7\xef\x87\xb3\x86\x5c\x99\x32\x60\x87\x8e\xa6\x61\x38\xb5\xdc\xa1\x8d\x6b\x0b\x96\x36\x69\xef\x19\x09\x09\xf1\xc2\xee\x5b\x91\x15\xe1\xc2\x1b\xa6\xdd\xe3\xe5\x98\xef\xad\x4e\x8f\x13\x7b\xa4\x68\x98\xcd\x5a\x82\xeb\x10\x64\x5f\xd4\xbc\x89\x1c\xe0\x9f\xba\x3f\xd9\x7c\xff\xce\xcb\x19\x4a\xbf\x4f\xa3\x19\xe2\xc8\x21\x8e\x65\xf3\x81\x67\x86\x4d\x48\xfb\xee\x6a\x7e\xe9\x2c\xc5\xc7\xd1\xfe\x0c\x17\xf4\x6d\x50\x82\x58\xfb\x58\xe0\x5e\x67\x02\x80\x1e\xd8\x3e\x3d\x8f\x19\xd5\xca\x9e\x60\xb5\x16\xeb\x8c\x97\x25\x32\x4f\xb3\x62\xdd\xb5\x6a\xf8\x95\x54\x4e\x65\x25\x97\xca\x64\xcd\xa3\x51\x56\xe9\xdc\x10\xe6\x3b\x70\x41\x84\xc5\xfd\xed\xab\xb3\xbc\x08\xa2\xed\x02\x9b\xcc\x0c\xa0\x33\x8a\xae\x87\x93\x11\x47\xe6\xc4\x62\xe7\x58\x3d\x62\x50\xeb\x97\x24\x48\x36\x14\xb5\x83\x42\x2b\xdd\xea\x69\x62\xed\x9d\x32\x6c\x57\xaf\x95\x4e\x00\xec\x92\xeb\x6c\xc1\x25\x11\x99\x95\x44\xdb\x42\xb9\x0e\x3e\x03\x18\x6e\x06\x87\xd2\x72\x20\x1f\xf5\x8f\x16\xe7\x9e\xfa\xee\xe0\xab\x61\x27\xb7\x87\xaf\x16\xc3\xbb\x21\x9b\x87\xad\xf0\xf6\xe0\x39\xa4\x7d\x55\xcc\x48\x8e\xd2\xbd\xed\x6b\x9c\xb7\x39\xfa\x3d\xee\x37\xbb\xd6\x8d\x3d\xae\xf9\x68\xdf\xe7\xf3\xda\xae\xbf\x9a\x6c\x4e\xc8\x53\xf0\xd3\xa4\xea\x13\x3c\x91\x7c\xda\xe7\x82\xf6\x7d\xc8\x95\x7a\xf7\x4d\x28\x3c\x02\x85\xce\x7c\xe4\x49\x68\x6b\xb9\x1d\xdb\x79\x01\x0a\x5f\x0e\x62\xef\x2a\x96\xf9\x1f\xc7\x06\xb6\xc6\xa6\xec\xfa\x3c\x74\x9f\xd9\x2d\xe4\xfa\x5d\xbf\x4e\x5b\x87\xd2\x0f\x3d\xe5\xb7\x0e\xe2\xda\xc1\x24\xbc\x7b\x52\xc3\x75\xfb\xee\xf9\x83\xd6\xd0\x2a\x41\x50\x0a\x69\x6b\x59\xab\xdd\x74\x60\x60\x82\x68\x12\x78\x20\x0f\xf7\x8a\xa6\x71\x8c\x2a\x33\x9f\x81\x10\xd5\xd9\x59\x4d\x4b\xaa\x15\x75\x84\xda\x28\xa7\xa8\x12\x4d\xfc\x8e\xea\x1a\x08\xca\x99\xa9\xeb\xbc\x6e\xe8\x69\x12\xfe\xd3\x69\x74\xdb\x5c\xc8\x0e\xcd\x78\x35\xbd\xfc\x7a\x60\xce\x8d\xf8\xfd\xcf\x1a\x9c\xbf\xec\xcd\xd9\x11\x6e\x67\xec\xce\xf7\x0c\x7e\x22\xeb\x39\x82\x41\xeb\x0c\xa7\x0e\x94\x14\xeb\xe0\x15\x7e\xd9\x54\x34\x34\xd7\xde\xef\x1b\x93\x3f\x55\x0e\x0a\x22\xd9\x1a\xea\xfe\x74\x64\xe9\xbb\xac\xf9\x88\x65\xe1\x8e\xbb\x42\x55\x0e\x4a\x22\x2b\x22\xc4\x1a\xac\x2d\x62\xaf\xc1\xa5\x53\xe0\x0a\x6c\x1c\x26\xef\x0c\x74\x0d\xdf\xf5\x6c\x3a\xdd\x03\xbb\x37\xd4\x05\xbc\x0f\xfa\x6e\x71\x1f\xe0\xef\xee\xeb\x5d\x1b\x1c\x8e\x17\xc1\x7e\x50\xad\xac\x97\xfe\x78\x76\x5c\xf6\x8f\x6a\xcb\xe6\xd3\xcc\x3b\x74\x67\xdc\x78\x38\xba\x49\x6b\xfd\x23\x7b\xf5\xf1\x7c\x5e\x4d\x5f\x5e\x8c\xe4\x73\x33\x34\x90\xcf\x8a\xbd\x6d\x3e\xdb\x8f\x5d\xc7\x16\xda\x89\x81\xbd\x7d\xa1\x9d\x10\xd8\x3e\xdd\xf4\x86\x3e\x14\x62\xa7\x04\xf6\xee\x88\xfd\xc6\x6d\xf9\x58\x5f\xb2\x79\xa7\x1f\xc7\x6f\xa2\xfd\x5d\xf2\x08\x84\xcf\xcf\xa7\x97\x23\x08\x6f\x86\x9e\x00\xe1\x23\x22\xfb\xfa\xe2\x7c\x7f\xdf\xed\x0d\x3d\x41\x64\x47\x90\xe4\xf6\x02\x74\x0c\x3f\x86\xab\xcf\xa3\xe7\x97\x1f\x5b\xf2\x0b\xfa\xbf\xab\xa3\xcc\xc1\x53\xc5\xc5\x74\x6f\x73\xed\x08\x87\x4f\x15\xfe\x4e\xd7\x62\xda\x5c\xa4\x8f\xaa\x98\x0f\xb8\x91\x7f\x74\x3b\x79\xfb\x89\xf5\x51\xda\x68\x14\xdf\x8e\x3b\x3e\xc4\x99\xf8\x7a\x7a\x3d\x84\xe3\x46\xfc\x34\xb3\x9e\xef\xd7\xef\x8e\xf8\xf4\x59\x3f\xae\x7a\xa9\xbf\x15\x1d\x4b\x6a\x41\xfb\x71\x52\xfb\xc5\xab\x3d\x19\x97\x5d\x0d\xc1\x74\xf5\x6e\x69\xf9\xb8\xe8\xf3\x53\x24\xb6\xff\x05\x00\x00\xff\xff\x25\x99\xf8\x63\xdf\x28\x00\x00")

func clusterTfBytes() ([]byte, error) {
	return bindataRead(
		_clusterTf,
		"cluster.tf",
	)
}

func clusterTf() (*asset, error) {
	bytes, err := clusterTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cluster.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdd, 0x5d, 0x68, 0xc3, 0xaa, 0x14, 0xa7, 0xeb, 0xc5, 0x5c, 0xf7, 0x6, 0x28, 0x9f, 0x7b, 0x22, 0xb1, 0x92, 0x0, 0xa1, 0x16, 0x8e, 0xfe, 0xb4, 0xda, 0xf5, 0x14, 0x55, 0xf3, 0xa1, 0xba, 0x3d}}
	return a, nil
}

var _dashboardYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8c\xb1\x4e\x03\x41\x0c\x44\xfb\xfd\x8a\xe9\x52\xa1\x0b\xa2\xdb\x9a\x26\x35\x0a\xbd\x0f\xfb\x12\x4b\xce\x7a\x65\x6f\x0e\xf1\xf7\x68\xaf\x40\xa2\x9d\xf7\xe6\x51\xd7\x4f\x89\x54\x6f\x15\xaf\xa5\xf4\xf0\x5d\x59\x22\x6b\x79\x41\xa3\x87\x54\x9c\x92\x1e\xdd\xe4\x54\x00\x8f\xdb\x85\xa7\x07\x6c\x6e\x2c\xf1\x8f\x8e\x9f\x2e\x15\x9b\x9a\x14\x80\x35\x69\x35\x79\x17\x93\x71\xc4\x37\xb2\x9c\x40\x58\xc7\x24\x15\x23\x9e\x73\x78\x76\xa6\x21\x97\x36\x24\x76\xb2\x0f\xf9\xf2\xc6\x59\xf1\x76\x3e\x17\x80\xcc\xfc\xfb\xaa\xd7\xc3\xc9\xbf\x8f\xf7\x19\xcd\x5a\x00\xa0\xd3\xb8\x57\x2c\x3b\xc5\x62\xba\x2e\xb7\xa0\x8d\x1a\x2d\x4c\x79\x5f\x9d\x82\xf3\x37\x00\x00\xff\xff\x94\x42\xcd\x38\xe4\x00\x00\x00")

func dashboardYamlBytes() ([]byte, error) {
	return bindataRead(
		_dashboardYaml,
		"dashboard.yaml",
	)
}

func dashboardYaml() (*asset, error) {
	bytes, err := dashboardYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dashboard.yaml", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdd, 0x3e, 0x13, 0x14, 0xd7, 0xd7, 0xbd, 0xed, 0xe7, 0xca, 0xfc, 0xe7, 0x3a, 0x3a, 0x0, 0xc1, 0x74, 0x2d, 0xef, 0xf0, 0x9d, 0x47, 0xe0, 0x99, 0xba, 0x23, 0x47, 0x12, 0x37, 0x12, 0x52, 0xd9}}
	return a, nil
}

var _dashboard_dataJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9d\xdb\x6e\xdb\x3a\xd6\x80\xef\xfb\x14\x04\xd1\xff\x47\x52\x24\xd9\x96\x13\xb7\x69\x80\xc1\x20\x6d\xd1\x4e\x81\x66\xa6\xd3\xb4\x7b\x2e\xda\x40\xa0\xa5\x65\x99\x08\x45\xaa\x24\x95\xc3\x0e\x32\xcf\x3e\x20\x75\xa2\x25\xa5\x76\x8c\xb8\x76\x5a\xe6\x22\xb1\x16\x29\x8a\x5c\x5c\x87\x4f\x14\x23\xdf\x3c\x41\x08\x13\xce\x85\x26\x9a\x0a\xae\xf0\x11\x32\x22\x84\x30\xa3\x4a\xe3\x23\xf4\xd5\x1e\xa1\x52\x6a\x4b\xc6\x39\x65\xfa\x3d\xc7\x47\x28\xd8\x69\xa4\x31\xd1\x44\x89\x5c\x46\x80\x8f\x10\xde\xdd\x45\xef\x24\x99\x10\x4e\xd0\xee\x2e\x76\xaa\x01\x27\x63\x66\xaa\x68\x99\x83\x23\x9f\xd2\xb8\x47\x4a\x23\xc1\x5f\x0b\x26\xa4\x69\x53\x26\x63\xb2\x35\xd8\x41\xc3\x20\xd8\x41\xc3\xd1\x68\x07\x05\xdb\x6e\xd3\x9c\xa4\xf6\xda\xc7\xcd\x70\xd0\xff\xa3\x63\x06\x52\x2b\xb7\x9e\xbe\xce\x6c\xbd\x98\xa8\xe9\x58\x10\x19\xe3\xb2\xec\xd6\xfe\x3d\x7b\x82\xd0\xad\xa9\x8e\x21\xa6\xba\xd5\x5b\x9c\x70\xd0\xef\x63\x7c\x84\x86\xa3\x83\x61\x21\x91\x24\x9b\x7e\x16\x82\x69\x9a\x55\x3a\xc1\xd4\x56\x29\x3e\x6a\x90\xb6\x37\xa6\x70\x74\xf8\xfc\x60\x38\x18\x0c\x07\x87\x87\xfb\xb6\x94\x51\x7e\x6e\xb4\xfe\xf5\xcc\x1e\x66\x84\x03\x53\xb5\xde\x2b\xad\x63\xc2\x28\x51\x56\x13\x76\x8a\x6e\xab\xe1\xe0\x31\xb1\x92\x09\x61\xaa\x56\x9c\x1d\xd9\x07\xe0\x89\x9e\x9a\x6b\x0e\x66\xe4\xd0\x57\xdd\x99\x3a\x9e\x33\x56\x97\x4c\x28\x63\xee\x3c\x5b\xc1\x3b\x49\x62\x0a\xdc\x58\x47\xd3\x74\x22\x69\xfc\x51\x34\xf6\x53\x4c\x2a\x3e\x42\x87\x8e\xe6\x2f\x4d\x5b\x43\x47\x70\xe5\xb6\x81\x10\xbe\x36\xc7\xd5\x6c\xd4\x6d\x4f\x69\x1c\x03\x3f\x05\x49\x7b\x3a\x6f\x35\x1d\xbc\xac\x8f\x19\x24\xc0\xe3\xd9\x7e\x90\x8b\xa4\x7d\x1e\x42\x38\xca\xa5\x2c\x86\xd1\x2e\x49\xc9\x55\x9f\x94\xf2\x1e\xa9\x9a\x8a\xcb\xae\xe1\x6a\xa1\x09\xeb\xa9\x7d\x41\x58\xde\x0c\xa2\x33\x52\x46\xb9\x2d\x75\x5b\xb3\xc2\x4b\x1a\x17\x93\x59\x4b\xcd\x3c\x7d\x14\x94\xeb\x13\x61\x3d\xc7\x0a\x6a\x33\xc7\x22\x9b\xf5\xe7\x7a\xa2\x3f\xd4\x06\xd7\xb9\x7a\x06\x32\x02\xae\x49\x02\x1d\x2d\x67\xe6\x4a\x66\xde\x73\x55\x19\x76\x23\xef\x4e\x8a\x04\x1e\x83\x04\xeb\xb6\x13\x26\x74\xd3\x2f\x65\x67\xf1\x5f\x17\x20\x25\x8d\xa1\xb1\xfc\xa2\x30\x23\x11\xf4\x19\xae\xd2\x24\x3a\xef\x5c\x45\x69\xc8\x32\x88\x3f\x50\xde\xed\xb0\x26\x32\x01\xad\x9c\x08\xe6\xc6\x30\xe3\xdc\x57\x99\xed\x9e\xca\xd3\x2d\xca\x23\x09\x44\xc1\x56\x4a\xb4\x06\x99\x0a\xa5\xc3\x78\x1c\x2a\x2d\x24\x84\x9a\xa6\x10\x46\x22\xe7\xfa\x26\x05\x3d\x15\xf1\xdf\xfe\xfb\x0d\x3f\xd5\x22\x33\x55\xac\xfc\x1b\xbe\xfd\x3a\x4a\xcf\xb6\xb7\xd1\xf8\x1a\x6d\x15\x95\xdc\xc8\x54\x5b\xe5\x5b\x21\x53\x62\x0c\x0e\xdf\x94\x6d\xdd\xde\xce\xd6\x93\x30\xb1\xd1\x05\x1f\xe3\x5a\x7c\x5b\x7e\x6a\xf4\xa4\xa7\x12\xd4\x54\xb0\xb8\xa5\x3f\xd3\xd5\xb7\x52\xa4\x6d\x37\x36\xf2\x4f\x90\x94\x06\xd1\x3a\xe1\x74\x4a\x27\xba\x7b\x86\xb6\x71\x0f\x7f\x16\x19\x0a\x06\xe8\xcd\x2b\xf4\x3d\xb7\x53\x67\xc6\x68\x47\xdd\xcc\xa9\xae\x83\xdf\x8d\xeb\x17\x44\x42\xdc\xf5\x0c\x25\xa4\x6e\x79\xbd\x75\x8a\xb0\x0a\xcb\x94\xc7\xf4\x82\xc6\x39\x61\xb8\x63\xa1\x55\x1d\x1b\x73\x9b\x0e\x5c\x91\x2b\xda\x32\xf5\x71\x1e\x9d\x17\xf3\xef\x8e\xcb\x78\x71\xe9\x2e\x66\xe8\x3d\xd9\xa3\x55\xbb\xdf\xbb\x6b\x2f\xee\x71\xa2\x6b\x72\x05\x3f\x30\xbb\x49\x6d\x02\x6a\x6a\x34\x31\x6b\x25\x64\x0c\xac\xd3\x09\x53\x20\x92\x57\x44\xc1\x6c\xca\xad\x03\x55\xa7\x7a\x11\xa9\x3a\x62\x67\x30\x8d\x71\xed\x3c\x92\x7e\x76\x9c\xe0\xba\x3b\xe9\x84\xd1\xa4\x2f\x46\x5b\xf9\x07\xb8\xa8\x3b\x3d\x93\xf0\x4b\x15\xfc\x02\xb9\xf6\xe5\xbc\x5c\x3b\x23\x58\x3a\xd9\xbe\xf0\xc9\xd6\x27\xdb\x07\x4d\xb6\x24\xa3\x6e\x9e\x9d\x12\x1e\x33\x90\x75\xa2\x35\xc5\x3d\x99\xb6\xac\xd6\x4a\xb5\x94\x6b\x90\x17\x84\xbd\x25\x91\xb6\x37\x0e\xfb\x73\x32\x71\xd9\xcc\xc6\xa7\x62\x5d\xa4\xe2\xe3\x8f\xef\x1f\x38\x17\x0f\x7f\xfb\x5c\x9c\xaa\xf5\x27\x38\x9f\x88\xd7\x97\x88\x31\x13\x24\xd6\xa0\xf4\x6e\x29\xaa\xab\xf5\xac\x01\x14\x72\x29\x6d\x70\x99\x6d\x74\x1d\xf7\xcb\x87\x4b\xa4\xf0\xdf\x25\x83\x0f\x5d\xe9\x79\x2b\xe2\x76\x92\x7a\x24\x38\x87\x48\x43\xfc\x93\x32\xfb\xe8\x17\xca\xec\x7d\x59\x3d\x33\xbf\xec\xd4\xde\x50\xae\x34\xe1\x11\xd8\x94\xae\x40\x5e\x80\x34\xb9\x3c\x48\xcf\x5a\xd9\xbb\x09\x77\x96\x07\x8a\x71\xce\x49\xf0\xc1\x9c\x04\x5f\x5d\xbc\x9d\xe1\x53\xd0\x92\x46\x36\xfe\xf7\x75\xfa\x4e\x1c\x98\x8d\x91\x1a\x4c\xb6\x1d\xcd\x0d\xe4\xf3\x18\x68\x31\x6d\xad\x5c\x5d\x9f\xbb\x43\x5f\x4e\x4f\xaf\xe6\xe9\x69\x9d\x2c\xf5\xcf\x3c\x1d\x83\x44\x62\x82\x4e\x40\x29\x92\x80\x42\x19\x48\x74\x0a\x91\xe0\xf1\x1c\x9a\x4a\xd5\x27\x50\x82\xe5\xe5\x9a\x6e\x37\xc2\xf9\x85\x8f\xe2\xe7\xfe\x1c\x83\x5f\xcf\xd0\x2c\xf2\xc8\x85\x3c\x72\xad\x1f\xb9\x3a\xcb\x26\x2f\x97\x60\x2e\x47\x3f\x1e\xba\x3c\x74\xad\x0a\xba\xa6\x5a\x67\xa1\x84\xef\x39\x28\xad\x1e\x1d\x7d\xd9\xde\x5b\x47\x57\x6b\x81\xb0\x7b\x6a\x6f\x63\x68\x6c\x51\xbd\x6d\x34\x94\x7d\x2a\xf5\xee\x51\xcc\xa3\xd8\x7d\xba\xea\x51\xec\xf7\x41\xb1\xf6\xea\x57\xf0\x7c\x09\x14\x7b\xee\x49\xcc\x93\xd8\x83\x92\x58\x3b\x15\x5f\xc2\x58\x09\x1b\xfd\x7f\x08\x11\x9b\x48\x5f\xed\xae\xaf\x90\xc0\x96\xd4\xda\x3a\xa8\xeb\x71\x10\x54\xb3\xac\xf5\xba\x72\x66\xf4\x06\x2e\x68\x04\x0a\x6d\xfd\x07\xc6\xa7\x56\xbb\x55\xa1\xb9\xc2\xb6\x27\x2c\x4f\x58\x9e\xb0\x3c\x61\x21\x74\xd7\x62\x57\xf0\x62\x09\xc4\x0a\x3c\x62\x79\xc4\x5a\x15\x62\xc5\xe3\x30\x25\x4a\x83\x0c\xa3\x26\x93\x3d\x26\xd0\xfa\xc1\x00\x7e\x0e\x6e\xdd\x5f\x83\x1b\xb3\xd4\x75\x7f\xdd\x3d\x2e\x64\xb3\xff\xb7\xa2\x05\x3a\xb1\x83\x44\x6f\x88\x26\x63\x93\x6a\x3d\xa7\x79\x4e\xf3\x9c\xf6\xdb\x70\xda\xdc\xbd\xdc\xed\x95\xb0\xe1\xc1\x32\x98\xf6\xbb\x70\x5a\xe0\x39\x6d\x53\xf6\x78\xab\x3c\x2d\x1e\x9e\xd9\x4d\xdc\x15\x6d\x6c\xa3\x3f\xd0\x62\xdb\xc3\xfb\x4e\xde\x30\xbe\x73\xc7\xba\x22\xa0\x9b\x52\xa5\x45\x22\x49\x1a\x7e\xcf\x09\xd7\x94\xc1\xd6\x60\xef\xe5\x8e\xd5\xa1\x24\xba\x5f\x7f\x45\x1e\xee\x51\xe0\x0e\x83\xd5\x3f\xcc\x7c\x39\xc8\x76\x97\x55\x64\xd1\xf3\xc7\x08\x78\x27\x40\xb8\xdd\xb4\x5f\x3e\xde\x44\x9f\x5d\xd6\xf1\x34\xb7\x12\x9a\xbb\x83\xe4\x3c\xc4\x79\x88\xdb\xec\xc5\xb6\xe1\x68\x09\x8a\x1b\x79\x88\xf3\x10\xf7\xa0\x10\x97\x88\x30\x85\x54\x69\xa2\x55\x48\x18\x13\x51\x38\xbe\xd6\xa0\xee\x58\x1b\xfa\x23\x18\x0c\x0f\xec\xaf\x55\x23\xc4\xb1\xe9\xcb\x42\x10\x71\xc7\x08\x56\xba\xc8\xe6\x5e\xd3\x4e\x4e\x48\x79\xae\x60\x55\xba\x2b\xdf\x99\xd2\xf2\xb2\xfb\x2b\xf5\xd4\x74\x75\xb5\x9a\x9d\x4b\x68\xf7\xd0\xec\x14\x48\xf6\x38\x14\xfb\x0f\x20\xd9\x6a\xf5\xfa\x7a\xb3\xc9\x37\x15\xf2\x1a\x7d\x51\x24\x01\x73\xab\x81\x4e\x5e\xf9\xc7\xcd\xeb\x5b\xc6\x3c\x79\xe5\xf1\xd7\xe3\xef\x46\xe1\x6f\x7b\x0d\x73\x7f\x7f\x09\xfa\x0d\xf6\x3d\xfe\x7a\xfc\x7d\x50\xfc\xa5\xed\x15\xb4\x4c\x8a\x08\x94\x0a\xa3\x2c\x0f\x95\xdd\x99\xbe\xc0\x7f\x07\x3c\x43\xc1\x60\xb0\xbe\x95\xc9\xc5\xb1\x76\x9d\x90\xf0\xfa\xe3\x17\xf4\x45\x53\x46\xff\xb2\x2f\xc9\x43\x9f\x88\x06\xb4\xf5\x7f\x1e\x14\xd6\x07\x0a\x3f\x9c\x0e\xe4\xd1\x01\x79\x74\x58\x3f\x3a\x74\x56\xce\x96\x62\x87\x03\x8f\x0e\x1e\x1d\x1e\x14\x1d\x12\x11\x26\x42\x8a\x5c\x9b\x29\xd9\xe0\xdd\x68\x33\xfd\x5c\xe1\xa6\xb3\x45\xf4\xb1\xde\xbd\x65\x8b\x68\x62\xa3\x9f\x30\x36\x5b\xc8\xde\x09\xf4\xa9\x3d\x10\xcf\x4f\x7e\xbf\xd8\x2f\x0e\x4c\x4f\xca\x66\x8d\xbf\x1a\xaf\x33\xa3\x0e\x06\x85\x0b\x60\x15\x4d\x21\x25\x7f\x82\x54\x85\xa9\x0f\x8b\xd7\x72\x2b\x7d\xcd\xca\xd7\x80\xcb\xf3\xa2\xa6\x26\x49\x33\xf7\xce\x2e\x84\x52\x77\x8e\x64\x37\x03\x39\xc1\xf5\x65\x35\xa4\x19\x23\x9a\xf2\x64\x91\xf7\xa7\x13\xc6\xfe\x34\x26\xd8\xb5\xd7\x86\x2c\x66\xa6\x52\xc3\x95\x9d\xc8\xe0\xc5\x70\x6f\x3f\xd8\x0b\xf6\xf7\x82\xc1\xe8\xe8\x70\xf0\xfc\xc5\xec\xac\x5e\x94\xad\x7e\x75\x84\xa8\xf7\x34\xa7\xc2\x59\x9f\x25\x2d\x88\x85\xa6\x22\x4c\x28\xa7\x65\x14\x29\xec\x2a\x2c\x1c\xac\x77\x73\x4e\xb5\xb4\xee\x06\x10\xca\x23\x96\xc7\x70\xcc\xfa\xc0\xa7\x76\xaa\x22\x59\xb8\x4d\xa5\x39\xd3\xb4\xeb\xde\xd5\x6b\xe0\xbb\x27\x34\x0c\xd3\x84\x57\x84\xf0\xf7\x1c\xe4\xf5\x62\xbd\x6f\xec\x2b\x98\x91\x26\x70\xd5\xda\x5e\x80\xd5\x39\xcd\xbe\x48\x76\x7a\xcd\xa3\xbe\xf8\xda\x8d\xa3\x9a\x24\xd6\x2c\xd4\xbf\xcb\xfe\xb4\x8c\xa3\xb2\xce\xb3\x96\xec\xae\xea\x65\xb0\x2d\x46\xe7\x14\xe4\x0a\x3e\x17\x4d\xf5\xb2\xe4\x43\x18\xea\x31\x63\x0b\x59\xe6\xd3\x30\x24\x8c\xad\xcc\x1a\xed\xd0\x43\x09\x2a\x67\x7a\x4b\x8b\xec\x7c\x2b\x18\xec\x2c\xba\xc9\xec\xe9\x4d\x18\x4a\xc2\x13\x08\xd5\xad\x6a\xbf\x71\x74\xfb\x5e\x36\xdd\x32\xd0\xfe\xe0\x3b\xcf\x9c\x67\xde\x83\x7a\x87\x55\xd7\xc2\x5e\xf3\xfe\x59\xea\x98\xe3\x24\x7b\xcf\xaa\xd7\xbb\x7e\xc3\x5b\x7b\xcf\xfe\xbe\xfd\x0d\xef\x3d\x7b\x58\xcf\x99\x71\xc4\x79\x7e\x33\x5b\x79\x0d\x5e\xa3\x80\x15\x37\x5b\x7d\x4f\x1c\xe7\xba\x54\xc7\x89\x7e\xba\xe3\xf4\xbc\x29\xbf\xfb\x3a\xfc\x35\x79\x4c\xf5\x8a\xfe\x9f\xe0\x30\x4b\xa8\x61\xae\xa7\x94\x5f\x38\xf0\x0b\x38\x8a\xfd\xdb\x7c\xb5\x8a\xe5\xf7\x9a\x95\x26\xc5\xfd\x10\xe6\xe2\x72\x37\x18\x55\x3b\x66\xb1\x16\xa5\x10\xcf\x9c\x97\xd1\xe8\xdc\x2e\x05\x94\x67\x97\x5a\x0c\xab\x7b\x42\x77\x66\xf1\xc8\xb9\x11\xaa\x90\xd0\x1e\xec\xbb\x07\x41\xb3\x4b\x17\x8f\x9c\xcf\x81\x7b\xb0\x3f\x70\x4b\x9c\x1b\x98\xa1\xf3\x39\x28\xbf\x4d\xe6\xac\x1a\x83\x31\x86\xae\xc5\xdd\x7d\x15\xb7\xe1\xe7\x6e\xc3\xee\x55\x86\x07\xee\x81\xb3\x4c\xf3\x22\x76\xfb\x5b\xf5\x65\x46\x7d\x7f\x09\xbb\xbc\x81\xc7\x52\x5c\xaa\x12\x90\x9c\x67\xf6\xb5\x41\xa3\x8f\x20\xed\x4d\x04\x8f\x00\x9d\x08\x4e\xb5\x90\x86\x70\x6d\xfd\xdc\xae\x9e\xe1\x41\xf1\x13\x04\x85\xf4\xa2\x86\xec\x83\x27\xb7\x4f\xfe\x17\x00\x00\xff\xff\x02\x19\xf3\x53\x5c\x68\x00\x00")

func dashboard_dataJsonBytes() ([]byte, error) {
	return bindataRead(
		_dashboard_dataJson,
		"dashboard_data.json",
	)
}

func dashboard_dataJson() (*asset, error) {
	bytes, err := dashboard_dataJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dashboard_data.json", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9f, 0x7c, 0x79, 0x8b, 0x7b, 0x20, 0xbc, 0xa5, 0x2b, 0xfc, 0xd5, 0x36, 0x5f, 0x32, 0x44, 0x17, 0x13, 0x40, 0xca, 0xe5, 0x37, 0xf1, 0x1b, 0x57, 0x5c, 0x1b, 0x48, 0x1, 0xc2, 0x51, 0xf0, 0x5a}}
	return a, nil
}

var _datasourceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xb1\x6e\xeb\x30\x0c\x45\x77\x7d\xc5\x45\x8c\x87\x37\x05\x45\x87\x2e\x9a\xb3\xe4\x07\xba\x33\xd6\xb5\xcd\x42\x96\x0c\x91\x72\x9b\xbf\x2f\x9c\x00\x6d\x37\x82\x87\x38\x38\x1c\x30\xd6\x32\xe9\x8c\x49\x33\xb1\xb3\x99\xd6\x12\x64\xd3\xf7\xe7\x18\xf1\x1a\xc2\x80\xac\xe6\xa8\x13\x92\xb8\x58\xed\x6d\xa4\xc1\x2b\xb4\x18\x9b\xbf\xf4\x2d\x89\x13\x89\x1b\x4b\xd2\x32\x87\x01\x9f\x8b\xf8\x7f\x83\xec\xa2\x59\x6e\x99\xd0\x02\x5f\xf8\x10\xdc\xc4\x18\xfe\x98\x62\x38\xa3\xc8\xca\x88\x5c\x25\x39\xcd\xcf\x4f\x12\x00\xbf\x6f\x8c\xd8\x5a\x5d\xe9\x0b\xbb\x05\x40\xc6\x91\x66\x8f\xe5\xd7\x3d\x00\xb5\xcd\xd7\x74\x74\x02\xbd\xe5\x88\x7f\xc7\x91\xda\x85\x93\xf4\xec\x11\xde\xfa\x61\xda\x7f\x1f\x02\x98\xd4\x8f\xac\x1f\xfa\x61\xb5\x5c\xc4\x25\x06\x00\x70\x5d\x79\x2d\xce\xb6\x4b\x8e\x38\xbd\xd9\xe9\x3b\x00\x00\xff\xff\x50\x28\x34\xbf\x28\x01\x00\x00")

func datasourceYamlBytes() ([]byte, error) {
	return bindataRead(
		_datasourceYaml,
		"datasource.yaml",
	)
}

func datasourceYaml() (*asset, error) {
	bytes, err := datasourceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "datasource.yaml", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xff, 0x82, 0x88, 0xc0, 0xa4, 0xae, 0x3, 0xd5, 0x18, 0x5f, 0xf7, 0x9a, 0xce, 0x18, 0xcf, 0x8c, 0xeb, 0x4c, 0x24, 0x3f, 0xba, 0x8b, 0xe6, 0x78, 0x8f, 0xb9, 0x5e, 0x3f, 0xa9, 0x91, 0xc7, 0x9d}}
	return a, nil
}

var _outputsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x10\x5c\xb9\xe8\x0d\x5c\x79\x04\x0f\x10\xa6\xcd\x20\x85\x98\x84\x99\x49\x55\x4a\xef\x2e\xc4\x06\xac\x48\xdd\x05\xfe\xfb\x2f\x7f\x52\xd1\x5c\x14\xec\x18\x45\x31\x0e\x24\x16\x66\x03\x30\x61\x28\x04\x27\xb0\x87\x19\xef\xe2\x5a\xda\x61\xce\x4e\x88\x27\xe2\xee\xb8\x58\xb3\x18\xd3\x04\xbe\x3f\x87\x22\x4a\xfc\x53\xc0\x5e\xdc\xf0\xce\x3b\xdf\xb7\xe7\xd6\x80\x57\x8a\xfa\xe7\xff\x90\xd0\x2b\x89\xba\x0a\x7f\x6f\xb8\x91\xf2\x38\xc8\xa5\x0e\xdc\x17\xad\xe8\x7a\xcc\x56\x93\x39\x3d\x9e\xfb\xf5\x8a\x7c\x94\x5f\x01\x00\x00\xff\xff\x6f\xd3\xe1\x13\x47\x01\x00\x00")

func outputsTfBytes() ([]byte, error) {
	return bindataRead(
		_outputsTf,
		"outputs.tf",
	)
}

func outputsTf() (*asset, error) {
	bytes, err := outputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "outputs.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0xd7, 0xa7, 0xff, 0x49, 0x7d, 0xf7, 0xf2, 0xf0, 0xad, 0x3d, 0x63, 0x4d, 0x52, 0x40, 0x3f, 0x63, 0xb2, 0x30, 0x39, 0xfe, 0xb0, 0xbb, 0x6b, 0x82, 0xfe, 0xd3, 0xf7, 0xce, 0xb2, 0x30, 0x67}}
	return a, nil
}

var _variablesTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x4d\x6e\x83\x30\x10\x85\xf7\x9c\x62\xe4\x75\x6b\x45\xa9\xd2\x9f\x05\x27\xa9\x2a\xcb\xd8\x13\x6a\xc5\xd8\xae\xc7\x26\xa5\x15\x77\xaf\x30\x55\xd3\x08\x14\x58\xce\xf7\xcd\x63\xf4\xa0\x97\xd1\xc8\xc6\x22\x30\x65\x33\x25\x8c\xc2\xc9\x0e\x19\x7c\x57\x63\x55\x5d\xa0\x0c\x41\x18\x47\x49\x3a\x85\x42\xf9\xec\xd2\x6d\x25\x0d\x61\x25\xa4\x45\x97\x36\x63\xae\xa5\xd5\xa0\x10\xfd\xe7\xb0\xe1\xe8\x66\xeb\x4d\xff\x0d\x74\xad\x71\x1b\x21\x56\x12\xad\x19\xf3\xae\xe8\x31\x92\xf1\x6e\x32\x00\xa6\x93\xa0\x86\x4e\x86\x0a\x40\xe3\x51\x66\x9b\xa0\x2e\x08\x80\xc9\x1c\x7d\x94\xf7\xdd\x40\x1f\x96\x41\x79\x6a\x60\x07\xfe\xc4\xcb\x48\xcc\x9c\xef\xf9\xee\x81\xef\xd9\xd5\x4e\xf0\x94\xda\x88\x65\xb1\x06\xf6\xc2\x1f\xf9\xf3\x64\x8c\x8b\xb3\x32\x61\x5c\xfd\x96\xba\x11\x41\x12\x9d\x7d\xd4\x0b\x46\xf4\x2e\x42\x6e\xac\x51\xe2\x84\xc3\x02\x77\x32\x25\x8c\x9d\xa7\x24\xb4\x3f\x3b\xeb\xa5\x16\x39\xda\x5b\x9e\x35\x0a\x1d\xa1\x38\x1a\xbb\x3c\xa5\x04\x24\xdc\x8a\x8b\xde\x27\xd1\x58\xaf\x4e\x42\x63\x6f\x14\xce\x25\x5f\x7a\x7d\x2d\x1d\xcd\xed\x02\xf4\xde\xe6\x6e\xfe\x2d\xa6\x8e\xda\xf0\xdb\xe1\x1f\x21\xf3\x35\x91\xc3\xae\x8c\xc7\xbb\x0a\xe0\xad\x1a\x7f\x02\x00\x00\xff\xff\x13\x26\x6f\x52\x0b\x03\x00\x00")

func variablesTfBytes() ([]byte, error) {
	return bindataRead(
		_variablesTf,
		"variables.tf",
	)
}

func variablesTf() (*asset, error) {
	bytes, err := variablesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "variables.tf", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdc, 0x86, 0xe, 0x94, 0xf, 0xf5, 0xd6, 0xbc, 0x4d, 0x91, 0xd2, 0xcd, 0x71, 0x7a, 0x6a, 0x8f, 0x4c, 0x1f, 0xeb, 0xe7, 0x50, 0x1b, 0xf, 0xb4, 0x63, 0xa9, 0xb4, 0x10, 0x3c, 0x21, 0x8e, 0x65}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cluster.tf":          clusterTf,
	"dashboard.yaml":      dashboardYaml,
	"dashboard_data.json": dashboard_dataJson,
	"datasource.yaml":     datasourceYaml,
	"outputs.tf":          outputsTf,
	"variables.tf":        variablesTf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"cluster.tf":          &bintree{clusterTf, map[string]*bintree{}},
	"dashboard.yaml":      &bintree{dashboardYaml, map[string]*bintree{}},
	"dashboard_data.json": &bintree{dashboard_dataJson, map[string]*bintree{}},
	"datasource.yaml":     &bintree{datasourceYaml, map[string]*bintree{}},
	"outputs.tf":          &bintree{outputsTf, map[string]*bintree{}},
	"variables.tf":        &bintree{variablesTf, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
