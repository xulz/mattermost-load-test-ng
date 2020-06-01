// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/cluster.tf (12.427kB)
// assets/dashboard.yaml (231B)
// assets/dashboard_data.json (30.868kB)
// assets/datasource.yaml (296B)
// assets/outputs.tf (453B)
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

var _clusterTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x73\x1b\xb7\x15\x7e\xe7\xaf\x38\x5d\x7b\x26\x97\xf1\xee\x92\x92\x2c\xcb\x4a\x94\x8c\xd2\x38\x9d\xce\xd8\xb5\xa7\x49\x9b\x07\x8f\x66\x07\x04\x0e\x49\x54\x58\x00\x05\xb0\x94\x68\x57\xfd\xed\x1d\x60\xb1\xcb\xe5\xf2\x22\x5a\xb6\xec\xa4\x8e\xfd\x42\x01\xe7\x86\x73\xf9\x0e\x2e\xab\x8d\x9a\x73\x86\x06\x12\x72\x65\x13\x78\x3b\x00\xd0\x46\x4d\xb8\x40\x38\x83\xa4\x2c\x53\xa1\x08\x73\x68\x5d\x32\x00\x30\x38\xe5\x4a\x82\x9f\xa9\x6c\x8a\xc4\xba\xf4\xc0\x8f\xcf\xd1\x58\x3f\x71\x06\xc9\x7f\xbf\x83\x83\xec\xe8\x49\x32\xb8\x19\x0c\x0c\x5a\x55\x19\x8a\x41\x76\x71\x89\x8b\x42\x13\x6e\x12\x48\x2e\x71\x51\xab\xf2\x63\x92\x94\x08\x41\xe6\xc3\xb7\x73\x62\x32\x2a\x2a\xeb\xd0\x84\xf1\x9b\xf4\x12\x17\x81\xc9\xdb\x55\x8d\x05\xa7\x5e\x0e\x9c\x81\xb7\xf0\x4b\x4f\x6e\xed\xac\x58\xce\x7c\xb5\xae\x97\x4b\xeb\x88\xa4\x98\x40\x42\xb4\x2e\x2c\x9a\x39\x9a\x5a\xbd\x23\x53\x0b\x67\xe1\x27\xc0\xdf\xbc\x1d\x5b\xac\x20\x5a\xa7\x0f\xdf\x52\x55\x49\x97\x71\xc9\xf0\xfa\xc6\x1b\x74\x33\x18\x00\x50\x25\x25\x52\xe7\x97\x5f\xcb\x79\x00\xbf\xcc\x10\x18\x4e\x48\x25\x1c\x54\x16\x4d\x58\xe1\x44\x19\x50\x95\x81\xf3\x17\x7f\x0d\x64\x6e\xa1\x83\x3a\x6b\x67\x49\x18\xf0\x94\xc1\xb3\xe3\x4a\xba\xaa\x1e\x9b\x29\xeb\xe0\x0c\x2c\x8a\x49\x16\x17\xc9\x75\xa3\x99\x94\x1c\x96\xff\xce\x20\x21\x25\x4f\x87\x13\x7a\x30\x64\x6c\xc4\xc8\xd1\xf0\xf8\xc9\xc9\x70\x9c\xc0\x03\x18\x9d\x64\xc3\x23\x78\xfe\xcb\xcf\x03\x80\xc6\x1d\x45\x34\xc0\x2f\xd6\xfb\x65\x65\x7c\x35\x34\x41\x7a\x37\x84\xd9\x25\x2e\x32\xce\xc2\xea\x2b\xe9\x3a\x36\xac\x49\x0b\x04\x3e\x45\x34\x2d\x2c\xd2\xca\x70\xb7\x28\xa6\x46\x55\xba\xe0\xcc\x3b\xff\x75\x58\x68\xf2\xf0\xad\x57\xb0\x4a\xe1\x25\x65\x9c\xdd\x24\x8f\x76\xd3\x14\x53\x65\x2d\xaf\x49\x07\x00\x17\xde\x39\x6c\x21\x49\xc9\x29\x24\x46\x29\x57\x8c\x85\xa2\x97\x05\xc3\x39\xf7\x79\x50\xc7\x69\xa2\x4c\x81\x84\xce\xa2\xd5\x6b\x74\x81\x88\x2a\xe9\x50\xba\xc8\x02\x30\x57\xa2\x2a\xb1\xb0\xfc\x8d\xf7\x9d\x50\xea\xb2\xd2\x5f\xae\xb1\x66\x73\x22\x2a\x7c\x04\x49\x87\x3c\x79\x04\xb2\x12\xe2\xab\x55\x41\x31\x08\xfb\x0a\xf2\xe4\x2b\x82\x6e\x9a\x64\x08\x75\xec\x8b\xd0\x97\xb2\x2f\x8e\x66\x99\xb1\x14\x3a\xf1\x29\x89\x73\x68\x4a\x65\x5d\x21\x38\x45\x69\xb1\xf0\x0c\x81\x9a\xa1\x75\x5c\x12\x17\xab\x39\x9f\xa9\x12\xf3\x3a\x23\xf3\x25\x5f\x47\x44\x1a\x45\x24\x1b\x0d\x31\x58\x2a\x87\x29\x5e\x23\x6d\xec\xe1\x52\x70\x89\x6d\xe0\x01\x92\xab\x99\x47\x9b\xd7\xf0\x27\x48\x27\x90\xcf\x89\xc9\x05\x1f\xe7\x54\xa8\x8a\xe5\x4d\x1e\xe5\x63\xa5\x5c\x3a\xe1\x92\xdb\x19\x32\xb8\xf8\x06\x98\x02\xa4\x33\x05\x5f\xfc\x4a\xb8\xe3\x72\x1a\x0a\x2c\x30\xa5\x5c\x72\x97\x65\xd9\x17\xdf\x80\x15\x88\x1a\x46\x9e\x5a\x62\x4c\x23\xaf\x71\x8a\x0e\xd2\x54\xaa\x94\xce\x90\x5e\xa6\x14\x8d\xe3\x13\x4e\x89\x43\x48\xff\xfd\x12\x52\x98\x39\xa7\xed\x69\x9e\xdb\xc3\x14\xab\xf4\x0a\xad\x4b\x47\x19\x29\xc9\x1b\x25\xc9\x95\xcd\xa8\x2a\x73\x86\xe3\xcc\xa8\x71\x65\x9d\x46\x43\x51\x7b\x9f\x65\x5c\xe5\x47\xa3\x67\x3f\xfd\xf9\xe9\xd3\x1f\xb3\xa9\x9e\xc2\x7f\xc0\x56\x4c\x01\xd1\xce\x23\x19\x10\xc6\x20\x5d\xda\xd1\xce\x05\x7b\x16\x50\x69\x46\x1c\x6e\x99\x0f\xae\x10\xc2\xd3\x69\xa3\x4a\x74\x33\xac\x6c\x2a\x15\xf3\xfe\xd5\xca\x38\x34\xfd\x15\xbe\x84\x4e\xa0\x18\xb7\x2e\x73\xc4\x64\xd3\x37\x50\x63\x5c\x27\x11\x98\xba\x92\x1e\xe9\x8b\xca\x88\x9b\xa5\x18\x47\x0c\x5c\xbf\x99\x6c\x11\xd3\x33\xb4\x9c\x77\xe8\x20\x57\xda\xe5\x35\x8a\x5d\x84\xdc\xd8\x05\xcc\x25\x3a\xc3\xa9\xbd\x1b\x38\x47\xe6\xff\x1b\x44\x4e\xdc\x41\x26\x88\x99\x86\x9a\xda\x07\x86\xef\x0a\xaf\xd1\x71\x2d\xc4\xfe\x01\x9c\x7f\xe0\xd5\x3d\xe3\x55\x8f\xd6\x2e\xac\xc3\x92\x3a\x01\x28\xc9\x58\xe0\x76\xca\x0d\x52\x09\x63\xa1\x42\x05\x1f\x4f\x94\x74\x54\xc9\x09\x9f\x8e\x7a\x5e\x6b\x1c\xc3\x44\x36\x35\x64\x42\x24\x09\xce\x50\xd6\xe6\x06\x05\x12\x8b\x79\x1c\x2f\x8e\xb3\xe3\xec\xa0\x20\x25\x3b\x3e\xca\x18\x8e\x7b\x06\x30\x7d\x39\x85\x94\xc3\xad\xd4\x2b\x6a\xa7\xdc\xcd\xaa\x71\x50\xc9\xe5\xb8\xa2\x97\xe8\x96\x3f\xa2\x01\x36\x6f\xd0\x37\x9f\x1f\x64\xa3\x6c\xd8\x52\x14\xe1\xcf\x42\x70\x59\x5d\xdf\x6a\xd8\x3b\x31\x2d\x1d\xcf\x08\x96\x4a\xa6\x06\xbd\x01\xb7\x85\x27\x2e\x3e\x8d\x38\xdd\xa3\x46\xe3\x4b\xaf\x47\x04\xd6\x11\xe3\x6e\x13\xdc\x18\xbf\x45\x64\x33\x1d\x85\xed\xd9\x55\xb4\x51\xd7\x8b\xbb\xf5\x94\xc0\x5a\x77\x94\x3e\xac\xaf\xfe\xdb\x0a\xf2\x6b\xc0\xde\xe7\xf3\x4a\x6b\x0b\xfb\x5b\xef\xd5\x4d\xf5\x26\xbe\xf5\x2d\x36\x7c\x07\x23\xf8\x1e\x46\x70\x0a\x43\x6f\xb2\xb5\x8a\x72\xe2\xb0\x68\x9b\x54\x41\x18\x33\x68\xfd\xfa\x9d\xa9\xf0\xae\x4d\x23\x98\xbc\xdc\x6a\x77\x3a\xd4\xd6\xe6\xf4\xb9\xb5\x94\x4f\xb7\x01\xf9\x3d\x34\xb3\xfb\xdb\xf4\x6e\xe3\x94\x53\x2e\xaf\x3f\x0c\x00\x6e\x12\x65\xca\xe0\x3b\x74\x34\x0f\xd3\xb9\xe5\x0e\x6d\x5a\x73\xb0\x3c\x86\xbd\xc7\x24\x24\xa4\x13\xbb\xce\x45\xe6\x84\x0b\xcf\x98\x77\xf7\xd2\xdb\x64\x2f\x69\x76\x63\x22\x29\x0b\x9f\x4b\x09\x24\xf6\xb0\xfe\xe1\xf3\xa1\xde\x57\x6e\x41\xc0\x48\xd8\xc2\xd1\x5e\xc0\xb3\x51\x33\xa1\x14\x6d\x00\x86\xa0\xbf\xbd\xfc\x09\xd9\x1d\x51\xa3\xb1\x30\xab\xd5\xbe\x1e\x5e\x64\xde\x8c\xf7\xd4\x6e\x0f\x8b\xd8\x56\xbc\xe2\xe6\xa7\xd7\x1d\xdb\xc9\xe6\xb5\x67\x2d\xad\x47\x52\x2a\x22\xce\x6b\xc3\xe7\x3e\x53\x97\x08\xbd\x27\x18\xef\xdd\x75\xba\x7a\x43\x3d\x4f\x94\xa1\x58\xf8\x43\xb9\x51\x8b\x06\xb8\xb7\x86\xb7\xd0\x4a\x70\xba\x68\xa3\xdc\xfc\xb9\x67\xac\x23\xf9\x7d\x44\xc6\x43\x53\x90\x0e\x67\xf0\xed\xb7\xcf\x5e\xfe\x34\xf0\x36\x25\xff\xac\xef\x0d\x93\x53\x48\x0e\x86\xa3\x83\x74\x34\x4c\x47\x4f\x42\xa1\x24\x3f\x3b\xe2\xb0\x44\xe9\x92\xd3\x16\xac\x1a\xe0\x0f\x55\xf4\x6c\x32\x41\xea\x67\x93\x73\x21\xd4\x55\x5b\x5e\x61\xf2\x3c\xa0\x6f\x87\xb5\x9d\xb1\x87\xa7\xcf\xb9\x75\x3f\xac\xee\x36\xba\xd3\x7f\xc1\x38\xfb\x5c\xd1\x70\x11\x92\x74\x88\x2e\x56\xd4\xfc\x3d\x86\xc1\x5b\x41\x8c\x3c\x25\x57\xf6\xd4\x1e\x9e\x9e\x9e\xc6\xce\xd9\xa4\x5f\x9b\x50\xde\x7b\xb1\x75\x86\x8e\xf1\x68\xd3\xca\x76\x1a\x7f\x3e\x56\xc6\xbd\xa8\x84\xe3\x9a\x18\xf7\x0f\xbd\x02\x5a\x5d\xc2\x1f\x51\xa0\xc3\x97\xe3\x7f\x79\x2f\x6d\x59\xe7\x3e\xb3\xe7\x54\x6c\x24\x78\x55\xed\x62\x6f\x67\x3d\xfb\x56\xff\xed\x8c\xe1\x5d\x9d\x9b\x7f\xdd\xba\x37\xec\x50\x6e\x06\x3e\xdf\x6e\x06\xfd\xaa\x31\xcc\x16\x4d\x19\x74\x36\x8d\xfd\xa1\x78\x29\xbe\x63\x4f\x56\x57\x00\x1b\xaf\x5f\x7a\x72\x86\xd2\x1f\xd4\x7c\x35\xad\x31\x6d\xac\x44\x36\xde\x70\xcf\xdc\x9a\xb4\x2e\xae\xae\xd1\xce\x52\xbc\x1d\xcd\xcf\x70\x43\xbb\x34\x4a\x10\x6b\x6f\x33\xdc\xd3\x0c\x00\xd0\x37\x9b\xfe\x96\x75\x1b\x53\x4d\xec\xa1\x52\x6b\xb1\x28\x78\x59\x22\xf3\x5b\x4f\xb1\xe8\x72\xc5\x3d\x27\xa9\x9c\x2a\x4a\x2e\x95\x29\xe2\xab\x41\x51\xe9\xa9\x21\xcc\x83\xe2\x84\x08\xbb\x01\xdf\x3a\xcb\x4b\x20\x59\x2e\x30\x46\x66\x83\x77\xb6\x7a\xd7\xbb\x93\x11\x47\xc6\xc4\x62\xe7\x5e\x65\x0b\x43\x4d\x5f\x92\x30\xd2\x6e\xdb\x56\xbc\xd0\x8c\x2e\xe9\x34\xb1\xf6\x4a\x19\xb6\x4a\xd7\x8c\x0e\x00\xec\x25\xd7\xc5\x84\x4b\x22\x0a\x2b\x89\xb6\x33\xe5\x3a\xfe\xd9\xe0\xc3\x76\x72\x53\x58\x76\xc4\xa3\xfe\xd1\xf8\xb9\x47\xbe\x3a\xf9\x7a\xb3\x90\x8b\xdd\x77\x4b\x9b\x4f\x08\x6c\x1c\x30\xee\x62\xe7\xd9\xac\x79\x56\x2a\xc8\xd4\x83\xfc\x3b\x3e\xc7\x78\x9e\xbd\x1f\x64\x3e\xda\xbd\xde\xb6\xd7\x15\x6f\xed\x87\x7c\x5f\x59\x95\x57\x83\xcd\x1d\xe2\x14\xe4\xc4\x50\x7d\x86\xa7\xb4\xcf\xfb\xac\xd4\x3c\x10\xb8\x52\xaf\x3e\x0a\x84\x57\x80\x50\x99\xb7\xbc\x09\x2c\x39\x97\x73\x2b\x4f\x00\xe1\xe9\x38\xf5\xa2\x52\x39\xfd\x7a\xdb\xc4\x92\xd9\x94\x5d\x99\xbb\xce\x33\xab\x89\x5c\x3f\xec\x76\xb7\xb9\x3b\x5b\x2c\xd1\x3a\x6d\x04\xa4\xb5\x80\x41\x78\xf8\xa2\x86\xeb\xe6\xe1\xeb\x5c\x6b\x68\x88\x20\x10\x85\xb0\x35\xa8\xd5\x34\x1d\xd8\xa0\x20\x19\x04\x1c\x98\x86\xbb\x96\x58\x38\x46\x95\x85\x8f\x40\xb0\xea\xe0\xa0\x86\x25\xd5\x0c\x75\x06\xb5\x51\x4e\x51\xd5\x1c\x39\x1c\xd5\xb5\x23\x28\x67\xa6\xce\xf3\xba\xa0\x87\x59\xf8\x9f\x0f\x93\x8b\x78\x49\xb5\x4b\xe3\xc9\xf0\xf8\xf1\x06\x9d\xed\xf0\x87\xd7\x1a\x84\x3f\xe9\xe9\xec\x0c\x2e\x35\x76\xf5\x3d\x80\x17\x64\x31\x46\x30\xfe\xc0\xc3\xa9\x03\x25\xc5\x22\x48\x85\x57\x6d\x46\x43\xbc\x57\xfc\x3e\xb2\xfc\x50\x39\x98\x11\xc9\x16\x50\xd7\xa7\x23\x97\xbe\xca\xe2\x57\x0c\x16\xae\xb8\x9b\xa9\xca\x41\x49\x64\x45\x84\x58\x80\xb5\xb3\xd4\x53\x70\xe9\x14\xb8\x19\x46\x81\xd9\x7b\x3b\xba\x76\xdf\xd3\xd1\x70\xb8\xe6\xec\xde\x54\xd7\xe1\x7d\xa7\xaf\x26\xf7\x0e\xfc\xee\x3e\xdf\x34\xc6\xe1\xf6\x24\x58\x37\xaa\x19\xeb\x85\x3f\x1d\xed\x17\xfd\xbd\xca\x32\xbe\xcd\xbf\x47\x75\xa6\x51\xc2\xde\x45\x5a\xd3\xef\x59\xab\xb7\xc7\xf3\x64\xf8\xe4\x68\x4b\x3c\xdb\xa9\x0d\xf1\xac\xd8\xbb\xc6\xb3\xf9\xda\x61\xdf\x44\xbb\xa3\x61\xef\x9e\x68\x77\x30\x6c\x1d\x6e\x7a\x53\x9f\xca\x63\x77\x31\xec\xfd\x3d\xf6\x91\xcb\xf2\xb6\xba\x64\xe3\x4e\x3d\x6e\x3f\x89\xf6\xbb\xe4\x1e\x1e\x3e\x3c\x1c\x1e\x6f\xf1\x70\x3b\x75\x0f\x1e\xde\xc3\xb2\xc7\x47\x87\xeb\x7d\xb7\x37\x75\x0f\x96\xed\x01\x92\xcb\x03\xd0\x3e\xf8\x18\x8e\x3e\xb7\xee\x5f\x9e\x37\xe0\x17\xe8\x7f\x57\x5b\x99\x9d\xbb\x8a\xa3\xe1\x5a\x73\xed\x0c\x6e\xde\x55\xf8\x33\x5d\xe3\xd3\x78\x90\xde\x2b\x63\x3e\x61\x23\xff\xcd\x75\xf2\xe6\x1b\x9b\x5b\x61\x23\x12\xbe\x1b\x76\x7c\x8a\x3d\xf1\xd3\xe1\xd3\x4d\x7e\x6c\x87\xef\x47\xeb\xe1\x7a\xfe\xae\x0c\xdf\x5d\xeb\x47\xce\x97\x07\xf0\x2b\x82\x44\x64\x40\xc0\xa2\x26\x86\x38\xec\xc3\x8c\xa9\xfc\x76\x5c\x81\x36\x38\xf7\x28\x44\x17\x54\x70\x0a\x0c\x35\x4a\x86\x92\x2e\x60\x8c\xee\x0a\x51\x0e\x1e\x84\x7d\x38\xd1\x3a\x72\x12\xc9\x20\x66\x52\x3d\x92\xed\x4c\xd0\xc2\x6b\xaa\xf7\x9b\xa9\x53\x69\xfb\x31\x41\x7d\xab\xb4\xfe\xf8\xde\xd6\x6a\x0c\x95\x5f\x75\xbf\xfe\x3b\x64\x07\x8f\x43\x74\xfa\x30\xb0\x4e\xd1\x47\x03\x58\x87\x85\xb5\x0b\x9a\x25\xc9\x4e\x64\x18\x34\x5f\x72\xae\x5f\xf1\x6c\x66\xad\xdb\xd0\x1e\xa5\x5d\x7f\xea\xb0\x6f\xff\x09\xd4\xb7\xf7\x9f\x57\x9e\xec\xde\xda\xce\xc9\xa6\x8c\x3e\x79\xbf\x0a\xfa\x6d\x75\xba\x26\x06\x9f\x53\x0f\xfa\x5f\x00\x00\x00\xff\xff\x28\x82\x00\xe6\x8b\x30\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xac, 0xd9, 0xc9, 0x94, 0x77, 0xf0, 0xde, 0xdd, 0xf8, 0x14, 0x7f, 0x0, 0x5e, 0xd, 0x5e, 0xfc, 0xc2, 0x5f, 0xdd, 0x30, 0x37, 0xa2, 0xc0, 0x16, 0x71, 0x8, 0xa4, 0xc4, 0x43, 0x67, 0xea, 0x21}}
	return a, nil
}

var _dashboardYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8c\xb1\x6e\x02\x41\x0c\x44\xfb\xfd\x8a\xe9\xa8\xa2\x23\x4a\xb7\x75\x1a\xea\x88\xf4\x3e\xec\x03\x4b\x66\x77\x65\x9b\x8b\xf2\xf7\xd1\x52\x44\xa2\x9d\xf7\xe6\xd1\xd0\x6f\xf1\xd0\xde\x2a\xde\x4b\x19\xde\x77\x65\xf1\xa8\xe5\x0d\x8d\xee\x52\x71\xb8\x4b\xba\x5e\xe2\x50\x80\xee\xd7\x13\x4f\x11\xd8\xba\xb1\xf8\x2b\xce\xdf\x21\x15\x9b\x9a\x14\x80\x35\x68\x35\xf9\x14\x93\x7c\xe6\x37\xb2\x98\x40\x58\x73\x92\x8a\xf4\xc7\x1c\x1e\x83\x29\xe5\xd4\x52\x7c\x27\xfb\x92\x4b\x6f\x1c\x15\x1f\xc7\x63\x01\xc8\xac\xff\x9c\xf5\xfc\x74\xe2\xff\xd3\xc7\x8c\x46\x2d\x00\x30\x28\x6f\x15\xcb\x4e\xbe\x98\xae\xcb\xd5\x69\xa3\x46\x0b\x53\xdc\xd6\x4e\xce\x51\xfe\x02\x00\x00\xff\xff\x4e\xda\x5e\xba\xe7\x00\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x59, 0xd3, 0xf5, 0x60, 0x9d, 0xb1, 0xc3, 0x4f, 0x3f, 0x8, 0x37, 0x26, 0x8f, 0xe5, 0x97, 0x54, 0xbb, 0xa4, 0x6b, 0x3d, 0xee, 0xdf, 0x22, 0x37, 0xc, 0xdd, 0xb7, 0x38, 0xb7, 0xdb, 0x2a, 0x9f}}
	return a, nil
}

var _dashboard_dataJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9d\x6d\x73\xdb\x36\xf2\xc0\xdf\xe7\x53\x60\x30\xff\xff\x8d\xd3\xb1\x5b\x4b\x89\xe3\xd8\x33\x7d\xe1\xa8\x49\x2e\x33\x71\xeb\xf3\x43\xaf\x73\x49\x86\x03\x91\x2b\x09\x63\x12\x60\x00\x50\xb6\x9a\x73\x3f\xfb\x0d\xc0\x27\xf0\x41\x32\xa5\x48\xb6\xd2\xa0\x2f\x6a\x11\x04\x81\xdd\xc5\x62\xf1\xcb\x02\xa2\xbe\x3c\x41\x08\x13\xc6\xb8\x22\x8a\x72\x26\xf1\x31\xd2\x45\x08\xe1\x90\x4a\x85\x8f\xd1\x07\x73\x85\xb2\x52\x73\x67\x98\xd0\x50\xbd\x63\xf8\x18\xf5\x76\xcb\xd2\x80\x28\x22\x79\x22\x7c\xc0\xc7\x08\xef\xed\xa1\xb7\x82\x8c\x08\x23\x68\x6f\x0f\x5b\xd5\x80\x91\x61\xa8\xab\x28\x91\x80\x55\x3e\xa1\x41\x4b\x29\xf5\x39\x1b\xf0\x90\x0b\xdd\xa6\x18\x0f\xc9\xce\xfe\x2e\xea\xf7\x7a\xbb\xa8\x7f\x70\xb0\x8b\x7a\x4f\xed\xa6\x19\x89\x4c\xdf\x27\xa5\x3a\xe8\x1f\xe8\x24\x04\xa1\xa4\x5d\x4f\xcd\x62\x53\x2f\x20\x72\x32\xe4\x44\x04\x38\xbb\x77\x67\xfe\x7e\x7a\x82\xd0\x9d\xae\x8e\x21\xa0\xaa\x26\x2d\x1e\x33\x50\xef\x02\x7c\x8c\xfa\x07\xcf\xfb\x69\x89\x20\xf1\xe4\x92\xf3\x50\xd1\x38\xb7\x09\xa6\xa6\x4a\xfa\x51\x81\x30\xd2\xe8\x9b\x07\x2f\x8f\x8e\x0e\x8f\x7a\x87\x87\x07\x47\x2f\xcd\xdd\x90\xb2\x6b\x6d\xf5\x0f\x9f\xcc\x65\x4c\x18\x84\xb2\xb0\x7b\x6e\x75\x4c\x42\x4a\xa4\xb1\x84\x19\xa2\xbb\x5c\x1d\x3c\x24\xa6\x64\x44\x42\x59\x18\xce\x68\xf6\x1e\xd8\x58\x4d\x74\x9f\xfb\x95\x72\x68\xab\x6e\x0f\x5d\xc8\x49\xa0\x40\xaa\xbd\xac\xa8\xa8\xd6\x62\x8e\xb4\x5c\x08\x33\x40\xd5\x46\x47\x34\x0c\x6d\x17\x31\x05\x6f\x05\x09\x28\x30\xed\x58\xa5\x54\x63\x41\x83\x33\x5e\xba\x5e\xea\x0f\xf8\x18\x1d\x59\x83\x76\xa3\xdb\xea\x5b\x05\xb7\x76\x1b\x08\xe1\x99\xbe\xce\x07\xb2\x68\x7b\x42\x83\x00\xd8\x05\x08\xda\xa2\xb7\x19\xa4\xc3\xe2\x32\x84\x31\xb0\xa0\x2a\x06\x99\x8e\xeb\x8f\x21\x84\xfd\x44\x88\x54\x8b\xfa\x9d\x88\xdc\xb6\x95\x52\xd6\x52\x2a\x27\xfc\xa6\xe9\xf2\x8a\x2b\x12\xb6\xd4\x9e\x92\x30\x29\x75\x68\x28\x1a\x52\x66\xee\x56\x46\x46\x17\xde\xd0\x20\x75\x03\xbb\xd4\x72\xb9\x74\xea\x24\x61\x78\xc6\x29\x53\xa7\xdc\x4c\x43\xec\x73\xc6\xc0\x57\x10\x94\xc3\xcf\xe3\x6a\x84\x28\x5c\xe7\x7d\xd1\x5e\x43\xaa\x18\x84\x0f\x4c\x91\x31\x34\x8c\x1f\xeb\xee\xb4\x3b\x24\xfa\xd9\x83\x6a\x79\x73\xac\x04\xb0\x00\x04\x98\x40\x30\x0a\xb9\x2a\xe5\x92\x66\x70\x7f\x9b\x82\x10\x34\x80\x9a\x62\x32\x26\x3e\xb4\x4d\x05\xa9\x88\x7f\xdd\xe8\x45\x2a\x88\x63\x08\xde\x53\xd6\x14\x58\x11\x31\x06\x25\xad\x98\x68\x47\x45\x3d\x0f\x6e\xe3\x34\x4e\x11\x05\x3b\x11\x51\x0a\x44\xc4\xa5\xf2\x62\xfd\x3f\x33\xac\x5f\x28\x93\x8a\x30\x1f\x7e\xfe\xeb\x23\xfe\x3f\x09\x62\x0a\xe2\x23\xbe\xfb\xd0\x8b\x3e\xd9\x91\x4c\x4f\x16\x2e\x22\xa2\x1d\x0c\x2b\x1a\x81\x97\xea\x58\xad\x42\x99\x02\x31\x25\xe1\x1b\xe2\x2b\x33\xff\x7a\x95\xdb\xa9\x37\xbf\x29\xda\xf9\x52\x74\x7e\x77\x57\x6d\x28\x02\x25\xa8\xaf\xeb\xb4\x0a\x5d\xad\x2c\x60\x64\xc2\x1f\x3e\xa9\x96\x6b\xcb\xe9\x71\x2c\xca\xee\x76\x17\x5b\x49\x26\xd1\xce\x6a\x96\xda\xb8\xa9\x2e\x9b\x6a\xaf\x66\xa3\x57\xf7\xd9\x28\xfb\x54\xfa\xab\x9a\x08\x90\x13\x1e\x06\x35\x3f\xd6\xaa\xbd\x11\x3c\xc2\xc7\x48\xcf\xd5\x4a\xf9\x39\x8c\xb3\x89\x59\x7b\xe0\x62\x42\x47\xaa\xf9\x84\x32\x21\x1c\x9f\x71\xa9\x24\x8a\x41\xa0\x0b\xf0\x39\xb3\x66\xba\x2a\x16\x33\x6b\xa6\x47\xf2\x1c\x24\x0f\x93\x6c\x25\x6b\x06\x33\x22\x20\x68\x86\x33\xc9\x85\xaa\x45\x6a\x13\xc9\xbc\x7c\x15\xa6\x2c\xa0\x53\x1a\x24\x24\xc4\x8d\xf0\x91\xd7\x31\x4b\x6c\x29\xdf\x2d\xb9\xa5\xb5\x38\x34\x4c\xfc\xeb\x74\x72\xda\xca\x6a\xb1\xb3\x80\xa6\xed\xd1\x02\x0b\xb5\xda\xed\x21\xb9\x08\xbd\x2d\x11\x6e\x46\x6e\x61\x41\x4c\x28\x9d\x53\x4e\xb4\x25\xaa\x7e\x47\x86\xa0\x43\x3d\x1e\xf0\x84\xd5\xef\xf1\xf1\x2b\x22\xa1\xe1\xab\xe9\x02\x53\x15\xbb\x58\x61\x1a\xc5\x96\x3e\xf7\x4e\xcc\x4e\xa2\x36\x7a\xd8\xa0\x9c\x8d\xc9\x31\x6b\x8e\x3b\x09\xe9\xb8\xcd\x1d\x4d\xf9\x7b\x98\x16\x42\x57\x10\x2f\x33\x81\xa3\xab\xfc\xba\x95\xae\x2a\x05\xab\xe2\x95\x65\x9f\xf9\x7c\x55\x9b\x6f\xf7\xe1\x55\xad\x7a\xea\x53\xf5\xb8\xb3\x1a\x5c\x59\x1e\xe8\xd8\xea\x1b\x60\xab\x89\x52\xb1\x27\xe0\x73\x02\x52\xc9\x6d\x80\x2c\x94\xf7\x8e\x3a\x61\x96\x91\xdf\x4c\x71\xf9\xe0\xb4\xb5\xa4\xed\xb6\x06\xbb\xba\xda\x6c\xab\xe9\xeb\x3c\xb3\xbb\x03\x30\x07\x60\xcb\x88\xea\x00\xec\xfb\x01\xb0\x7a\x7a\xeb\x68\x05\xfe\x7a\xd1\x01\xbf\x5c\x7a\xcb\x21\x58\x77\x04\xab\xaf\xc4\x37\x30\x94\xdc\x04\xff\x85\x0c\xb1\x65\xf9\xad\x56\xd1\x37\x08\x5f\x2b\x5a\xed\x31\xa0\xeb\xdb\x00\xa8\x5f\x93\x68\x08\x02\xf1\x11\x1a\xe4\x93\x19\xfd\x02\x53\xea\x83\x44\x3b\xff\x86\xe1\x85\xb1\x6e\x7e\x53\xf7\xf0\xd4\x01\x96\x03\x2c\x07\x58\x0e\xb0\x10\x9a\x97\xe1\x5a\x85\xb0\x7a\x8e\xb0\x1c\x61\x6d\x8a\xb0\x82\xa1\x17\x11\xa9\x40\x78\x7e\xb9\x90\x7d\x4b\x9c\xb5\x40\x81\x87\xa1\xad\xe5\x2d\xb8\x35\x89\xae\xe5\x6d\xf7\x6d\x11\x9b\x39\xa1\xa3\x38\x3a\x35\x4a\xa2\x5f\x88\x22\x43\xbd\xd2\x3a\x4c\x73\x98\xe6\x30\xcd\x61\x1a\x9a\x93\x07\xeb\xbd\x5c\x05\xd3\x1c\xa7\x39\x4e\x5b\x2f\xa7\x69\xd2\xa0\xcc\x17\x40\x64\x65\x63\x8d\xc4\xd4\x4b\x59\x21\x89\xd2\xad\x33\x34\x9c\xa1\x9d\x9c\x36\x9e\xa2\x9f\xd0\xbd\x8f\xfa\x3a\x12\xb7\x3d\xbc\x34\x9c\x64\x27\x6b\x6b\x6e\xf6\x00\x7b\x9b\xb6\x19\x36\xc4\x7a\x13\x2a\x15\x1f\x0b\x12\x79\x9f\x13\xc2\x14\x0d\x61\x67\xff\xc7\xa3\xa3\x5d\xd4\xb6\xdf\x59\xc8\x93\xae\xd1\x2d\xc6\xdd\x0d\x61\x85\x6d\xce\xb9\x06\xae\x5b\x30\x3e\x3a\xda\x9b\x8f\xcf\x16\xc4\x6d\x07\xaf\x9d\x02\x61\xe8\xe4\xec\x1d\xca\xf6\x2a\xd1\xa5\x8d\x2e\x0e\xce\x36\x02\x67\x73\xc0\xcc\x31\x99\x63\xb2\xad\x62\xb2\x46\xea\x6c\x25\x28\x3b\x70\x4c\xe6\x98\x6c\xad\x4c\x36\xe6\x5e\x04\x91\x54\x44\x49\x8f\x84\x21\xf7\xbd\xe1\x4c\x81\x7c\xa4\x64\xd9\x89\x96\x60\xc1\x9a\x5f\x72\xd3\x1c\xb9\x37\x9a\x29\xb3\xfb\x34\x43\xe2\x51\x96\x48\x58\xaf\xc5\xd6\x85\x9f\x17\x5a\xc0\xcd\xda\xf3\xde\xec\xd9\x12\xf6\x9c\x00\x89\xb7\xd9\x9c\xff\x04\x12\x6f\xd6\x9a\x83\x6d\xce\x45\x9e\x42\xc4\xc5\x0c\x5d\x49\x1d\x1d\x1d\xd2\x6e\x12\x69\x8d\xa7\x38\xac\x9d\x27\xa7\x4d\x14\x8e\x6b\xf3\xf2\xad\xc9\x35\xf6\x0f\x57\xc0\xda\xde\x33\xc7\xb5\x8e\x6b\xd7\xca\xb5\xb4\xf1\x5d\x49\xc1\x7d\x90\xd2\xf3\xe3\xc4\x93\xe6\xfc\x78\x87\x33\xfc\x3f\xa0\xde\xfe\xfe\xe3\xed\x10\x77\x27\xd7\xc7\x64\x83\xc1\xd9\x15\xba\x52\x34\xa4\x7f\x9a\xaf\xed\xa3\x73\xa2\x00\xed\xfc\xbf\x3b\x3e\xf6\x78\xfb\x92\x0b\x87\x03\x7d\xb7\xec\xe0\x52\x62\xdb\x84\x0e\x8d\x94\xd8\x4a\xec\xf0\xdc\xa1\x83\x43\x87\xb5\xa2\xc3\x98\x7b\x63\x2e\x78\xa2\xf4\x90\x6c\xf1\xa9\xb1\x8a\x9c\x1b\x3c\x1c\xd6\xc5\x1e\x8f\x7b\x06\xac\x8b\x25\xbe\x91\xa3\x5e\x6f\x39\x3a\xaf\x2b\xe2\xf8\xc9\x9d\xeb\x72\xc0\x84\x36\x02\x4c\x95\xf9\x18\x80\xf4\x05\x8d\xb3\x19\x55\x66\xdb\x1e\x25\x93\xf2\xec\xc5\x0a\x34\xd4\xef\x72\x6a\x6b\xee\xeb\x23\xea\x2f\x8a\xe8\xfc\xf6\x08\x41\xc7\x13\x75\xd1\x9a\x6a\x7f\xac\x37\x4b\x34\x90\x48\x17\x3c\x10\x0d\xf5\xff\x46\x34\x54\x9c\x0c\x62\x3c\x00\xcf\x8f\x93\x76\x02\xd8\xd5\x71\xfd\xe7\x8f\x38\x91\xd6\x1b\x10\x16\x9d\xc4\x5a\xb0\x31\xa3\xeb\xb7\xc0\x79\x0b\x15\x5c\x49\x10\xa8\x53\xfe\x64\x29\xe2\x59\x42\x5f\x39\x93\x0a\xa2\xaf\xd3\xb8\xb1\x75\x67\xda\xec\xa2\xd7\xab\x8d\xe9\x45\x83\x10\xd6\xab\xd5\xbb\x20\x84\x2e\x3a\x0d\x36\xa4\xd3\x5f\x1f\x31\x15\x9f\xff\x2b\xf9\x48\xe9\xbf\x94\xdf\x10\xaa\xd6\xab\xe2\x6f\x6a\xd2\xcd\x1f\x7f\xd9\x92\x23\x6b\x83\xb3\x2b\xf4\x9e\x93\x00\xed\xf4\xa2\xfb\xf2\x75\x0e\x28\x5b\x7d\x71\xeb\x28\xcd\xd1\xe4\xd6\xd0\xa4\x4f\x23\x12\xd6\x5d\xf6\x71\x72\x6b\xab\xd1\xe4\x7a\xf7\xe5\x74\x1c\x7d\x1d\xc5\x6a\x36\xe7\xde\x7f\x40\xf0\xaf\xcc\xc9\x6d\x00\x44\xbf\x3a\x5f\xe7\x48\x74\x3d\x24\x9a\xae\xf2\x91\x39\x89\xe2\x9d\x42\x74\xb9\x60\xf7\x0e\xed\xa1\x5a\xed\x37\x02\xa0\x5b\xe5\x57\xc9\x68\x04\x62\x5e\xa6\xab\x56\x79\x40\xfc\x09\x04\xdd\xb2\x62\xeb\x42\xde\x00\x2d\xfa\x6e\xc1\x2a\xd0\xbb\x84\x65\xf1\x42\xe9\xcc\x93\xdd\xc4\xdb\x9e\x63\xfb\x9d\x8f\x36\x39\x04\x6a\xf5\xa2\x4e\x67\x97\x1c\x02\x7d\x27\xa7\x97\x1e\x9c\x75\xea\x89\xb3\xe7\x07\xab\xa0\x4e\xb9\x83\x95\xef\x4c\x54\x53\x80\xdf\xc3\xdb\x58\x1d\xa9\xac\x89\x54\xaa\x89\x09\x06\xea\x86\x8b\x6b\x4f\x09\xc2\x64\x44\xd5\xc2\x13\xce\x6d\xe9\x88\xa7\x3f\xbc\x6c\xdf\x3c\x6b\x9e\x09\x5d\x76\x5b\xed\x8f\xf4\x7c\xc8\xba\x81\x62\x8e\x05\x04\xf8\x40\xa7\x8b\x8f\x78\x77\x32\xc0\x4a\x28\xb5\xa4\x6d\xce\x97\xb1\xcd\xb6\xd0\xcc\xaf\xa9\xa5\xd1\xf9\x1f\x3f\x65\x63\x7b\xdf\x3e\xa2\x83\x9a\x56\x4f\xb6\xa0\x26\x76\x48\xe3\x90\xe6\x41\x91\xa6\x91\xbe\x59\x8d\x69\xba\xfc\x56\x8f\xa3\x97\x15\xe9\x25\x4c\xc6\x94\xfd\x0e\x42\x66\xbb\xc5\x2f\x7e\x7c\xf1\x63\x1f\x7f\x07\x74\x93\xaf\xe8\x52\x11\xe5\x5d\xfa\xb1\x37\x48\x84\x78\x2d\x15\x19\x76\x4a\x19\xac\x25\x0b\x62\xba\x0b\xa9\x9c\x6c\x20\x19\x42\x2b\xdc\x92\x6b\xf9\xfa\x56\x79\x97\x83\xb3\x4b\x1a\x01\x4f\xd4\x42\x78\xd9\x80\xbe\x79\xb7\xcb\xc3\xc8\x3d\xca\x6a\x4e\x5b\x7a\x40\x57\xdd\x1f\x4b\xb3\x43\xd6\xd8\xad\x67\x0b\xf0\xab\x07\xec\xeb\xf4\xc9\x7b\xd8\xf2\xbd\xbe\xcb\xc1\x99\xfd\x22\x31\x47\x85\x6e\xb7\xef\x6f\xc2\x85\x4f\xb2\x66\xf5\x9c\xd3\xd3\x49\xab\xdd\xdb\x4f\x3d\x1c\x4b\x7f\x02\x11\x29\xd7\xe9\x7e\xfa\x2b\x93\x52\xcd\xd2\x69\x11\x10\x71\x9d\xd6\x54\x64\x5c\x8e\xbc\xf5\x62\x9e\xcc\x78\x56\xc9\x5e\x0c\x62\x84\x8b\x6e\x15\x44\x71\x48\x14\x65\xe3\x2e\x3f\x07\x4a\xc2\xf0\x77\xed\x80\x4d\x6f\x2d\xd9\xab\x32\x96\x0a\x6e\x55\xf3\x1f\xfa\xd3\xac\x8d\xc2\x87\x2b\xce\xd0\xf1\x5b\x01\x66\x97\x72\x44\x19\xcd\x8f\xbc\x19\xd7\xf0\xd2\x09\xd2\x7a\x1c\x22\x0f\x8e\x76\x00\xa0\xcc\x0f\x93\x00\x4e\xc2\x36\xbc\x2b\xce\x54\xa6\x21\xd7\x6e\x2a\x4a\x42\x45\x9b\xd3\x33\xff\x5d\xd2\xe6\x03\x25\xc2\x95\x91\x0f\x21\xfc\x39\x01\x31\xeb\x26\x7d\xe9\x21\xbd\x4a\xe9\x18\x6e\x6b\x26\xc6\xf2\x9a\xc6\x57\x22\xbc\x98\x31\xbf\x6d\x33\xb1\x19\x07\x15\x19\x9b\x81\x95\xff\xca\xe5\xc1\xd5\xbb\x0d\xc1\x75\x59\x7b\xe5\x2c\x50\xa6\x9a\x59\x37\x12\x09\x97\x69\x43\x95\x6d\x49\xf3\xb7\xfc\x21\x56\x13\x1c\x0b\x57\x1c\xa5\xeb\x08\x66\xfc\x66\xaf\x77\x90\xbf\x5a\x0a\x2b\x9e\x15\xe2\xca\x73\x31\xf5\xaf\x0d\x86\x66\x4f\x67\x16\xf3\xf2\xe4\x89\x1d\x1c\xf1\x81\xb5\x8c\xe4\x33\xce\x5c\x3c\xb3\x2f\x7a\xe5\xeb\xac\xf0\x81\xf5\xb9\x67\x5f\x3c\xdb\xb7\xef\x58\xab\x43\xdf\xfa\xdc\xcb\x7e\x7b\xf6\x53\xae\x03\x8d\xc0\xb3\xbc\xe2\xde\x5e\xec\x86\x5f\xd8\x0d\xdb\xbd\xf4\x9f\xdb\x17\xd6\x57\x28\x0e\x03\x5b\xde\x5c\x96\x8a\xf9\xfe\xe4\x06\xad\xf1\x50\xf0\x1b\x99\x79\xaf\xb5\xd7\x54\x84\x10\x73\xec\x66\xef\x12\xa4\x42\x67\x20\x4c\xb8\xd6\x48\x77\xca\x19\x55\x5c\xe8\x50\x62\x9e\x4c\xcc\x3f\xe4\xf0\x7e\xfa\x5f\xaf\x97\x96\x4e\x8b\x68\xf6\xec\xc9\xdd\x93\xff\x05\x00\x00\xff\xff\x43\x09\x2b\xee\x94\x78\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa4, 0x48, 0x12, 0x83, 0xcb, 0xbb, 0x98, 0x6, 0xb0, 0xed, 0x7f, 0xa6, 0x41, 0x18, 0xa2, 0xad, 0xed, 0x9, 0x68, 0x77, 0xca, 0x76, 0xa1, 0xce, 0x5e, 0xec, 0x65, 0x1d, 0xc0, 0x8f, 0x94, 0x18}}
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

var _outputsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x3d\x8e\xc3\x20\x10\x85\x7b\x4e\x81\xd0\x56\x5b\xd0\xb8\xde\x6a\xcb\x2d\xf7\x00\x68\x0c\xa3\xc8\xf2\x1f\x62\x06\x27\xc8\xf2\xdd\x23\x61\x53\xd8\x71\x9c\x0e\xe9\xbd\xef\xe3\xcd\x18\xd9\x47\x96\xaa\x19\x88\x61\xb0\x48\x4a\xce\x42\xca\x09\xba\x88\xf2\x47\xaa\xaf\x19\xee\x64\x4a\xaa\xc1\x7b\x43\x18\x26\x0c\xfa\x7b\x51\x62\x11\xa2\x08\x5c\xfd\xdb\x45\x62\x0c\xa7\x82\xe0\xc8\xd8\x35\xd7\xae\x2e\xcf\xbd\x01\x6e\x38\xf0\x87\xff\xbb\x11\x1c\x23\xb1\xc9\xe5\xe3\x86\x1e\x39\x34\x96\xfe\xf3\xc0\x6b\xd1\x56\xdd\x8e\xd9\x6b\x7c\x18\x1f\xe9\x1a\xcf\x95\x53\x98\xaa\x3a\xda\x16\x79\xe5\x5f\x0c\x54\x99\x35\xd7\xa5\x78\xc4\xff\x30\xbd\x61\x1b\xe8\x0d\x58\x8b\x44\xa6\xc5\xa4\xa9\x6a\x31\x65\xfa\x19\x00\x00\xff\xff\x9b\x4f\xd0\x61\xc5\x01\x00\x00")

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
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x13, 0xf6, 0x0, 0x1, 0x5d, 0xc9, 0xe4, 0x88, 0xa3, 0x96, 0x10, 0xb1, 0xe9, 0x86, 0x1e, 0x1d, 0xae, 0xa9, 0xb2, 0x5a, 0x40, 0x88, 0x37, 0x83, 0xf4, 0x9b, 0x98, 0x3a, 0x7c, 0xe3, 0xae, 0x58}}
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

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
