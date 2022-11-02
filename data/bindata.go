// Code generated by go-bindata. DO NOT EDIT.
// sources:
// colours.txt (6.69kB)
// dogs.txt (1.348kB)
// gladiators.txt (221B)
// greek.txt (2.178kB)
// metals.txt (828B)
// trees.txt (559B)

package data

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
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

var _coloursTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x58\xdd\x9e\xa4\xae\x0e\xbc\xcf\x5b\x46\x8c\x9a\x31\x10\x36\x40\xbb\xf6\xd3\x9f\x5f\xe1\xcc\x7f\x7a\xcf\x45\x57\x55\x2b\x02\x42\xc8\x87\xbc\x2c\x72\x13\x27\x5e\x95\x41\x75\x58\x72\xe2\x6d\x63\x0d\xe2\x93\xc3\x99\xd8\x78\xe1\xd6\x25\x88\xcd\xaa\x47\x6f\xc4\x96\xbd\xac\xbf\x54\xb5\x08\xb1\x75\x27\xb6\x91\xb5\xe8\xc8\xc4\x99\x83\x4b\x3f\x20\xde\x5e\x88\xf3\x82\x2e\xb2\x84\x26\x2e\x0e\xd5\x8f\xbb\xf5\x4f\xe1\x03\x2d\x86\x49\x27\x2e\x7c\xea\xc5\xc4\xe5\xcd\x89\xb8\x72\x3a\x84\xb8\x56\x03\x86\x26\xef\xbf\xfc\x67\x60\xb0\x39\x87\x0f\x19\x5c\x19\x1d\x44\xe6\x55\xcd\x9c\x38\xc2\xaf\xee\x57\x21\x6e\x07\x71\xab\x1c\xbc\x8f\xf6\xaf\x3a\xd8\x3a\x71\xeb\xc1\x0f\xda\x24\x2f\x3c\x3a\x71\x37\x2e\x5d\x1b\x71\x77\x33\xe2\xb1\x48\xec\x73\xac\x97\x27\x5e\x9d\xf8\xaf\x9b\x77\x23\x7e\xb3\x09\x13\xbf\xbb\x24\xe2\xf7\x08\xa1\x85\x0f\x65\x5a\x38\x2f\xee\xb4\x70\xe1\x82\x7f\x65\xd5\xe4\xde\x69\xe1\x58\x24\xe2\x86\xf0\xd6\x70\xab\x75\x35\xc3\x73\x6f\xe6\xa0\x45\xf8\x25\x1f\x24\xed\xe2\xbf\xb4\x88\xee\x42\x8b\x44\x1e\x2b\xd3\xa2\x5c\x12\xc8\x16\x76\x5a\x34\xd2\x41\x8b\xb6\xc4\x37\x28\x73\x9c\xe0\x8e\xc9\x68\xc7\x7e\x3e\xd4\x2e\x91\xfe\x7f\xfa\xcd\x81\x66\xc6\xe9\xfc\xc4\xef\x39\x42\xa6\x11\xd8\x5e\x5a\xcc\x5b\xf3\x4c\x8b\x0d\xf9\x85\x9f\x86\x30\x86\x79\xa5\x1d\x3f\xe8\x79\xc1\x84\x1c\xd7\x3d\x56\xe1\xf1\x97\x16\xbc\x73\xf1\x17\xd3\xe2\xbd\x6b\x12\x33\xa5\xc5\x87\xad\x98\xa5\x8f\x3f\x03\xb3\xf2\x11\x8b\x17\x5a\x82\xd3\x29\x93\xcb\x7a\xd3\x12\xba\x4a\xcb\xac\x2b\x2d\xe1\x25\xf9\xa4\xb7\x7c\x53\x9f\x03\x85\x63\x8a\x81\xbd\xff\xc6\xb1\x2c\x26\x8d\x96\x91\x12\x17\xc1\x30\x63\xa5\x65\x6c\x1b\x2d\xa3\x9c\xf3\x7f\xe9\x5a\x76\x5a\x46\xec\x63\x0e\x34\xa2\x1c\x9c\x69\x99\xaf\x31\xb0\x5c\x69\xd4\x6f\x95\xd5\x4e\x4a\xbc\x70\x48\xa7\xc4\xa9\x8f\x46\x69\xda\x1d\x27\x4a\x6c\x9a\x7c\xd2\xe6\x51\x94\x21\xef\xda\x70\x29\x73\x60\x86\x89\xb3\x98\xf7\xc9\xf3\xb2\x8f\xcd\x78\xc7\x8d\xc2\x71\x7f\xd0\x6a\x62\xba\x1f\x68\x5a\x25\x28\x71\xe0\xd1\xc9\x11\xe8\x39\x56\x2d\x3c\x2f\x68\x6b\x79\x5e\xb1\x89\x73\x2f\x12\x47\xe1\xae\x5e\x28\x71\xe3\xc5\xa6\xc9\x40\xe2\x89\x96\x78\x45\x93\x76\x64\x89\x29\x9e\x21\x70\x02\x28\xc9\xca\x41\x49\x8c\x57\x3c\x2d\x26\x98\x90\x98\xb4\x2e\x60\x43\x13\xeb\x9a\x28\x49\x96\xd2\x29\x49\x70\x9e\x7f\x43\x9b\x7c\xd0\x30\xe1\xf2\x21\x0e\x5e\x4c\x1b\xd8\xce\x1b\x94\x97\xe0\x47\xf8\x73\x3d\x57\xde\x31\xf9\x03\xe7\xcf\x6c\xde\x8b\x39\xd7\x83\x63\x4e\xe7\xe1\xf2\x3c\x16\xe6\xbd\x3f\x37\xf3\xc4\x1e\x32\xda\xbc\xd0\x05\xa7\x2a\x1d\x52\x14\x4f\x49\xf8\x29\xf3\x3f\x0c\xf6\xa1\xcb\x7d\x9d\x72\x2c\xa0\xd6\xcb\xe8\x9f\x42\x13\xef\x0e\xde\xb6\x39\xb2\x16\x7f\xd0\x4f\x4a\x87\x27\x37\x9e\x83\xff\xaa\xd0\xd6\xf9\x19\x17\x52\x7f\x18\xaf\xa4\x05\x46\xfe\x90\x18\x76\x4a\x4b\x81\x21\x3d\x22\x63\x08\xf5\x34\xfb\x72\x4a\xda\xe3\x79\x0c\x4e\x69\x12\x2c\xcd\x58\xe3\xe5\x37\x0e\x64\xb2\xc7\x06\x6d\x6e\xc2\x6c\x6b\x3a\x2d\x3a\x99\x8f\xf5\xc1\x1b\x04\x4f\x92\x7c\x81\xc3\x4b\x9e\x7c\xbe\x9c\x6f\x1b\x96\xc3\xf7\x02\xbb\x75\x4c\xc7\x33\xba\xf3\xec\x94\xbc\xa4\x63\xa2\xc7\x3a\x39\x04\xef\xe7\x65\x93\x8e\xb7\xf2\xa2\xdb\xec\xb4\x74\x81\x17\x4b\x5e\xa7\x01\x7d\xd2\xa6\x62\x78\x18\xae\x15\xfd\x8c\xf0\x1b\x42\xf9\x59\x08\x8f\x13\x50\x26\x6c\xe6\xd7\x73\xf1\x43\xbe\x64\x6e\xae\xb7\x69\x5c\xde\xb2\x37\x4a\x7e\x2d\xe8\x27\x58\x0d\x58\x1e\x0f\x94\x42\x38\x03\x1d\x8b\xff\xcc\x36\x34\x37\x2c\x5d\x78\xf2\x55\xe7\x75\xbf\xda\x21\xbc\x52\x8a\x31\x4d\x34\x46\xf3\x87\x3a\x53\x82\x13\x03\x0e\xc3\x52\x8f\xaa\x2b\xa5\xbb\x62\xdd\x57\xde\x36\xf4\x41\x2b\x6b\xe9\x21\x42\x2b\x9b\x31\x6e\x94\x55\x0c\xa7\xec\x1f\x35\x16\xb4\xb8\x0a\xad\x92\x9c\x56\x31\x03\x74\x06\x8e\x5d\x68\x95\xa2\xf9\x3f\x8c\xe5\xa6\x55\x9a\x44\xa7\x55\x2e\x5a\x55\x9a\x18\xad\x5a\x76\x93\x9b\x56\xc5\x75\x6d\xe8\x48\xff\xaa\xd0\xea\x38\x18\xab\x5b\x3d\xb4\xd0\xea\x19\x76\xb9\x7a\x20\x30\xad\x7e\x95\x50\x6c\x38\xd4\x4d\x6b\xe8\xd6\xa7\x9d\xaf\x31\xed\x60\x1d\x45\x48\x78\x37\x21\x59\x16\x92\xc5\xcb\x4d\x92\x4c\x6b\x13\x92\xd4\x3a\xb7\x9b\x64\x95\x42\xb2\xee\x72\x31\xc2\x87\xac\x17\xc7\x4a\xb2\xef\x15\x71\xf1\x43\x98\xd4\xe3\x11\x99\x24\x4b\xb0\xad\x1f\xac\x45\x4a\x12\x92\x5c\x25\x3c\xc0\x21\xad\x91\x94\x55\xf8\x85\x14\x40\xca\xeb\x26\xf9\x33\xb8\xe3\x76\x9b\xb7\x9d\xa4\x4b\x14\xed\x37\xc9\x98\xee\x13\x2e\xf6\x1f\x59\xe2\x26\x79\x49\xec\x06\xcf\xb0\xb1\x25\x2f\xb4\x71\x99\x73\xdf\x04\x4b\x41\x9b\xe8\x97\x83\xa2\xfc\x07\xf3\x72\xeb\xfa\x62\xa3\x4d\x3a\xd3\xa6\xb0\xf3\x4d\x8b\x21\x40\xcf\xff\x85\x36\x85\xd1\x6f\x1a\x32\x61\xb3\x9b\x36\xe3\x2c\x88\x38\x10\x5a\xf6\x29\xfe\xd2\x66\xd2\x0e\xda\x4c\x4b\x07\x46\xa7\xcd\x39\xd3\xe6\x3b\x6d\xc1\x65\xd7\xca\x45\x69\x0b\x47\x17\xe1\xad\x3f\xb8\x68\x97\x47\x89\xd0\x36\xd2\xd1\x30\xf4\x90\xdd\x69\x87\xf3\x88\x7b\xb2\x22\x79\xda\x91\x47\xec\x42\xbb\xc8\x32\xca\x4e\xbb\x14\x67\xda\xe7\x0a\xe3\xc8\xef\x72\x37\x09\xda\x0f\xf4\xbe\xeb\xce\x8d\x76\xcd\x0b\x12\xab\x5d\x0b\xed\xfa\x42\x6f\xc6\x49\xd1\xca\x17\xc3\x45\xb7\xf5\x17\xa4\x84\xff\xab\xca\x0a\x77\xb0\x7b\x34\xa1\x1d\x01\x3b\xcf\x67\x5b\xd3\x4a\xbb\xf7\x43\x13\xed\x08\xc9\xda\xc0\x55\x3e\xf0\xc0\xab\xed\xc1\x2f\x31\xd0\xfd\x03\x22\xe5\x17\x0b\xaf\x73\x36\x23\x2f\x4e\xfb\x28\x2d\xfb\x29\xb4\x8f\x38\x0f\xa6\x03\x53\x2d\x2b\x84\x76\x25\x44\x84\xee\x85\xe0\xea\xe5\xcf\xd0\xa9\x2a\x1d\xc2\xfd\x78\x50\x82\x0e\x9c\xbb\x1e\x5e\x85\x0e\xc9\xe6\xe9\x04\x57\x3a\x14\x59\xd1\x68\x74\xe8\x7e\x60\x8b\xe9\x40\x9c\x8e\x9b\x0e\xcd\x6c\x7c\x33\x1d\x7e\x2a\x1d\xf3\x40\x1d\x5e\xe4\x6e\x23\x9d\x26\x74\x78\x9d\xd1\xff\xf0\x50\xe4\xb4\xc7\x88\x99\xcd\x0a\x1d\xa3\x9d\xa4\x09\x89\xcf\x4e\x6a\x36\x1a\x8e\x3c\xec\xe8\xf2\xc8\xa4\x65\xd5\xdd\x3f\xc8\x11\x30\x84\x34\xfc\x74\x60\x99\xd0\x66\xb6\xa2\x2f\x8f\x9b\xbe\x18\x21\x1d\xaf\x3c\x55\x07\xaf\xf2\x03\xdb\x86\xff\xfb\x2e\x01\x1a\x0c\xca\xcb\x33\xf9\x2f\x2e\x05\x58\xbd\x68\x82\x98\x51\xfc\x8b\x5f\x4c\x5f\x72\x89\xd1\x97\x17\xfc\xfe\x0c\x35\xfa\x9a\xcb\xfd\x35\x8a\xa2\xd1\xc9\xcb\x30\x3a\xb9\xec\x1c\xee\x74\x72\xf0\xc9\xa0\xb8\xe9\x14\xab\x74\x4a\xad\x62\x74\x1e\x7c\x2a\x9d\xba\x16\x9e\xae\xfd\x54\xe3\xcc\xe5\x8b\xc3\xa1\x8d\xa3\xc8\x4d\xa7\x22\xe7\xb7\x9b\x4e\x5f\x94\x4e\x3f\x7d\x65\x3a\x3d\xf2\x44\xcf\x7a\x3a\x9d\x3e\xa2\xe8\x89\xdc\xef\x1c\x70\x11\x64\x0c\xd3\x35\x1e\x21\x46\xc6\x2f\x99\x71\xe1\x57\x7c\xef\xae\x09\x42\xe3\x37\x2a\x02\x48\x12\x32\x45\xbe\xf5\x1f\xde\x64\x9a\x19\x20\xbf\x80\x05\x34\xf8\x20\x82\x77\xeb\x9a\x4e\x32\x67\xc4\x2d\xf3\xc5\xe6\x9e\xdb\xdc\x9e\x17\xc7\x54\x28\x64\x20\x46\xeb\x64\xbe\x73\x21\xc3\x49\x30\x2f\xd2\xfa\x6c\x03\xff\x63\x3e\xcc\x07\xd9\x48\xe7\x4d\x36\x66\xb5\x63\x37\x1c\x49\xe6\x45\x8c\x32\xaf\x5c\x76\x90\x22\xfc\x64\x5e\x83\x1b\x65\xde\xa5\x74\x06\x17\x37\x85\x38\x30\xc4\xfd\x29\xf4\x2d\xff\x21\xb6\x04\xe4\x04\x53\x4d\xf3\x64\xfd\xa3\x74\x19\x20\x83\x6b\xce\x6c\xb3\xef\xbc\x00\x0b\xc3\xb3\x7c\xf0\x0a\x7b\x99\x62\xe2\xc1\xbd\x33\x66\x36\x6b\x1c\x90\xcd\x66\x6f\x42\x49\x35\x3d\xc3\x53\x5b\x05\xd8\xe7\x4b\xfc\x50\x7b\x0e\x52\xe6\xe8\x5a\xf4\x87\xff\x0c\x74\x10\x6f\x78\x3c\xca\x48\x32\x31\x8f\xae\xad\xe1\x7a\x0f\xfd\x0b\xea\x12\x07\x02\x7e\xe6\xf1\x92\x07\xcd\x47\xfb\x47\xbd\x50\x36\x9e\x94\xc5\xb8\xa8\x3c\xfc\xc6\xc1\x43\xda\x5c\x7e\x31\x1c\x5d\x23\x33\x8f\x1b\x8c\x18\x98\x25\xe0\xdf\x40\xde\x29\x4b\x17\x8f\x6f\x9a\x6b\xa6\x6b\x99\x09\x75\xd6\x13\xb1\x32\xab\x61\xb6\x6a\x86\xb2\xe1\xa4\xac\xd9\x1b\x53\xd6\xb2\xc2\xbc\xe1\xe8\x01\x0d\x77\x02\x29\x7a\xd6\x96\x8e\x93\x29\xfb\x32\xcb\xe4\xec\x29\x71\x4a\x73\x68\x4f\x07\x6e\x7c\x41\x16\x46\x45\x86\x92\x79\xfe\xdb\x7d\xce\xd5\x4b\x9b\x6b\xe8\xe5\x8d\x96\x3f\x4e\x20\x7b\xc0\xb6\xb2\xb7\xb9\x88\x43\x9a\x29\xe5\x61\x4f\x5e\xf3\x21\x60\x82\x2b\xe5\x1b\x96\x4c\x85\xcb\xea\x41\x85\x2b\x53\xe1\x78\xe9\x49\x45\x96\x61\x4c\x45\xf6\xf0\xa2\x54\xa4\xb2\x01\x3b\x22\x7e\x91\x70\x2a\xf2\xe2\x95\xa9\x28\xef\x30\xae\xe2\x30\xd7\xe2\x99\x57\x2a\x1e\x17\xdf\x54\xc6\xbe\x4b\xa7\x32\x7a\x96\x9d\x9c\x9b\x36\xf2\xa5\x49\xbc\x10\xa2\x6f\xf2\x74\x84\x90\xab\x91\x9b\xbe\xe4\xc1\xf9\x1a\x50\x70\x74\x5e\xf8\xe0\x41\x5e\xe0\x16\xbd\xdc\x7f\xc9\x31\x0f\xaf\x3a\x32\x79\x70\x32\x01\x95\xfd\x83\xd2\xa1\xeb\x2f\xc9\x8e\x27\x43\x51\x59\x38\xf6\x35\x39\x79\xef\x9e\xb9\x90\xff\x45\x46\x54\x99\xc3\xa8\xf2\x62\x4e\x95\x93\x6e\x7a\x32\x04\xfe\xad\x03\x32\x57\x6e\x54\xb9\xcc\x8f\x09\x95\x4b\x3a\x70\xaf\xc6\xd3\x30\xe6\xd1\x84\x48\xc7\x2c\x60\x2a\x47\x7b\x3a\xee\x5a\xd0\x62\x76\xf2\x32\x38\xab\x2a\x9c\x8e\x5f\x44\xf2\x5c\x85\x03\x00\x65\x1e\x30\xdc\x0a\x77\x89\xc2\x0b\xd7\x02\x91\x1b\x19\xf0\xc8\x02\xd6\xd5\xe7\x65\xbd\xb4\x20\xba\xfc\x2b\x9b\x66\xb8\xb6\x2a\xad\xe3\xa9\x0b\x96\x55\x11\xef\xb8\x52\xd5\xc4\x48\x89\xe6\x77\x97\xaa\xe5\xa4\x3a\xdd\x77\xd5\xaa\x80\xaa\x85\x2a\x6a\x8e\x74\x28\x5a\xbd\xf9\xfd\x06\xbd\x99\x66\x86\xf6\x14\x80\xd5\x46\xfe\x06\x3f\x46\x1f\x27\x5f\x4c\xd5\x0d\x6f\xe1\x59\x10\xcc\x51\xbe\x54\xcf\x95\x57\xe4\x65\xd5\x03\xe5\x1f\x3a\xf7\x68\x73\x0d\x61\xa5\x1b\x2c\x7d\xaa\xfd\xb9\x82\xa8\x53\x43\x6c\xac\x42\x35\x34\x4f\x98\x27\xb3\x8e\x04\x90\xb9\x45\x23\xeb\xfc\x97\xeb\x89\x2e\x07\x7c\x65\x1d\x65\x67\xaa\x23\x2a\x56\x61\xf4\x7e\xd3\x9f\xa1\xe9\x6c\xf0\x2e\x08\xf4\xe9\xa6\xe0\x6d\x53\xa6\x60\x2d\x22\x14\xfc\xc5\x07\xcd\xbc\xaa\x7b\x77\x42\x96\x51\x28\xf8\xfd\xce\xdc\xf1\xde\x9f\x52\x60\xda\x21\xeb\xcf\x6f\xa6\xc0\x21\xb2\x51\x48\xbe\x29\xe4\xe5\x86\x74\x38\x66\x51\x17\x5a\xbb\xae\x42\x81\xd4\x21\xa6\xa5\x4d\x4c\xf2\x30\x4e\x5c\xa0\x26\x52\x9a\x59\xc6\x7c\xc5\xff\xe0\xe9\xdb\x3b\x6e\x22\xb9\x8f\xb1\xdc\x14\x23\x53\x8c\xd6\xa4\x7f\xd3\xe4\x4e\x8d\xd7\xd5\x84\x1a\x6f\x1b\x92\x81\x86\xa5\x6c\x8c\x22\x97\x1a\x4a\x9a\xc6\x26\x19\x98\xe7\xdd\x1f\xf2\xac\x68\x67\x1d\x0e\xb6\x71\x5e\x46\xc2\x03\xf0\xec\x93\x90\xe6\x4f\xf1\x38\x15\xa8\x4b\x5b\x85\xd8\x43\xd1\xb4\x1a\xfc\x59\xe3\x5a\x0f\xa4\xb6\x1f\x22\xb8\xfb\x8e\x26\xe3\xa5\x7b\x99\xe3\xbd\x25\x38\x51\x4b\x9c\xab\x82\x9e\x71\x12\x92\xb0\xfe\xff\x0c\x71\x68\x9b\xe4\x8e\xa0\xd1\x92\x7b\x7f\x38\x2a\x0c\xb0\x09\xef\xc3\x0c\x8c\x35\x6d\xc2\xed\x90\xe7\xff\x25\xb2\x52\x13\x43\xa5\xdd\xa4\x62\xaa\xdf\x38\x13\x45\xa1\x76\xf0\xea\xd7\x07\x9d\xd2\x70\x00\xe7\x1d\x44\xfd\x80\xc8\xd5\x7d\xf2\xdc\xc2\x4f\x11\x40\x35\xdc\xd5\xe0\x37\xb5\xc3\xd3\x39\x97\x42\x39\x53\xd3\x55\x12\xba\x50\x3b\x01\xb0\x89\xff\xa8\x2c\xbc\x52\xd3\x90\x02\x44\x69\x4e\x4d\x1b\x56\xe2\x94\x0a\x9b\x68\x88\xbb\xf4\xe4\xaf\xc0\x9b\x5a\x19\xdb\x46\xcd\xb9\x4e\xf8\xde\x0d\x37\xed\x3c\x17\x1b\x0a\xa7\xa5\x79\x2c\xa3\x51\xab\x92\x3a\xb6\xbe\xe2\x84\xb4\xaa\x65\x1a\x47\x45\xd6\xdc\x6a\xf8\xe8\xd4\xfe\x0c\x0d\x64\x48\x38\xe5\x27\x30\xda\xa1\x95\x5a\x57\x2c\xbe\xd3\x1c\xe3\x62\x2c\x67\xc7\x56\xb6\xc9\xd7\x83\x4f\xf0\x68\x3d\xf4\x94\xfc\x7c\xdc\x6d\x3d\x3c\x2f\x6e\x4a\xad\x8f\x55\x9d\xda\x58\xbe\x3f\xa5\xb6\x61\x83\xda\x28\xf8\xad\xcf\x56\x8d\x82\x42\x12\xfc\x5d\x97\xb7\x51\x76\xac\xe6\xa4\xeb\x93\xb1\x45\x78\xa2\x4a\xcc\x4f\x7c\x6d\xc4\x46\x6d\xb4\x43\xa9\x5d\x9c\x2b\xb5\x4b\xc3\xa8\xdd\x89\xb3\x87\x50\xe7\x85\x51\xd5\x76\x4e\x3c\xf1\x60\xea\x6c\xe8\xad\xcf\x8f\xbb\x66\xfe\xad\xca\x4a\xc8\x58\xbe\x7f\x68\x36\x13\xd7\x47\xc8\x9c\x3b\x14\x9a\x57\x5c\x84\x33\x8d\x9b\x3a\x8e\x15\xe0\x92\x29\x5a\x9e\x1d\x8c\x2a\xd4\x85\xf1\x3b\x01\x46\x5d\x72\xed\xb3\x38\xed\x52\x0a\x75\x94\x1a\x86\x06\x11\x9c\xbc\x77\xc8\xbf\xdc\xa8\x1f\xdc\xd3\x41\x1d\xf6\x6e\xf2\xc1\x63\xa6\xac\xdf\xbc\x68\xac\xd4\x75\x0e\xae\x48\x0f\xbb\xe2\xfc\x4c\x17\xd3\x67\xaa\x7c\xb9\x6d\xff\x48\xe7\xd6\xa9\xbb\xc9\xea\x20\xaf\x98\x9e\x57\x7e\x53\xf7\x96\x98\x7a\xf0\x2a\xd7\x5c\x87\xe0\x27\xa1\xef\x70\x80\x31\x3f\xfc\xf4\x10\x39\x7c\x34\x28\x2d\xba\x32\x9a\xc1\x78\xfa\xe0\x67\x0d\x86\x69\xa5\x3e\x50\x19\xce\x13\xf7\x8f\xc4\x82\x62\x9f\x67\xc3\x58\x1c\x38\x3f\xdf\x43\xfc\x19\xae\xe8\x78\xc0\x0d\xdc\xe0\x13\xd0\x70\xbc\xfa\xe8\x83\xfa\xa5\xcf\x77\xca\x7e\x61\x2e\xc3\x7a\xfc\x7c\x9a\x7f\xb1\x49\x49\xca\x8f\xe8\x70\xb5\x2f\xb6\x03\x49\x2b\xbd\xb8\xe8\xc3\xb1\x4a\xa1\x97\x94\xd1\xe8\x25\xb1\xea\x1e\x3a\x55\xd6\xf9\x39\xe5\x25\x6d\xbc\x14\x37\x35\x75\x87\x33\x7b\xe9\x3c\xc0\x2f\x45\x76\x0e\x94\x4e\x2f\x0d\x5d\x95\x0b\xbd\xdc\x57\x77\x7a\x0d\x4b\x5c\xe8\x9a\xf9\xfe\xc5\xea\x23\x06\x5d\x6c\x88\xdf\x17\x37\x5e\x94\xe6\xc7\x8d\xe4\x03\x75\xee\xc5\x1d\x1b\x79\x71\x1f\x4d\xe9\x9a\xdf\x3e\xe0\xd6\xaf\x27\xf3\xbf\xe4\xe2\x93\x2e\x14\x9c\x0f\x3e\x1f\xb7\xae\x43\xdb\x29\xf7\x64\x84\xe5\x6b\xe6\xe6\xdf\xa8\x28\xe1\x33\x61\xd3\x9a\x07\x5d\x8a\xd3\xa7\xfc\xaf\xd8\x86\x11\x06\x9a\x49\x35\x44\xa0\xec\x84\x78\xbc\xca\x5f\x2e\xbc\x0e\xba\x65\x1e\x8b\x1f\x1a\x99\xe9\xcd\x79\x91\xb7\xd2\x9b\x0b\x1f\xf4\x96\xd6\xe9\x2d\xa3\xd1\x5b\xf7\x7d\x04\x77\x7a\x6b\x29\x17\xdb\x8a\xc9\xbc\x35\x92\x17\x7a\x7b\x5e\x54\xe8\xed\xb1\x30\xbd\x07\x52\x57\xa5\xf7\xc8\xfd\xf0\xf8\x5f\x00\x00\x00\xff\xff\x93\x37\x2a\xed\x22\x1a\x00\x00")

func coloursTxtBytes() ([]byte, error) {
	return bindataRead(
		_coloursTxt,
		"colours.txt",
	)
}

func coloursTxt() (*asset, error) {
	bytes, err := coloursTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "colours.txt", size: 6690, mode: os.FileMode(0644), modTime: time.Unix(1667406788, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x23, 0xa3, 0x5f, 0xf2, 0x3, 0x92, 0x29, 0xb1, 0x38, 0x13, 0xc6, 0x2c, 0x54, 0xe9, 0xa7, 0xef, 0x11, 0x21, 0xfc, 0x5d, 0x58, 0x7b, 0x9c, 0x4f, 0xdf, 0x1d, 0xf2, 0x50, 0x16, 0xc1, 0x1, 0x25}}
	return a, nil
}

var _dogsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x53\x51\xb6\xeb\x20\x08\xfc\x67\x97\x44\x89\x70\x35\xe2\x41\x6d\xda\xae\xfe\x1d\x34\x7d\x1f\x95\xa9\x22\x83\xc3\x04\x4f\x8c\x6a\x80\xe7\x49\x95\x27\xf2\xc4\x8d\x9b\xd4\x1e\x98\x0c\x30\x1f\xd8\x19\x30\xcb\x40\xc0\x99\x92\x10\xe0\xec\x5d\x28\xaa\xc6\xf2\xfb\xd3\xf4\x02\xfc\xe2\x8d\x99\xe1\x40\x3b\x68\xc0\x81\x9d\xea\x9f\x78\xec\x8b\xe4\x20\x3c\x66\x29\x1e\x53\xa1\x1d\x84\x7c\xdf\xea\xaf\xda\x51\x54\x23\xeb\xac\x11\x0e\x0d\xd9\x4f\x95\xec\x50\x2a\x70\x68\xd1\x54\xa9\x13\x1c\x6a\xbb\xa0\xda\xef\x9a\xda\x57\x05\x0e\xed\x2c\xec\xa1\x8b\x6f\xbe\x9f\xb4\x37\x3d\xeb\xff\xf4\xf7\x85\x43\xb0\x3a\xba\x49\x8a\x9f\x9b\xa0\x45\x0f\x63\x60\xfd\xc0\x61\xca\x5a\x2e\x3f\x99\x29\x81\x77\xbe\xab\xcd\x52\xa2\xee\x8d\x0b\xfb\x90\xf3\x7c\xf0\x2a\x18\xf0\x85\x81\x75\x03\xcf\xf7\xd8\x54\x21\x30\x1e\xbb\xed\xc0\x5b\x80\xc0\x92\xc4\xd7\x47\xf8\xc0\xb2\xf4\x08\x2c\x55\x35\x7b\xf4\x3a\x2c\x4d\x56\xb8\x89\xaa\xf8\x35\xb5\xbc\xe2\x4c\xbe\xf4\xfc\x81\xa0\x21\xe3\xbe\xec\x68\xf1\x69\x29\x9e\xa5\x96\xc4\x59\x23\x06\xde\x73\x58\x88\x5d\xe2\x88\xbf\xb6\x23\x56\x7f\x48\xd4\xc3\x45\x32\x88\x5b\xfc\xe8\xb7\x7d\x75\xc2\xa8\x6f\x1f\xb4\x3a\x28\x04\xd1\xe8\x45\x06\x34\x0d\xbb\x33\x9f\x46\x35\xf0\xd0\xba\x90\x3f\x30\x79\x35\xaf\x9e\xb4\xc4\x45\xed\x80\xea\x33\x85\xb4\x3b\x4c\x46\x9f\x7d\x68\xf4\xd9\xa3\x67\xbc\xa4\x0c\xad\x7d\xe8\x0b\x8d\x80\xd1\xcc\x29\x18\x5f\xb8\x1c\xc0\x9a\x33\x4a\x54\xe0\xd5\x20\xeb\x0b\x6f\xb4\x01\xae\x86\x1b\x75\xa9\xf2\xd0\xfc\x61\xc8\xfb\xb9\x99\xa8\xb3\x3a\xd0\x4b\xab\x53\x66\x55\xc9\x64\xbe\xf9\x47\x90\xe7\x0b\xfb\x17\x0a\x2e\x47\x76\x72\x74\xa1\x61\x25\x73\x68\xdb\xbe\x0b\x45\xac\xf4\xa0\xcd\xb2\x70\x1f\x78\x9e\x0e\x7d\x26\x85\xb4\x1e\x64\xc9\x2f\x33\xf6\x35\x94\xa2\x77\x60\xaa\x70\x61\x19\x4e\xe0\x51\xfc\xc0\x9d\xb4\x44\xf8\x59\xea\xda\x9a\x5f\x33\x0a\x54\xba\x4f\xd7\xa5\x60\x8d\xa0\x63\x78\xc3\xfe\x9e\x86\x4d\x4a\xd1\xba\x80\x57\x69\x44\x7b\xf8\x8d\xb2\xd4\xe4\x0c\x4d\x86\x37\xd3\x8a\x8e\x01\x4d\xa5\x0e\x32\x68\x7a\xed\x34\xbd\x02\x8b\x87\xed\x48\x07\x86\xd5\x47\xd6\xf4\x5a\xf7\x54\x97\x9b\xdb\x7e\x65\x9b\xc9\x7f\xdb\x6a\x6d\xa6\xb4\xf7\x64\xa9\xd5\x66\x11\x68\x1f\xfb\x7d\xcb\xa6\x63\xbf\xc9\xc1\xf3\xff\xf9\xd8\x3a\x96\x99\x05\x3a\x5e\xfa\xa1\xe8\x71\xd9\xb8\xbb\xd9\x1b\x59\x26\x87\x75\x97\xe9\x4c\xd4\x7e\x32\x77\xa6\xb6\x32\x59\x56\x5f\x9d\xe5\x69\xb8\xf3\xb6\x53\x7f\x3e\x8f\x5e\x74\x26\x16\xe8\xcd\xa4\xa6\xd5\x47\x1f\x78\x2c\x87\xc1\x20\xb3\xa5\xd8\x4b\xbe\xbd\x20\xdc\x24\xbf\x41\xdf\xd4\xc7\xb6\xf9\xbd\x9a\x19\x70\xf3\xe6\x7e\x6b\x51\x19\xdf\x30\xa5\x8e\x22\xf0\x71\x1e\xcf\xfb\x17\x00\x00\xff\xff\xe3\x18\xda\x70\x44\x05\x00\x00")

func dogsTxtBytes() ([]byte, error) {
	return bindataRead(
		_dogsTxt,
		"dogs.txt",
	)
}

func dogsTxt() (*asset, error) {
	bytes, err := dogsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dogs.txt", size: 1348, mode: os.FileMode(0644), modTime: time.Unix(1667406788, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0x56, 0x5a, 0x32, 0x1, 0xe8, 0xf, 0x60, 0x5c, 0x99, 0xd8, 0xcc, 0xbf, 0xf3, 0xb2, 0x37, 0xec, 0x91, 0x1b, 0x8, 0x1e, 0xb0, 0x7e, 0xda, 0x66, 0x32, 0x76, 0xb0, 0x8b, 0x41, 0xcb, 0xd6}}
	return a, nil
}

var _gladiatorsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xcd\x31\x12\x83\x30\x0c\x04\xc0\xfe\x7e\x79\x98\x03\x2b\x28\x12\x23\x9b\x90\xe1\xf5\x99\x74\xdb\x2d\xdf\x7c\x32\xb0\x38\x1f\x61\xa3\xb7\x0c\x6c\xce\xb7\xb0\xe5\x17\x7b\xfa\x8a\x97\x26\x9c\x43\x05\xb7\xbd\xcf\xb0\xd8\x11\x7f\x8d\xce\x55\x38\x19\xb3\xab\x70\xf6\x54\xd8\x17\xa5\x45\x8e\xb2\x44\x65\x3b\x34\x31\x5a\xd6\x69\x89\x61\xa5\xc0\x27\xf7\x4b\x78\x72\x35\x36\xb0\x09\xcb\xe5\x6e\x13\x2d\x97\x22\x56\xd3\x90\xa3\xf3\x3e\xd0\xaf\x98\x2a\x1c\x9d\x81\xa2\xad\x2a\x54\xb7\x48\x0c\x16\x9b\x02\xff\x3f\x6f\xcc\xca\x17\x03\x9f\xcb\x1b\x03\x37\xab\x2c\x0b\x77\xfa\xf6\x0b\x00\x00\xff\xff\x3d\x2f\xc3\x00\xdd\x00\x00\x00")

func gladiatorsTxtBytes() ([]byte, error) {
	return bindataRead(
		_gladiatorsTxt,
		"gladiators.txt",
	)
}

func gladiatorsTxt() (*asset, error) {
	bytes, err := gladiatorsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gladiators.txt", size: 221, mode: os.FileMode(0644), modTime: time.Unix(1667406788, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1f, 0xc, 0x5, 0xbb, 0x70, 0x1c, 0x50, 0xf7, 0xd1, 0x98, 0x95, 0xb0, 0xb, 0x67, 0xf2, 0x5d, 0x11, 0x9e, 0x27, 0x86, 0x83, 0xf9, 0x57, 0x95, 0x82, 0x2a, 0x85, 0x72, 0xf4, 0xd4, 0x4b, 0x6f}}
	return a, nil
}

var _greekTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x55\x51\x96\x2e\xad\x09\x7c\x67\x2f\x59\x14\x9f\x5d\xb7\x25\xbf\x8a\x07\xed\x9c\x71\xb6\x94\xf7\x6c\xe0\x6e\x2c\xa7\xe8\x99\x87\x96\x52\x5b\x11\x28\x40\x0b\xd6\xc6\x12\x2d\xd5\x5a\x4b\x60\x1c\x7c\x6c\x7b\x08\xb6\xc2\x87\xe8\xd5\xb1\x39\xbf\x7c\x70\x1f\x5a\x38\xc3\xc5\x3d\x78\xcb\xc9\xae\xa1\xa2\xb7\x76\xf4\xc1\xf5\xbb\xe9\x13\xbe\x44\x1b\x95\x58\x02\xdb\xd5\xff\xfe\x97\xa8\xbf\x17\x37\xd7\x4b\x21\xda\xb5\xed\xaa\x50\xa2\x6f\x6e\xf4\x59\xed\x57\xee\xb0\x8d\x5f\x78\xb8\x3a\x74\xfc\x03\xd1\x51\x14\xd4\x3d\x4a\xb5\xc5\xd7\x8f\x2b\xbc\x6b\xa9\xf8\x81\xb8\x54\x74\xa4\x29\x63\xff\xfc\xbb\xad\x79\xa9\x3f\xd0\x27\x44\x67\x0d\xbf\x52\xc5\xf4\xd6\x5c\x34\x78\x57\x60\xd7\x67\xa9\x68\xdc\x3e\xf4\xd9\x04\x3c\x15\xa6\xd7\x00\xe5\xfa\xb9\x32\x36\x3a\xed\x5b\xa5\x61\xa6\xb2\xb5\x8f\x0e\xfd\x12\xdd\xda\xa8\x59\x74\x57\xed\xba\x28\x31\x38\x6d\x9c\x3c\x37\x52\x6c\x6f\x87\x1e\xfd\xe8\x43\xff\x7f\xd0\x1a\xc2\x67\xf5\x21\x1f\x8c\x8b\x4b\xb6\x39\xf1\xe0\x81\x4f\x98\x06\x35\x7f\xc2\xb6\x77\x0d\x7a\xf7\xf3\x2c\x0b\x5b\x52\xf4\x03\x0b\x93\xa2\x57\x7f\x38\xc5\x40\xca\xa6\xb9\xdb\x4a\xd5\x94\x2d\x8d\x4f\xb0\xb6\x4b\xd1\xb5\xe8\x34\x25\xda\xda\x4c\xa5\xa0\x84\xcf\x25\x05\x63\xeb\x13\x52\x30\xab\x32\xd8\xa5\xaa\xe7\x18\x3e\xa4\x54\xeb\xe0\xb1\x6a\x39\xb5\x28\x90\xd2\xe0\x1f\xea\x6b\xe6\x52\xda\xd9\xe8\x03\x6b\xf3\x37\xdf\x67\x49\x09\x1f\xbc\xe8\x94\x96\x1a\x4e\x60\x40\x2e\xc5\x95\x0a\x2e\x1d\x9a\x62\xd6\x5c\x9e\x95\xbc\xbb\x34\x2e\xe5\xb1\x0b\x1d\x1b\x41\xf9\x7a\xe9\xc2\x53\xb4\x91\x30\x97\x31\xea\x58\x04\x3c\x6a\x3e\xce\x7a\x72\xba\xca\x13\x26\x28\xd5\xae\xa1\x94\x2e\xb0\x06\xdb\xf5\x98\x0a\x1a\x0a\xdf\x87\x71\x9d\xce\x9b\xe0\x4b\x10\xba\x5d\x10\x28\x75\x57\xba\x11\x61\x37\xef\x05\x5d\x0d\xd2\x1b\x4f\xf7\x36\xb9\xf5\x84\x4f\xa5\x38\x97\x15\x08\x9e\x8d\x98\x10\xfc\x47\xc7\x85\x90\x3f\xda\x55\xfe\xe8\x86\xfc\x79\xc2\xb0\xe4\x26\xdf\x6f\x6d\xba\x5f\x69\x63\x57\x53\xae\x8f\x43\x1b\x84\x19\x44\x5a\xdc\x4e\x0a\xca\x1d\x5a\x20\x55\x69\x5d\xd5\xe8\x3e\x4c\x09\xe6\x91\x8a\x0f\xa4\xa2\xf0\xf6\x8a\xb2\x3d\x28\x9e\x8f\x4a\x45\xc3\x78\xc7\x67\x51\x9a\xa7\x78\x57\x67\x55\xac\x9d\x1b\xa1\x39\x14\x56\x81\x8a\xe8\xbf\x59\xf1\x6e\xf6\x77\xd5\xa5\x62\x4d\x84\xe5\x1b\x12\xe6\xfe\xda\x7c\x8a\xcd\xe9\xed\xe4\x09\xe7\x75\x47\x0b\x6d\xe2\xfc\xbc\xaf\x3e\x57\xae\xdf\x06\xfe\x7f\x98\x04\xf5\x74\x3e\xe5\x4c\x44\xd2\x3b\xf1\xe0\x23\xcf\x5c\x36\x4f\x83\x98\x96\x4c\x58\xd3\xc5\xc0\x58\x51\x6a\xb5\xcb\xfb\xcb\x6d\x6b\xbe\xc4\x5c\xcc\x1b\x59\x63\xb3\xda\x8d\x41\x15\x99\x11\xf6\xc5\x53\xff\xd6\xe5\x43\x1a\xed\x8d\xc3\x18\x2e\x69\xda\x4d\xa5\xa9\x17\xff\xfb\x3f\xee\x91\x3a\xfc\x09\x6f\xc4\x1a\xcb\x47\xc3\xae\xe0\xe8\xd2\xf0\x14\xdf\x15\x2a\xcd\xe8\xcd\xe6\xfb\x59\xff\x82\x92\x8a\xed\x29\xf6\x87\xf2\x14\xf5\x21\x5d\xc7\x76\xe9\x1a\xeb\xe8\x12\x46\x53\x39\xb2\x9e\x74\x34\xed\x24\x4c\x47\x83\xde\x08\x82\x99\xb6\xc8\x4f\xed\x24\x4e\x53\xba\x5d\x3c\x6e\xf4\x07\xc7\x4c\xc3\x3e\xd0\x7d\x1d\xfe\xee\x31\x93\x96\xfd\x59\x90\x7e\xa2\x1b\xdf\x3f\xd4\xf4\x92\xa1\x51\x6c\x91\xfb\x03\x1d\xcb\x28\x7d\x6e\x6f\xe8\xb9\x16\xb0\x2b\x45\xce\x16\x39\x33\xec\x1f\xc8\x48\xdb\x86\xf9\x87\x38\xcf\x9f\x2f\xf1\x82\x4c\x3d\xbf\xce\x5a\x3c\xe2\xb8\x6c\xa6\x1c\x4c\x08\x8f\xb7\x93\x78\xd0\xd9\xbf\xef\x9a\x3a\xf8\x5d\x19\x30\x02\xf2\x62\xb2\x74\xca\xc4\xad\xbc\x7c\xa2\xe1\x15\xe4\xfe\x44\xd6\x83\x49\x07\xb0\x3c\x4d\xc4\x02\x53\xfc\x07\xf2\xcf\xca\x5e\xe3\x43\x66\xb5\x86\xfe\x03\xbc\x6c\x50\xff\xac\x5e\xf2\x27\x67\x4a\x50\x0c\xfb\x92\x69\xc1\xd6\xc3\x8d\x06\x4b\x2e\xce\xf6\x90\xa9\x64\x6c\x67\x22\x11\xcc\x9a\xbe\x21\xfc\x62\x99\x9e\xbe\x90\x1e\x9d\x61\xda\x73\xa4\xc9\x33\xbc\xc4\x93\xe6\xce\xf0\xec\x8c\x93\xcd\xe6\x2d\x11\x33\x7c\x63\x59\xc6\x2f\x31\xe5\x3a\x6c\x4a\xf3\xdc\xfd\x2d\x53\xf3\xa4\x09\x41\x32\x2d\x8d\x99\xbc\x5b\xba\x4f\xc8\x62\xae\x42\x16\xc8\x0f\x59\x08\x9d\xb6\x64\xd9\x9b\xc1\xcb\x02\x43\x96\xad\x33\x99\x12\x6b\x9f\x2f\xd9\x6c\x31\x2c\x9b\x5b\x63\xa7\xab\x37\x1a\xee\xac\xb0\x44\xec\x88\x09\x33\xd6\xac\x4a\xcb\x4a\xf5\x80\xec\x9a\x85\x9e\xcd\xe9\x30\x28\xbb\xea\xd0\xed\x04\xd9\xd2\x76\x45\xfa\x7c\x57\xb0\xc7\x6c\x0b\x2c\x86\x69\xdb\xd6\xc1\xb1\xbe\x4a\x22\xdb\xd3\x0e\xb7\x7c\x46\x1a\xbb\x4f\x96\xe6\x27\x94\xee\xa5\x78\x96\x7c\x67\x90\xbe\xb3\xb5\xca\x37\x9e\xf5\xff\x00\x00\x00\xff\xff\xdc\x93\x09\x62\x82\x08\x00\x00")

func greekTxtBytes() ([]byte, error) {
	return bindataRead(
		_greekTxt,
		"greek.txt",
	)
}

func greekTxt() (*asset, error) {
	bytes, err := greekTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "greek.txt", size: 2178, mode: os.FileMode(0644), modTime: time.Unix(1667406898, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc2, 0x61, 0x82, 0x64, 0xd9, 0xa3, 0x1f, 0xc5, 0x78, 0xa6, 0x55, 0xe9, 0x13, 0xab, 0x99, 0x44, 0x34, 0xb, 0x70, 0x3a, 0x7e, 0x4a, 0x30, 0x11, 0x5d, 0xcb, 0x52, 0x7c, 0xc1, 0x93, 0x5b, 0x85}}
	return a, nil
}

var _metalsTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x52\xd1\x8e\xed\x20\x08\x7c\xf7\x2f\xa9\x72\x2a\x59\x04\x83\xd8\x9b\xee\xd7\xdf\x14\x7a\xf6\x89\x81\xc8\xe0\x0c\x40\x75\x12\xda\xa3\x00\xef\x41\xf2\x80\x81\x46\xf5\x29\x1d\x60\x11\xd0\x7e\x90\x5f\x74\x73\x22\x5a\x63\x7b\x2f\x87\xf6\x78\x53\xa1\x8d\x8c\x5c\xdf\x48\x1f\xb5\x60\xae\x98\x4f\x70\x45\xe8\xa6\xf9\x54\x0f\x60\x2f\x55\x27\x9a\xe4\xc0\xaa\x73\xa2\x95\xba\xa3\xa1\x81\x8d\xe5\xd0\x3c\x92\x7d\x04\x59\xbb\xd7\x34\x0d\x26\x24\x59\x8e\xf9\x7b\xb4\x23\xc2\x36\x9d\x0f\xf8\xa0\xc5\x90\x0f\xa3\xe9\x15\xc8\x40\x62\xc8\x09\x4d\x39\xbb\x4e\x48\x35\xa7\x72\x2b\x1d\x3e\x51\xec\xb0\x82\xbe\x2b\x07\x05\x49\x8b\x60\xf4\x46\x95\xc2\x20\xde\xe1\x71\x8b\xe1\x9f\x61\x12\x33\x42\x2b\x4c\xde\x23\xa1\x0b\x6d\x68\x08\xe1\xed\x18\x22\x06\x9c\x92\x2e\x0c\x90\x13\x04\x17\x96\x81\xe4\x92\x0e\x0d\x94\x86\x8c\x57\x62\xab\xdb\xee\x32\x94\xef\xa3\xe1\x33\x6a\xe8\xaa\xa9\x45\x50\xdb\x3d\x12\x4d\xdf\xf1\x6d\xa1\xfa\x83\x5c\x84\xba\xbe\xb9\x86\x25\xa2\x47\x2e\x4f\x57\x74\x4c\x60\x86\x50\x32\x19\x3c\x36\x3e\x79\x7b\x36\x4d\xe5\x2f\xf0\xd7\x86\x69\xb0\xfe\xc6\x4d\xd3\x81\x29\x70\x9a\xfa\xdf\xf1\x58\x32\x5a\xc7\x4c\xbb\x66\xae\x28\x7e\xbe\xb5\x7d\xa4\x81\xb6\xfd\xfb\xec\x41\xf6\x51\x8b\xfa\x82\x91\xf7\xb6\x2a\xa4\xe5\x0b\xe1\x50\x3b\x03\x12\x5f\x68\x65\x25\xef\x72\x53\x09\x47\x1d\xc4\x9f\xcb\x2d\x8e\xb5\x4b\xba\xec\xef\x31\x78\x7f\xd7\xeb\x3d\xf7\xe0\x7d\x67\x4e\x52\x9c\x1c\xe2\x13\xbe\xe5\x5c\x8e\x52\xb6\x65\xe1\x02\x49\x35\xb7\x7f\x99\x6e\xf7\x20\xf8\x25\xa9\xe5\x97\xac\xa6\x49\xff\x03\x00\x00\xff\xff\xca\x6c\x8b\x01\x3c\x03\x00\x00")

func metalsTxtBytes() ([]byte, error) {
	return bindataRead(
		_metalsTxt,
		"metals.txt",
	)
}

func metalsTxt() (*asset, error) {
	bytes, err := metalsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metals.txt", size: 828, mode: os.FileMode(0644), modTime: time.Unix(1667406788, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5c, 0x15, 0x57, 0xd3, 0xa6, 0x2e, 0xca, 0x1e, 0x1b, 0xed, 0x88, 0xdc, 0x11, 0xc4, 0xfe, 0x81, 0xf7, 0xa, 0x1b, 0x3a, 0xb3, 0x90, 0x39, 0x3, 0x67, 0x84, 0xe9, 0x7f, 0x15, 0x71, 0x81, 0x64}}
	return a, nil
}

var _treesTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x51\x51\xae\xe5\x2a\x0c\xfb\xf7\x2e\x0d\xcd\x2d\xdc\x26\x84\x17\xe0\x54\x5d\xd6\x5b\xc3\x6c\x6c\x44\x8f\x34\x3f\x06\x62\x09\xdb\x31\x33\x73\x25\xa8\x87\x04\xd8\xbb\x0a\x18\xc9\xe3\x53\x27\x05\x1c\x05\x1c\x5d\x1a\xf8\xf1\xcc\xc3\x91\xa8\x83\x86\x44\x4b\xbe\x5f\x8d\x8d\x48\xf4\xc4\x84\xc4\x07\x49\x24\x17\xa4\x1a\x1b\x95\xf9\x9a\xc5\xa3\x21\xe9\x92\x73\x19\xd2\xca\x97\x3c\x82\xb4\xe6\x94\x68\x6b\x22\xcb\xc1\x40\x2e\x12\xf1\xec\x63\xcc\x77\x5a\x6a\xbb\xd8\x6b\x43\xf6\xb8\x90\x7d\x4e\x6f\xb7\xfb\x81\xbc\xf2\xb2\x24\x81\xfc\xf4\x90\x31\x70\xf8\xf9\x12\xf2\x46\x10\x35\xfc\xd4\xc0\xb9\xf8\x21\x0a\xf3\x95\xde\x9f\x0b\xef\xaf\x93\x22\xa6\x9e\x2f\x94\x9a\x2f\xdf\x84\xab\x6e\xec\x9b\x4d\xc2\x86\x7f\x97\x5f\x66\x06\xdb\x41\xfc\xae\x56\xbb\x04\x94\x3b\x97\x72\x85\x28\x54\xcc\x1b\xb4\x9a\x40\x3d\xaf\x31\x61\x3c\xc2\x9b\xc0\x78\x36\xd7\x4a\x18\xf7\x42\xcd\x7d\xc8\xeb\xd1\x96\x7e\xed\xd8\x13\x53\x05\x6d\x4d\x93\x13\xce\x0b\x1e\x6c\xa7\xa0\x53\x0d\x5d\x18\xe8\x12\xa3\xda\x96\xe8\xb5\xc9\x0b\xdf\x7e\x7a\xfd\xf3\xff\x9e\xba\xc9\x19\x6c\x9c\x82\xee\x5d\x19\x08\x39\x5e\x99\xc1\x31\xf8\x13\x1c\x18\xf2\xdf\xf2\x4a\x8c\xc2\x23\xad\x51\x30\x7a\xac\x2c\x18\xcb\x98\x31\x9e\x4c\xf3\x10\x4c\x1a\xa3\xb6\x03\xd3\x23\xe4\x21\xe6\xea\xa2\x8e\x9b\xba\xdb\xb8\xab\xaa\xdf\x78\xe4\xfe\x1b\x00\x00\xff\xff\x7b\x57\xf0\x15\x2f\x02\x00\x00")

func treesTxtBytes() ([]byte, error) {
	return bindataRead(
		_treesTxt,
		"trees.txt",
	)
}

func treesTxt() (*asset, error) {
	bytes, err := treesTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "trees.txt", size: 559, mode: os.FileMode(0644), modTime: time.Unix(1667406788, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa7, 0x53, 0x77, 0xfc, 0xb7, 0xaa, 0x44, 0x6d, 0x27, 0x72, 0x1a, 0x1b, 0x99, 0x9, 0x35, 0x89, 0xab, 0xd9, 0x4f, 0xac, 0x5c, 0xc, 0xe1, 0xfd, 0x61, 0x2b, 0xdf, 0x31, 0xe3, 0x83, 0xca, 0x90}}
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
	"colours.txt":    coloursTxt,
	"dogs.txt":       dogsTxt,
	"gladiators.txt": gladiatorsTxt,
	"greek.txt":      greekTxt,
	"metals.txt":     metalsTxt,
	"trees.txt":      treesTxt,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"colours.txt": {coloursTxt, map[string]*bintree{}},
	"dogs.txt": {dogsTxt, map[string]*bintree{}},
	"gladiators.txt": {gladiatorsTxt, map[string]*bintree{}},
	"greek.txt": {greekTxt, map[string]*bintree{}},
	"metals.txt": {metalsTxt, map[string]*bintree{}},
	"trees.txt": {treesTxt, map[string]*bintree{}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
