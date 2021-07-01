// Code generated by go-bindata. DO NOT EDIT.
// sources:
// bin/syscall_x86_tester
// +build functionaltests,amd64

package syscall_tester

import (
	"bytes"
	"compress/gzip"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataSyscallx86tester = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x7d\x70\x54\xd7\x75\xff\xbd\xb7\x0f\x78\x36\xcb\xb2\x08\xc5\xc8\x86\xa6\x5b\xbc\x50\x39\x85\x45\x12\xb2\x8d\x6b\xd2\x22\x90\x04\xb4\x12\xc6\x18\xe1\x3a\x1e\xfb\x79\xb5\xfb\x56\xbb\x61\xbf\xba\xfb\x64\x23\x8f\xed\x80\x17\x62\xb6\x6b\xb5\x8b\x43\x1c\x7b\x9c\x74\x88\xf3\x31\x9e\x36\x9e\xfa\x0f\xd7\xa5\x4d\xa7\xc8\x15\x41\x76\xeb\x8c\x69\xd2\x66\x26\x75\x67\x42\x1a\xe3\x2e\x13\xe1\x21\x35\x06\x21\x09\xdf\xce\xb9\xf7\xbe\xdd\xb7\x6b\x29\x6a\x6b\x66\x3a\xd3\xec\xd1\x5c\xbd\xf7\xbb\xf7\x9c\x73\xcf\x39\xf7\xbe\x7d\xe7\xee\xde\xfb\x85\xae\x9e\x6e\x45\x51\x60\x93\x0a\x17\x08\x9d\x2e\x6a\x7a\x3b\x80\xe8\x1d\xa2\xbe\x1d\x3e\x2c\x44\x33\x7e\x0d\x2b\x30\x9f\x63\xa0\x7d\xbf\xa6\x53\x79\x48\x01\xa8\x68\x10\xc5\x05\xe0\xa8\x02\x1c\x3d\xa0\xe9\x54\x96\x02\x58\x2a\xdb\x14\x59\x38\xed\xd7\x74\x2a\x5e\x0d\xa0\x42\xed\xf0\xca\x76\x2f\x80\xa2\xa6\x53\xf9\xe6\x75\x00\x95\x79\xce\x76\x1f\x80\x63\x9a\x4e\xe5\xab\x2e\x80\x8a\x53\xde\xbd\x0e\x70\x9f\xd0\x74\x2a\x6b\x14\xa0\x5d\x01\xb7\x9b\xda\x55\x00\x8d\xeb\x80\xc6\x13\x9a\x4e\xa5\x04\x80\xca\x7c\x54\x7c\x78\x49\x01\x5e\x3a\xa0\xe9\x54\x3a\x01\x74\x3a\xda\x76\x9e\xb5\xc2\x07\x57\x02\x07\x5f\xd4\x74\x2a\x1b\x01\x6c\x74\xb4\xdf\x7d\xd6\x0a\x63\x06\x9a\x2f\xdd\xda\x75\xd6\x0a\x3b\xed\xbb\x08\xe0\xa2\x23\x3e\xeb\xe2\xb1\xfe\x75\xf1\xf0\xda\x78\x2c\x39\xb8\x2f\x90\x4d\x05\xda\x44\x5b\xa3\x8c\xed\xd6\x1d\x7d\xd8\xd8\x13\x5e\xf9\xa3\xdf\xf9\xf4\x97\xc7\xf6\x2c\xf5\xe9\x87\xdf\xfb\xfe\xd9\xdc\xa3\x6d\x9a\xd4\xaf\x48\x1e\x48\x7e\xb5\x3c\xb6\xc0\x42\xd9\xce\x63\xe9\x53\x57\x42\xd6\x6d\xf8\xcf\xc2\x8a\x97\x7f\xff\x67\x23\xb5\x36\xdf\xed\xb8\x6f\x00\xf0\x9b\x35\xd8\x57\x83\x9f\x70\x60\x6a\xfb\x83\x9a\xf6\x8e\x1a\x7c\x53\x0d\xde\x5e\x83\x6f\xad\xc1\x1b\x6b\x70\x80\x6c\x1f\xd1\x74\xf2\x7d\x09\x9a\xb8\x2f\xda\x31\x1b\x2f\x06\xe2\xb1\xfe\x10\xc5\xf0\x36\x18\xdb\xef\x32\xb2\x56\x38\x96\x34\x06\xb3\x66\x18\x91\x54\xda\x4c\x22\x6d\x66\x32\xa9\x0c\x22\xa1\x78\x2a\x6b\x22\x6b\x85\xcd\x4c\x06\x91\x58\xdc\x4c\xa6\x10\xb4\x52\x31\x44\xd2\x99\x58\xd2\x8a\x20\x3b\x94\x0d\x05\xe3\x71\x64\xad\x4c\x28\x91\x86\x61\x90\x62\x23\x6b\x05\x33\x96\x91\x08\xc6\x92\xd8\xda\xb3\x7d\xf3\x16\xa3\x2d\xd0\x5a\xbe\x6b\x81\x61\x0c\x24\x52\x49\xc9\x65\x40\x3e\x5f\x2a\xbf\x8a\x3b\xf1\xa7\x40\xe1\xff\xc5\xf8\x11\x2d\x89\xc5\x16\xd1\xe8\x7d\x5e\xd6\x79\x39\x56\x61\xc9\xf6\xe9\x13\x9a\x3e\x5f\x13\xbe\xcf\x5b\x08\xb8\x47\x34\x7d\x81\x02\x78\xe9\xaa\x02\x8d\x74\x75\x01\x4d\x74\x9d\x07\xac\xa0\xeb\x7c\xc0\x47\xd7\x05\x80\x9f\xae\x3a\xd0\x4c\xd7\xeb\x80\x35\x74\xbd\x7e\xa6\x59\x5b\xa7\x3a\xd5\xa9\x4e\x75\xaa\x53\x9d\xea\x54\xa7\x3a\xd5\xe9\xff\x0b\x7d\xb0\xf8\xd3\x53\xf7\xe4\xc6\xf5\xd2\x52\x05\x38\x30\xfa\xc1\x3a\xa0\x90\x9b\x66\x8c\x1d\x1a\xb1\x54\x76\x3a\x77\x52\xbf\x7f\xd4\xc9\xcf\x6e\xd5\x46\x34\x9d\xad\xd2\x47\x34\x9d\xe3\x55\xb4\xf2\x88\xd2\xed\xb9\x33\x8c\x31\xb6\x8a\x56\x20\x51\x6a\x3b\x77\x9a\x63\x5a\x89\x44\x69\x09\x73\x6e\x84\x63\x5a\x91\x44\x9b\x08\xbf\xc2\x31\xad\x4c\xa2\xb4\x54\x3c\x77\x8c\x63\x5a\xa1\x44\x9b\x09\x17\x39\xa6\x95\x4a\xb4\x85\xf0\x7e\x8e\x69\xc5\x12\xdd\x40\x38\xcd\x31\xad\x5c\xa2\x9b\x08\x3f\xc4\x18\x23\x7f\x5a\xcf\x3f\x98\xff\x69\xee\xdd\x0b\x3b\x77\xef\x2a\xdd\x0c\x72\x6b\xc5\x3a\x60\x38\xb7\xff\xc7\x8c\xed\x1c\xce\x79\xe9\x72\xf7\x9e\xb1\x91\x91\x67\x35\x7d\x67\xe9\xf3\x8c\xb1\x8b\x85\x15\xfe\xd1\x48\x31\x52\x2c\x92\x3c\xbf\x13\x7f\x45\xbb\xc1\xfe\x3b\x4e\x0b\xae\xcf\xd2\x3f\xcb\x7f\x9c\xbc\x3e\x34\x62\x2d\xef\xcb\x9f\xcd\x8d\x37\x46\xa9\x96\x42\xe6\x7d\x73\x74\xf8\xd5\xd5\xd4\x18\x29\xda\x77\xf2\x22\xe4\xd7\xd2\xbf\xfc\xdf\xbf\x5e\xfa\xf5\xd7\x27\x55\xe5\xf4\x3f\x4d\x58\xbe\xbf\xe1\xba\x7e\x60\x2d\xe3\xba\xbc\x3b\x85\xb2\x1f\x08\x65\xd6\x6a\xd8\x7a\xc8\xbe\xfd\x9f\xdd\x48\xd1\x1f\x14\xfd\xea\xa5\x28\x63\xec\xd4\x3c\xaa\x53\xca\x3d\x3b\xf9\x7f\xfe\x47\x15\xfb\xb9\x48\x73\xa1\xcb\x5d\xe8\xd5\x73\x8f\xeb\xda\xe2\x83\x14\xa1\xaf\x73\x83\x34\xff\xf0\x3c\xfd\x98\xa6\xe7\x3b\xfd\x5a\x29\xf4\x11\x63\x63\x5d\xd3\xb4\x26\x3d\xd7\x0b\x80\x44\x36\x11\x4f\x29\xff\x11\x63\xf9\xae\x09\x5e\xe1\xa6\x8a\x7d\xbc\xe2\x22\xaf\xd0\x0a\xbd\x13\x85\xbe\x8b\x63\x9a\xff\x35\x00\xa4\x29\xdf\xe3\xd7\xf3\xbb\xfd\xee\x52\xab\xd0\x08\xae\x6c\x3a\x77\xb2\xf9\x81\xd1\xe2\x9c\xf6\xdc\x65\xdb\xf3\xc3\xab\x15\x7b\xbe\x65\xdb\xa3\x71\x9e\xaf\xd8\x3c\x27\xae\x72\xcb\x72\x8f\x4f\x60\xf1\xa1\x9b\x28\xea\xf3\x9e\x3f\xc6\x15\x95\x2c\x87\x78\x84\x8b\x4f\x50\xf5\xcb\x57\x2b\xa6\x73\xbe\xef\xf0\x8a\x0b\x15\xe7\x5e\xe0\x15\xe3\x85\xae\x8b\x85\xde\x0b\x85\xbe\xf1\x31\xcd\x6f\xd4\x7a\x76\xef\x55\xc6\x0a\xbd\x13\x79\xb7\x3f\xdf\x55\x2a\x5d\x99\x9e\xc1\x4f\x9b\xfa\xf2\x67\xf7\xe4\xc6\xdb\x67\x77\xf8\x35\xdb\x99\x7f\xe7\x6a\x26\xb8\xc5\xfd\x55\x03\xa0\x3b\x8c\xe6\x36\x5e\x9e\x66\xac\xf5\x4d\xdb\x6c\xad\xd0\x77\xb1\x30\x78\x61\x4c\xf3\x7f\x46\x01\xc6\x3a\xfd\xfc\x7b\xb1\x7c\xa7\x30\x36\x6f\xf9\xbd\x63\x9d\xfe\x46\x28\x40\xbe\x77\xbc\xf4\x82\xe8\x46\x58\x3b\x91\x3b\xd9\xfe\xa0\xd3\xdc\x39\xc7\x07\x5f\x97\xe6\xae\x9f\x9e\x6d\xbe\xdc\x37\x5d\x33\x5f\xb6\x4d\xcf\x34\x5f\xbc\xb5\x51\x55\x66\x8a\xe3\x5c\xf6\x6c\xb3\xed\xf9\xc6\xd4\x6c\xf6\xbc\x39\x55\x63\xcf\xf1\xa9\x99\xec\xf9\xe7\x5a\x7b\xfe\x70\xea\x7f\x61\xcf\x11\xdb\x9e\xa5\x53\x73\xcf\x5f\x36\x39\xeb\xfc\x7d\x75\x72\xc6\xf9\x7b\x7e\xb2\x66\xfe\x8e\x4f\xd6\xcc\xdf\x77\x26\x3f\x36\x7f\xdf\xae\xf5\xec\x2b\x93\xce\xf9\x7b\xc7\xe4\x2f\x9f\xbf\xbf\xd4\xdf\x77\x6c\x7f\x6f\x9e\x9c\x2d\xfe\xdb\x27\x6b\xe2\xbf\x61\x72\xa6\xf8\x9f\xaa\xb5\xf2\x17\x57\x66\x8b\x7f\x93\xb0\x47\x36\xe6\x1e\xd7\x95\xc5\x7f\x52\x6d\xd6\xaa\x17\xa5\x59\x4f\x5f\xa9\x98\x35\xa9\x94\x87\x81\xcc\x99\xc8\xff\x34\xaf\x8c\x75\x68\x1b\x5f\xd4\xf4\x92\xff\x0a\x63\x39\x1a\x06\x37\x00\xa9\xf8\xdc\x9b\x8a\x08\xbc\xe0\xdb\x42\x7c\xd3\x13\x92\xcf\xc7\xcd\xd2\x73\x25\xa5\xd0\xeb\xce\xbd\x4e\x23\x9b\xef\xf1\x6b\xa5\x02\x77\x6e\xfa\xdc\x70\x95\xec\x66\x92\x1d\x99\x4b\xf6\x27\x52\xf6\xee\x2a\xd9\xed\x24\x7b\x70\x2e\xd9\x2f\x4d\x09\xd9\x4f\x55\xc9\xee\x22\xd9\x6d\x73\xc9\x2e\x98\x16\xb2\x3f\x86\x53\xf6\x73\x24\xeb\x9e\x4b\xf6\x7e\x29\xfb\xb5\x2a\xd9\xfb\x48\xf6\xf4\xe5\x39\x64\xff\x42\xca\x06\xab\x64\xfb\x49\xf6\xe8\x5c\xb2\xa1\xab\x42\x76\x95\x3d\xee\x34\xa1\x69\xe8\xf7\xca\xa1\xa7\x79\x54\x1a\xbe\x5c\x1e\x7e\x3e\x87\x9a\x1e\xa8\x7a\xbb\x47\xf8\xfb\xbf\xaf\xb4\x97\xb2\x86\xef\xbd\xb7\x1a\xb8\x77\xcf\x3d\xb9\x71\x77\xfe\xe7\x85\xc7\xfc\xcd\xa5\xb7\x2f\x31\x36\xfc\x35\x2f\x63\x6c\xf8\x90\x9b\x31\x76\xcb\xe8\xeb\x53\xaa\x75\x4b\xeb\x25\xfb\x25\xff\x30\x72\xe3\xda\xbd\xcc\xf2\xaf\xe1\xe5\xe8\x5f\x12\x57\xee\x94\x92\x3b\xe9\xbd\xe3\x83\xc1\x9f\xe5\x4e\xba\xef\x7f\xd0\x78\xa0\x9c\x13\x58\xab\xf9\xfb\x79\xb4\xb0\xc6\x3f\xea\xcc\xc5\x9e\xba\xcc\xd8\x81\xd1\x81\xd5\xc0\xc7\xf2\xaf\x3a\xd5\xa9\x4e\x75\xaa\x53\x9d\xae\x05\xb9\xf8\xef\xc1\x2a\x76\xc6\xcd\x60\xd6\xf4\xa5\x83\xd9\xac\x2f\xe8\x8b\xc4\xe2\x74\x6f\x45\xd7\xf8\xc2\x66\xd6\x8a\x25\x83\x56\x2c\x95\xf4\x0d\xc6\xc2\xbe\x60\x32\x5c\x55\x37\x10\x0b\xfb\xac\x94\x2f\x14\x4d\x3d\x92\xbc\xfe\x13\xab\x89\x48\x3d\x19\x74\x07\x63\x71\x93\xd7\xa5\xd2\x66\xd2\x97\xce\xa4\x1e\x8e\x85\xcd\x30\xd7\x79\x8d\xba\x09\x5a\x9f\xdc\xe0\xf8\x35\x72\x9c\xab\x59\xdf\x76\xad\x22\x78\x2d\x34\xc5\x67\xd1\x14\x4a\x25\x12\xc1\x64\xf8\x7a\x84\xa2\x66\x68\x2f\x44\x87\xb0\x43\x8a\xb8\x13\xaf\x6f\x83\xad\x06\x7d\xc9\xbd\xc9\xd4\x23\x49\x5b\xde\xf7\xd0\xaa\xec\x43\xfc\x17\x67\x65\xb9\xeb\xce\x0d\x72\x7f\xc4\x77\xce\x33\xf6\x18\x80\x9e\xf7\x19\xdb\x0d\xe0\xe9\xf7\x19\x8b\xd2\xf5\x12\x63\xc7\x00\x4c\x5f\x62\x6c\x9c\x64\x2e\x33\x46\x4b\xd8\x46\x39\x8f\x95\x47\x77\x41\x79\x4c\x57\x96\xbb\x35\xed\xb0\x22\x7e\x33\x5f\x01\xe0\xc2\x79\xc6\xee\x24\x86\xce\x05\x3a\xd5\xb5\x50\xfd\xfb\x8c\xcd\x93\x72\x3e\xb9\xb7\xa3\xf9\xbc\xd0\x0f\x8f\xde\xed\x71\xff\xde\xe2\x85\x96\xb6\x0f\xbf\x7b\xd3\x9d\x9f\x69\xf3\xaf\xdc\x06\x80\x6c\x78\xf7\x43\xc6\x4c\xe2\xe9\xf2\xe8\x87\xd4\x1e\x8f\xfb\x29\x57\x87\xc7\xfb\x45\xad\xc3\xd3\x98\x9b\xb7\xc5\xe3\x4b\x78\xfc\x1d\x9e\xe6\x4e\xcf\x9a\x4e\x4f\x4b\xaf\xc7\xb7\xd5\xd3\xd8\x31\xea\xf1\x76\x9c\xf2\xb8\x3b\xc6\x3c\x7a\xc7\xf7\x3c\x9a\xb0\xeb\x55\x00\xfa\xa5\x8a\x0d\x54\xf7\x06\x80\x2b\x1f\x32\xa6\xfd\x1f\x7c\x06\xd4\xa9\x4e\x75\xaa\x53\x9d\xea\x54\xa7\x3a\xd5\xa9\x4e\xbf\x8a\x74\xe6\x88\xa6\xbf\x72\x44\xd3\xed\x3d\xea\x6e\x88\x3d\xe8\x8b\x68\xfd\xf1\xbc\xa6\xdf\x08\xb1\x67\x7b\xb9\xdc\xab\x7d\x13\xe5\xee\x27\x34\x7d\x85\xc4\x1f\x7e\xc4\x52\xe3\x07\x34\x9d\xf2\xfa\x33\x4f\x6a\x3a\xdf\xef\xfd\xa4\xa6\xd3\x22\x67\x58\xee\xb3\xa6\x5c\xff\x53\xb2\x3f\x5a\x7b\x63\x44\xd3\x55\x00\xdb\xe4\xfe\xee\x25\x00\x96\xd1\x7a\x20\xa7\xe9\x74\x7f\x3c\xa7\xe9\x0d\x52\x6e\x29\xad\x1b\x00\x7c\xc4\x58\xea\x70\x4e\xd3\x19\x63\x29\xb2\xf3\x02\x63\xa9\x64\x4e\xfc\x1e\xff\x3f\xa5\xc6\x13\x15\xb9\xee\xa2\xa6\xef\x29\x6a\x7a\xa4\xa8\xe9\x0f\x17\x35\xfd\x8b\x45\x4d\x7f\xb6\xa8\xe9\xdf\x2e\x6a\xfa\x6b\x45\x4d\x3f\x55\xac\xf0\x6e\xdd\xb2\xe5\xb7\x7d\xcd\x7d\xfd\x83\x49\x6b\xd0\x77\x47\x60\x7d\xa0\x65\x6d\xeb\xed\x83\x1c\xb6\x3e\xd1\xd6\x12\x68\x69\xbf\x45\x54\x43\xf2\x84\xe2\xc1\xe4\x80\xef\x61\x33\x93\xa5\x25\x66\x6b\x6b\xa0\x25\xd0\xb2\xb6\xed\x09\x21\xc2\x05\x02\xad\xb3\xd8\x78\xf4\x80\xe8\xd7\x25\x4f\x0e\xbc\x54\xc6\x62\x67\xfb\x5b\x65\xcc\x23\x8a\xf1\x32\x16\xab\x29\x1a\x03\x81\xc5\x8a\xeb\x4c\x19\xf3\x13\x01\xb0\x63\xe7\xc2\x02\x8e\x0f\x97\xb1\x70\xf7\x78\x19\x5f\xc7\xf1\x1b\x65\x2c\x37\x4c\x17\x6d\xbc\x90\xc3\x96\x32\xe6\x33\x88\x9f\xa5\x10\x78\x11\xc7\x34\x97\x04\xf6\x08\xf9\x63\x36\x5e\xcc\xe1\xc1\x17\x6d\x2c\x76\x83\x8f\x94\xf1\x12\x8e\xdd\x27\x6c\xdc\x20\xfc\x2b\xe3\xa5\x55\x63\xea\x92\xab\xe3\xe9\x32\x96\x33\x6f\xc4\xc6\x37\x08\x7b\xcb\x78\x19\xc7\x1b\xca\xb8\xa9\x6a\x1c\x5c\xb8\xb1\x72\x76\x83\xcf\xf9\x5f\x30\xf2\xb0\xe5\x88\xe0\x57\xb1\x88\x7b\x94\x76\xe0\xdf\x00\xf0\x8a\x03\xdf\x0e\x80\xef\x54\xe0\xcf\x58\x13\xba\x1d\xf6\x2b\x58\x8a\x84\x7c\x0e\x6d\xfe\x21\x87\xbf\x0a\x1a\xf0\x5c\x4d\xff\xb5\xf6\xbc\x0c\xe0\xb9\x6f\xd8\xfc\x4b\xaa\xec\xa7\xf6\xef\x3a\xfa\x03\x1a\xf0\x96\x23\x5e\x0a\x1a\xf1\x2f\x8e\xfe\xa8\xbd\xe4\x18\x0f\x7a\x0a\xa7\x1c\xf1\x53\x70\x03\x96\x28\xc0\xfe\xe7\xc5\x33\xdf\x40\xfe\x3a\x0c\xa2\xd1\x59\xaf\x00\x9b\xbe\xac\xe9\xef\xc9\xf6\x4e\x05\xf0\x1d\x11\xe7\x13\x1a\xd4\x45\x78\x5b\xa9\xc4\xdf\x87\x65\xf8\x9c\x52\x7d\x9e\x21\xae\x00\x87\xe4\x7c\x21\xfe\x27\x95\xca\xf8\x78\xb1\x0c\xf9\x9a\xfe\x96\x2b\x95\xf9\xd5\xa0\x7a\xf0\x02\xd9\xf7\x4c\xa5\xff\x3f\xa3\xfe\x9e\xd5\xf4\x3f\x96\xf8\xbb\x35\xf2\xdf\x77\xd8\x43\xfa\xff\xd5\xd1\xce\x37\xe9\x28\x40\xbb\x6c\x5f\xa2\x2e\xc3\x25\x05\x48\x7f\x49\xd3\xbf\x25\xf5\x2d\x50\xab\xcf\x5f\xdc\xa0\xd6\xc4\xa3\x06\x77\xab\x80\xf7\x79\x4d\x37\xa5\xfc\x9e\x9a\xf6\x31\x00\x9b\xca\xf6\x34\x61\x80\xf8\x8f\xc8\x78\xab\x8b\xf0\x85\x1a\xfe\x1f\x2a\xe2\x79\xbb\x53\xea\x7b\xa6\xa6\xfd\x9b\xaa\x78\xde\x6c\xfb\xfe\x5c\x75\xc6\xb3\x09\x7f\xa5\x02\x23\xcf\x6a\xfa\x66\x55\xc8\xff\xad\x0a\xf8\x8e\x56\xe2\xd5\xae\x00\xaf\x38\xf0\x3f\xd6\xe8\x7f\xc7\xa1\x8f\xe2\xf3\x55\x05\xb8\x70\xa4\xc2\xbf\x53\xad\x7c\x5e\x34\xa8\x0b\x81\x50\xc6\xca\x5a\x83\x91\x48\x20\x84\xb0\x99\x31\x07\x62\x59\xcb\xcc\x18\x56\xc2\x08\xc5\x53\x49\x33\x0b\xc3\x08\xa7\x8c\x81\x78\xaa\x3f\x18\x37\xc2\x56\x2a\x93\x35\x82\x83\xfb\x10\x4a\x25\xd2\x71\xd3\x32\xc3\x81\xdb\x6f\x6b\x6b\x9b\x99\xc9\x88\xc4\x92\x31\x23\x98\xc9\x04\x87\x0c\x33\x69\x65\x86\x10\xc9\x04\x13\xa6\x11\x1e\x4c\x24\x86\x60\x18\x0e\x64\xc4\x92\x31\xab\x8a\x55\x1e\x77\x31\xf6\x6d\xb8\xcd\xb0\x4c\xb2\x29\x10\x82\x61\x74\xef\xea\xe8\xed\x32\xba\x76\x74\x1a\x06\x8c\x6a\xa9\x30\x8c\xce\xfb\x76\x74\xf4\x6e\xdf\x52\xdd\xc2\x0f\xbf\xc0\x30\xb6\xee\xe8\x33\xba\xb6\x49\x0d\xdb\x3a\x77\xc1\xd8\xda\x73\xd7\xe6\x8e\x1e\xe3\xae\xee\xee\x7b\xba\x76\x1b\xbb\x3b\x36\xf7\x74\x19\xf6\xc1\x9a\x50\x76\x90\xdb\x2f\xcf\xdb\x6c\xda\x54\x39\x54\x63\x7f\xb3\x68\xd8\x47\x72\x0c\xb2\x32\x30\x60\x5a\x46\x3a\x64\x58\xd1\xc1\xe4\xde\x40\xff\x3e\x79\x96\xc7\x29\x38\x13\x5f\x1a\x86\x19\x0e\x5a\x41\x79\x02\xa8\xc2\xde\x2a\xfb\x29\xf7\x12\xaf\xed\x56\x1c\x1d\xaa\xee\x80\x54\x95\x3d\xae\x3a\xfb\x63\x18\xe1\x6c\xca\x88\x06\x93\xe1\xb8\x59\xfe\xb2\xb4\xe2\x42\xf5\xd9\xa4\x8f\x9d\x2e\xaa\xf2\x5f\x9c\x48\xaa\xee\xb8\x1c\x34\x8a\xbc\x38\xdb\xe4\xf4\xc5\x08\xc7\x8d\x8c\x19\x4f\x85\x82\x96\x49\x6a\xad\x58\xc8\x48\xc7\x4c\x7b\x98\xab\xd4\xf3\xc3\x4f\x55\xda\x23\x69\x23\xfa\x08\x0c\xa3\x3f\x9b\x95\xce\xf1\x13\x4f\xf1\xea\x00\x05\xad\x54\xac\xda\xa8\xdd\xbd\x5b\xe4\x5c\x41\x20\x3b\x94\xb0\x82\xfd\x08\x64\xad\x8c\xb8\x46\xed\xbb\x58\xd2\x32\x33\x69\x04\x92\x29\xcb\x0c\x0c\x24\x07\x03\xfd\x83\xb1\x78\x78\x6d\x2c\x2c\xab\x3a\x36\x6f\x5f\x6b\x05\x07\xc0\xdb\xa2\xc1\x6c\x14\x81\xf0\x50\x32\x3b\x94\x10\x57\x2b\x23\x5a\x64\x6a\x51\x05\x8c\x0c\x02\x19\x33\x4e\x7c\xe2\x26\x1d\xb7\xa8\xc3\x98\x85\x80\x65\xee\xb3\x10\xe0\x73\x2c\x90\x49\xf1\x39\x10\x30\xa3\xf2\xa9\x88\x86\x33\x15\x24\x24\xc4\x74\x16\x12\xf6\x7d\x78\x28\x19\x4c\xc4\x42\x08\x0c\xa4\x2c\xfe\x4f\x74\x20\x94\xf5\x67\xb3\x08\x84\x52\x89\x84\x99\xb4\xf0\xdf\xa5\xe5\x32\xf7\x54\x65\xde\x73\x54\x81\x7c\xab\x0b\xb2\x3f\x8e\x6f\x06\x78\xb6\xa2\xca\x7c\xe8\x25\x05\xf0\x3b\xf8\xec\xef\x91\x6f\x73\xf0\x51\x9e\xf4\x96\x52\x7d\x66\xcf\xe6\xeb\x04\x70\x89\xb1\x94\x2a\xf3\xa7\x71\xa9\x6f\x9e\x83\x8f\xca\x0e\x99\xc7\xaa\x32\xaf\xf2\xaa\xc0\x69\xf9\xdd\xbd\x22\x79\x28\x6b\xd9\xe3\x38\x7b\x48\xf9\xd6\x19\x55\xe4\xc0\xb5\x7e\x3c\x08\x80\xc9\x7e\x29\x0f\x4b\xba\x44\x5e\x6d\xf7\xab\xca\xb2\x57\xe6\xbd\xaa\xcc\xcf\x0e\xbb\xc4\x77\xf9\xce\x7e\x89\x1e\x05\x78\x96\xa6\xca\xbc\xed\xb8\x4b\xd8\xe3\xf4\x83\x3e\x7c\x73\x92\x6f\xb3\xcc\xe7\xde\x70\x89\x1c\x9c\xf8\x6e\x70\xf0\x3d\x2d\xf5\xf3\x4c\x91\x3e\xb7\xbd\x33\xc7\xf9\x29\x07\x1f\xe5\x7f\x2d\x5e\xe0\x98\x83\xcf\x2b\x79\x9f\x71\xf0\xd1\x7b\xea\xb4\x17\xf8\x6b\xbd\x9a\x8f\xe8\x39\x07\x1f\xbd\xcf\x9f\xbe\x11\x55\xd9\x98\xdd\xef\x9f\x3a\xe6\x0b\xcf\x23\x7d\xc0\x01\xf5\xe3\x7c\xdf\x76\xf0\xf1\x73\xab\x2b\xab\xcf\x52\xda\x7c\xaf\x3a\xf8\x28\xef\x1c\x59\x09\xfc\xdb\x0c\x7c\x7f\x07\xf0\x5c\xcf\x25\xf3\x25\xf7\xba\x4a\x9b\x73\xbe\xfc\x03\xc0\xb3\x5a\x97\xcc\xbb\xbc\xb3\xf0\xfd\x48\xfa\xea\x92\xf9\x58\xe3\xba\xca\x99\x5c\xe7\xb8\xfd\x44\xda\xe7\x92\x79\xed\xf4\x2c\xfa\xde\x75\xf0\xf1\x7c\xad\x45\xcc\x97\x5a\xbe\xf3\x0e\x3e\xca\x7b\x5a\x5a\x00\x7d\x06\x7f\x3f\x90\xfd\xbb\x64\x7e\xbc\xa1\x86\xcf\xbe\x9f\x90\xfa\xec\xbe\x88\x6f\x8f\x83\x4f\x71\x14\xc7\x30\xa1\xd8\x02\x5c\xd0\xc4\xf3\xff\x5b\x8e\xe7\xe8\x3a\xdb\x07\x49\xfb\x6f\x05\xfe\x43\xad\xd6\x07\xb9\x6e\x74\xf2\x0d\x6c\x00\x1c\xc7\xba\xcb\x7c\xff\x15\x00\x00\xff\xff\x77\x40\x86\x09\xf0\x3d\x00\x00")

func bindataSyscallx86testerBytes() ([]byte, error) {
	return bindataRead(
		_bindataSyscallx86tester,
		"/syscall_x86_tester",
	)
}

func bindataSyscallx86tester() (*asset, error) {
	bytes, err := bindataSyscallx86testerBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "/syscall_x86_tester",
		size:        15856,
		md5checksum: "",
		mode:        os.FileMode(509),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"/syscall_x86_tester": bindataSyscallx86tester,
}

//
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"": {Func: nil, Children: map[string]*bintree{
		"syscall_x86_tester": {Func: bindataSyscallx86tester, Children: map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}