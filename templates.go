// Code generated by go-bindata.
// sources:
// templates/login.html
// templates/password.html
// templates/profile.html
// templates/register.html
// templates/verify.html
// templates/verify_email.html
// DO NOT EDIT!

package user

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
	info  os.FileInfo
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

var _templatesLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\xcb\x6e\xdb\x3a\x10\xdd\xe7\x2b\xe6\xf2\xa6\xb0\x8d\x40\x96\x93\xb6\x1b\x45\xd2\x26\x48\x56\x05\x9a\x22\xed\xba\xa0\x45\xca\x66\x4a\x91\x02\x1f\x8e\x03\x41\xff\x5e\x92\x92\x2b\xca\x28\xb2\x48\x69\x80\xf0\xbc\xce\x1c\x1e\x0e\xd5\x75\x84\xd6\x4c\x50\x40\x56\x53\xf5\x93\xcb\x1d\x13\xa8\xef\x2f\xba\xce\xd0\xa6\xe5\xd8\xb8\xc8\x9e\x62\x42\x15\x82\xb5\xf3\x5f\xe4\xda\xbc\x72\x0a\xe6\xb5\xa5\x05\x32\xf4\x68\xd2\x4a\x6b\x54\x5e\x80\x5b\x84\x1d\xd6\xb5\x54\x4d\x52\x33\xca\x09\x74\xc1\xe9\xd7\x0b\x23\x66\x9f\xc1\xe7\xcd\x87\xdb\x3f\xbe\x06\x2b\xd7\x2b\x83\x0d\x60\x6b\x24\xdc\x6c\xda\xe3\x10\xec\xc3\x6e\xf9\x9a\x2a\x25\x55\xd2\x50\xad\xf1\x8e\xea\xf7\xc2\xf9\xc5\x99\x36\x49\x20\x9e\x81\x90\x82\x4e\x91\x16\x13\xc2\xc4\x2e\xe1\xb4\x36\xae\xf8\x6d\x06\x9c\x45\x24\xb6\x52\x39\x55\x92\x4a\x72\xa9\x32\xf8\xff\xe1\x61\xe3\xd6\x6d\x14\x3e\x26\x7a\x8f\x89\x7c\xf1\x9c\xfc\xef\xba\x3d\xfe\x25\x0d\x57\xbf\x76\x4a\x5a\x41\x22\xa4\xbb\x9b\xbb\x8f\x53\xca\xc9\xbf\xd9\x9c\x77\x08\x04\x14\x26\xcc\xea\x0c\x3e\xc5\x07\x1e\xd4\x48\xb6\xd2\x18\xd9\x64\x70\x3d\x53\x63\x7e\xe6\xeb\x48\xf8\x3c\x0d\x22\x95\xfe\x9a\x2b\xc5\x5a\x13\xdf\xf3\x33\x3e\xe0\xc1\x3b\x5e\xf7\x01\x2b\x08\x0a\x41\x01\x5d\xb7\x7e\x32\xca\xa1\xea\xf5\x17\x3f\x43\xf7\xde\xdf\xf7\x03\x70\xd8\x58\x0d\x4b\x8f\x26\x6b\x78\xfe\x66\xa9\x7a\x85\xff\x8a\x02\x16\xee\xe4\x61\x00\xc9\x62\x15\x89\x7b\xb9\xac\xad\xa8\x0c\x93\x62\x19\xbb\x4f\x38\x63\xd7\x02\x10\x3a\x0f\xfb\xa5\xa8\xb1\x4a\xdc\xce\xfc\xfd\xcc\x9a\x19\x97\x4b\xe4\xa7\x16\xad\xd6\xad\xa2\x2d\x15\x64\xb9\x5c\xe4\x96\x43\xc5\xb1\xd6\x05\x9a\x0f\x01\x2a\x73\xce\xca\xc5\x55\xf0\x5e\x2d\xf2\xd4\x59\x79\x6a\x79\xb9\x58\xad\xa6\x8e\xfd\x2a\x92\x34\x68\xe6\x35\xf5\x5d\xa0\xa1\x66\x2f\x49\x81\x1e\xbf\x3e\x7d\x77\x42\x86\xb4\xdc\xbd\x9c\x53\xbb\xe9\x01\x8d\x32\x87\x04\x8e\xb7\x94\x83\x8b\x15\xe1\xa1\x0a\xdc\x50\x54\xde\x37\x98\xf1\xf4\xc7\x68\x67\x8e\x8b\xcf\x8a\xaa\x98\x68\xad\x01\x46\xa2\x22\xf0\x7b\x6c\x4f\x37\x8c\x20\x1d\x6a\xf3\xd4\xf1\xf9\x37\x6a\x8f\xae\xe0\xc5\x0d\xe8\x9b\xa4\xda\x31\xe9\x44\x6a\xb2\x07\x52\x93\xfd\x2e\x62\x43\xa3\x01\x4a\xdb\x6d\xc3\xdc\x09\x0f\x98\x5b\x67\x3e\x8d\xe6\x19\x6e\x9e\x7a\x24\xf7\x27\xfe\xfc\xd5\x52\x9a\xd3\xe7\xaf\xeb\xdc\x78\xf4\xfd\xef\x00\x00\x00\xff\xff\xec\x52\x54\x58\x36\x05\x00\x00")

func templatesLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLoginHtml,
		"templates/login.html",
	)
}

func templatesLoginHtml() (*asset, error) {
	bytes, err := templatesLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/login.html", size: 1334, mode: os.FileMode(436), modTime: time.Unix(1464039900, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPasswordHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\x4d\x6e\x83\x30\x10\x46\xf7\x3e\xc5\xd4\x52\x97\x05\x54\xa9\x1b\x4a\x38\x43\x6e\x10\x59\x99\x71\xb0\x04\xd8\xb2\x87\x00\x1a\xf9\xee\x21\xe4\x47\xca\xb7\x7c\x6f\xf1\x3d\x11\x24\xeb\x46\x02\x3d\x25\x8a\xa7\x60\x52\x9a\x7d\x44\x9d\xb3\x12\x61\x1a\x42\x6f\x78\x93\x1d\x19\xa4\xa8\xa1\xd8\xb8\x6a\x12\xaf\x3d\x01\xaf\x81\x0e\x9a\x69\xe1\xf2\x9c\x92\x6e\x15\x6c\x43\x77\x2d\xac\x8f\xc3\x8f\x75\xd4\x23\xc8\x0e\xef\x9b\x1d\x72\x57\xc3\x5f\xf5\xfd\xff\x66\x83\x89\x17\x37\xd6\x50\x81\x99\xd8\xc3\x6f\x15\x96\x87\xcc\xaa\x29\xf7\x93\x56\xa9\xe3\x33\xe9\x4b\x7d\x14\x59\xef\xf9\x55\x24\x42\x23\xe6\x7c\x0b\x00\x00\xff\xff\x8d\x61\x4f\xdb\xcc\x00\x00\x00")

func templatesPasswordHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPasswordHtml,
		"templates/password.html",
	)
}

func templatesPasswordHtml() (*asset, error) {
	bytes, err := templatesPasswordHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/password.html", size: 204, mode: os.FileMode(436), modTime: time.Unix(1464039900, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesProfileHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x51\x6f\xe2\x38\x10\x7e\xef\xaf\xf0\xb9\x3d\x01\x42\x84\xb4\x77\xf7\x02\x49\x5e\xda\xf2\x70\xd2\xb6\x5d\xb1\xdd\x97\xd5\x6a\x65\x88\x03\xee\x3a\x71\x64\x3b\x14\x14\xe5\xbf\xef\xd8\x4e\x48\x40\x5d\x84\xb6\x35\x12\x64\xc6\xdf\xcc\x7c\xf3\x79\x62\x51\x96\x31\x4d\x58\x46\x11\x2e\x14\x95\x3f\x72\x29\x12\xc6\x29\xae\xaa\x8b\xb2\xd4\x34\xcd\x39\xd1\xb0\xb7\xa6\x24\xa6\x12\x23\x0f\xfc\x17\x81\xd2\x3b\x4e\x91\xde\xe5\x34\xc4\x9a\x6e\xf5\x78\xa9\x14\x8e\x2e\x10\xac\x98\x6d\xbc\x44\xc8\x74\x94\x30\xca\x63\x54\x5a\xa7\x59\xaf\x2c\xd6\xeb\x09\xfa\xcf\xff\x7b\xba\xf7\xa5\x44\xae\x58\x36\x41\x3e\x22\x85\x16\xe8\xc6\xcf\xb7\x6e\xb3\xb2\xdf\x05\xf7\xa8\x94\x42\x8e\x52\xaa\x14\x59\x51\xf5\xa7\xe9\xcc\xe2\x4c\xe9\x91\x25\x3e\x41\x99\xc8\x68\xbb\x93\x93\x38\x66\xd9\x6a\xc4\x69\xa2\x21\xf8\x34\x03\xce\x3a\x24\x16\x42\x82\x2a\xa3\xa5\xe0\x42\x4e\xd0\xe5\x6c\xe6\xc3\x9a\x76\xb6\xb7\x23\xb5\x26\xb1\x78\x35\x9c\xcc\xe7\x3a\xdf\xbe\x01\x23\xcb\x9f\x2b\x29\x8a\x2c\xee\x64\xba\xbd\xb9\xfd\xa7\x85\x34\x7e\xdf\x3f\xae\x60\x09\x48\x12\xb3\x42\x4d\xd0\xbf\xdd\x86\x9d\x1a\xa3\x85\xd0\x5a\xa4\x13\x74\x7d\xa0\xc6\x61\xcf\xd7\x1d\xe1\x83\xb1\x15\x29\x32\xc7\xbc\x94\x2c\xd7\xdd\x73\x7e\x21\x1b\xe2\xbc\xf5\x71\x6f\x88\x44\x56\x21\x85\x42\xf4\xff\xfc\xf1\xc1\xcb\x89\x54\xb4\x5f\x96\xde\x5c\x4b\xa8\xa0\xbc\x19\x4c\xc3\xbd\x85\x54\xd5\xc0\x55\xb1\x5f\x2c\x41\x7d\x93\x5a\x24\xe8\xe5\x73\x41\xe5\x0e\xfd\x15\x86\xa8\x07\x32\xd8\x79\x8c\x7b\x83\x8e\xd2\x57\xfd\xa4\xc8\x96\x9a\x89\xac\xdf\x75\x37\x79\x1c\x05\x8f\xd3\x6c\xa5\xd7\x08\xd2\xf8\xc7\x28\xb3\x24\xd5\x85\xcc\xa6\x07\xfe\xea\xc0\x3a\x30\xf6\x67\x1e\x42\xf9\x5e\x50\x70\xb4\xe4\x44\xa9\x10\x1f\x8e\x04\x8e\x82\x71\xc1\xa3\xde\x60\xfa\xfb\x54\xf0\x46\xa0\xbe\xd1\x8a\x41\x32\x7f\x0a\x3f\x01\x3a\xe0\x0c\xae\xe1\xf0\x2d\xce\x4d\x15\x8f\xe4\x39\xcd\x62\xe0\xc1\x59\xd4\x1b\xba\xe0\x6f\xec\xbb\xf7\xc9\x01\x86\xbd\x60\x6c\x76\x8e\x58\x38\xed\x7a\x97\xdd\x88\x99\x79\x33\x07\x1e\x4c\xc0\xad\xe9\xa7\xef\xfa\xc1\x83\xb3\x85\xb9\xea\x63\xf3\x8a\xe3\x81\x97\x4b\x6a\x59\x35\x2c\x3b\x39\x9a\xb3\xb6\x13\x65\x47\xc6\x8c\x94\x89\x83\x9e\xf4\x5a\xc4\x21\x7e\x7a\x9c\x7f\x81\x39\xb2\xb0\x00\x2e\x8e\x46\xdf\xf6\xfe\xa8\xa7\xcc\x02\x38\x59\x50\x6e\x94\x0c\xed\x4d\x95\x91\x94\xe2\xe8\x8e\x29\xb8\xa1\x76\xe8\x01\xac\x09\x28\x60\x30\x9d\x18\x96\xe5\x85\x46\x0c\x6a\xc5\x0e\x68\x70\x18\x99\xd8\x23\x57\x3b\xe4\x18\x86\x9a\x17\x60\xd8\x19\x16\x12\xfa\xf2\x9e\xa1\xa0\x77\xd7\xe2\xab\x0a\xa3\xb1\xab\x13\x8c\x81\xf9\xfb\x9a\x78\xae\x9f\x4e\x36\xb0\x87\xd7\xec\x5b\xfb\x0c\xea\x4d\x85\x8f\xe0\x4d\x53\xc2\x38\x8e\xee\xcd\xcf\x49\xc6\x0e\x58\xd3\xad\x8d\x33\xb8\xda\xc4\x1f\x41\x34\x07\xd8\x2b\x5c\x90\x38\x7a\xaa\x9f\x4e\xd2\xdd\xc3\x6b\xc6\xad\xed\x48\xb7\xf6\x7b\x89\x6d\xa8\x64\xc9\xee\x69\x4f\xef\xab\xb5\xd1\x59\x2c\x8f\x62\x6b\xae\xc7\xde\x0f\x61\xec\x8a\xba\x54\xaa\x58\xa4\xac\x3d\xb3\x79\x6d\x1e\xe5\x0d\xc6\x26\x13\x3c\x74\xff\x3a\x24\x42\xe8\xe6\xaf\x43\x59\xc2\x6d\x51\x55\xbf\x02\x00\x00\xff\xff\xa9\x54\x51\x6b\x74\x08\x00\x00")

func templatesProfileHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesProfileHtml,
		"templates/profile.html",
	)
}

func templatesProfileHtml() (*asset, error) {
	bytes, err := templatesProfileHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/profile.html", size: 2164, mode: os.FileMode(436), modTime: time.Unix(1464039900, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRegisterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x95\x51\x6f\xda\x30\x10\x80\xdf\xfb\x2b\x3c\xb7\x13\x20\x44\x48\xbb\xed\x05\x92\xbc\x54\xe5\x61\xd2\xd6\x6e\x6c\x7b\x99\xa6\xc9\x10\x07\xdc\x39\x71\x64\x3b\x14\x14\xe5\xbf\xef\x6c\x27\x60\x50\x57\x55\x5d\x8d\x14\x72\xe7\xf3\xdd\x77\xe7\xb3\x53\xd7\x29\xcd\x58\x41\x11\xae\x14\x95\xbf\x25\x5d\x31\xa5\xa9\xc4\x4d\x73\x56\xd7\x9a\xe6\x25\x27\x1a\x26\xd7\x94\xa4\xa0\x45\x01\xe8\xcf\x22\xa5\x77\x9c\x22\xbd\x2b\x69\x8c\x35\xdd\xea\xf1\x52\x29\x9c\x9c\x21\x18\x29\xdb\x04\x99\x90\xf9\x28\x63\x94\xa7\xa8\xb6\x4a\x33\x1e\x58\xaa\xd7\x13\xf4\x21\x7c\x3b\xdd\xeb\x72\x22\x57\xac\x98\xa0\x10\x91\x4a\x0b\x74\x15\x96\x5b\x37\xd9\xd8\x67\xc5\x03\x2a\xa5\x90\xa3\x9c\x2a\x45\x56\x54\xbd\xd4\x9d\x19\x1c\xf2\x1a\x59\xf0\x09\x2a\x44\x41\x0f\x33\x25\x49\x53\x56\xac\x46\x9c\x66\x1a\x16\x3f\x4d\xc0\x99\x07\xb1\x10\x12\xaa\x32\x5a\x0a\x2e\xe4\x04\x9d\xcf\x66\x21\x8c\xa9\x37\xbd\x1d\xa9\x35\x49\xc5\x83\x61\x32\xbf\xcb\x72\xfb\x88\x19\x59\xfe\x59\x49\x51\x15\xa9\xe7\xe9\xfa\xea\xfa\xdd\xc1\xa4\xd3\x87\xe1\x69\x04\x0b\x20\x49\xca\x2a\x35\x41\xef\xfd\x84\x5d\x35\x46\x0b\xa1\xb5\xc8\x27\xe8\xf2\xa8\x1a\xc7\x39\x5f\x7a\x85\x8f\xc6\xb6\x48\x89\xd9\xe6\xa5\x64\xa5\xf6\xf7\xf9\x9e\x6c\x88\xd3\xb6\xdb\xbd\x21\x12\xd9\x0a\x29\x14\xa3\x8f\xf3\xdb\xcf\x41\x49\xa4\xa2\xfd\xba\x0e\xe6\x5a\x42\x04\x15\xcc\xa0\x1b\x6e\xac\x49\xd3\x0c\x5c\x14\xfb\x60\x19\xea\x1b\xd7\x22\x43\xf7\x5f\x2a\x2a\x77\xe8\x4d\x1c\xa3\x1e\x94\xc1\x36\x64\xda\x1b\x78\x95\xbe\xe8\x67\x55\xb1\xd4\x4c\x14\x7d\x5f\xdd\xf9\x71\x08\x01\xa7\xc5\x4a\xaf\x11\xb8\x09\x4f\xad\xcc\x90\x54\x57\xb2\x98\x1e\xe9\x9b\x23\xe9\x48\xd8\xef\x79\x0c\xe1\x7b\x51\xc5\xd1\x92\x13\xa5\x62\x7c\xdc\x12\x38\x89\xc6\x15\x4f\x7a\x83\xe9\xbf\x5d\xc1\x89\x40\x7d\x53\x2b\x06\xce\xc2\x29\xfc\x45\xe8\x88\x19\x54\xc3\xe1\x63\xcc\x5d\x94\x80\x94\x25\x2d\x52\xe0\xe0\x2c\xe9\x0d\xdd\xe2\x9f\xec\x57\xf0\xc9\x19\x0c\x7b\xd1\xd8\xcc\x9c\x50\xb8\xda\xf5\xce\xfd\x15\x33\x73\x32\x07\x01\x74\xc0\xb5\xc9\xa7\xef\xf2\xc1\x83\x67\x17\xe6\xa2\x8f\xcd\x11\xc7\x83\xa0\x94\xd4\x52\x75\x94\x9e\x8f\x6e\xaf\x6d\x47\xd9\x96\x31\x2d\x65\xd6\x41\x4e\x7a\x2d\xd2\x18\xdf\xdd\xce\xbf\x41\x1f\x59\xb3\x08\x2e\x8e\xae\xbe\x87\xfb\xa3\xed\x32\x6b\xc0\xc9\x82\x72\x53\xc9\xd8\x5e\x55\x05\xc9\x29\x4e\xbe\xb7\x6f\x13\xc8\xde\xcc\x7b\xf6\xac\x28\x2b\x8d\x58\xea\x99\x23\xf3\xf4\xe5\x43\x6b\x63\x68\x65\x5e\x81\xe0\x75\xee\xd7\xf6\x32\xec\x82\x34\x0d\x46\x63\x17\x21\x1a\x03\xef\x4b\xd0\x69\x4e\x18\xc7\xc9\x8d\xf9\x7b\x12\xda\x19\xb6\xc4\xad\xf0\x3c\x5c\xeb\xfb\x35\x58\x4b\x30\x7b\x80\x1b\x06\x27\x77\xed\xdb\x93\xc4\x7b\xf3\x16\xfa\x20\x3b\xee\x83\xfc\xbf\x60\x1b\x2a\x59\xb6\xbb\xdb\xe3\xfd\xb0\x32\x7a\x16\xe5\xc9\xda\x96\xf5\x54\xfb\x2a\xc4\x2e\xa8\x73\xa5\xaa\x45\xce\x0e\xdb\x36\x6f\xc5\x13\xbf\xd1\xd8\x78\x82\x17\xff\xdb\x9b\x09\xa1\xbb\x6f\x6f\x5d\xc3\x71\x6b\x9a\xbf\x01\x00\x00\xff\xff\xe3\x40\xc4\x3a\xb6\x07\x00\x00")

func templatesRegisterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesRegisterHtml,
		"templates/register.html",
	)
}

func templatesRegisterHtml() (*asset, error) {
	bytes, err := templatesRegisterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/register.html", size: 1974, mode: os.FileMode(436), modTime: time.Unix(1464039900, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVerifyHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x3d\x8b\xe3\x30\x10\xed\xfd\x2b\x06\xc1\x95\x91\xc3\x41\x38\x48\x6c\x17\x57\xa7\xba\xe3\x0e\xb6\x5a\x14\x6b\x6c\x0b\x64\xc9\x2b\x8d\xbd\x09\xc2\xff\x7d\xc7\x89\x93\x25\x45\x54\x18\xa3\xf7\x34\xef\x83\x49\x49\x63\x63\x1c\x82\x18\x23\x86\xf7\x09\x83\x69\x2e\x62\x9e\xb3\x94\x08\xfb\xc1\x2a\x62\xa8\x43\xa5\x31\x08\x90\x7c\x9f\x15\x91\x2e\x16\x81\x2e\x03\x96\x82\xf0\x4c\x79\x1d\xa3\xa8\x32\xe0\xa3\xcd\x24\x6f\x23\x36\x3d\xc6\xa8\x5a\x84\x74\x05\x96\xf3\x69\x34\x75\x7b\xd8\x6d\x7f\x1c\x1e\x77\xbd\x0a\xad\x71\x7b\xd8\x82\x1a\xc9\xc3\xcf\xed\x70\xfe\x06\x97\xe1\x1b\x65\x4d\xcb\x84\x1a\x1d\x61\xb8\x61\x73\x56\xe4\x57\x13\x15\xbb\x61\x49\xa8\xad\x8a\xb1\x14\xcf\xca\xab\xa5\x94\x4c\x03\xf2\xb7\xf7\x16\x95\x8b\xf2\xff\xc2\x31\xa8\x39\xc9\x5d\xe6\xcd\x8f\x01\x54\x5d\xfb\xd1\x11\x74\x2a\xc2\x09\xd1\xc1\xb4\x12\x65\x71\x0a\x90\x57\x4c\xe2\x00\xd6\x32\x08\x01\xb5\x09\x58\x13\x6a\x60\xcf\xd4\x21\x14\x0a\xba\x80\x4d\x29\x52\x92\x7f\x29\x18\xd7\x46\x79\xf4\x9c\xec\xdf\x9f\xe3\x3c\x8b\xea\xfa\x5f\xe4\xaa\x82\x61\xe9\xc4\x38\xd8\x41\xc4\xda\x3b\x1d\xe5\xc3\x47\xd1\x23\xf1\x1c\xa2\x61\x83\x1f\xa3\x99\x4a\xc1\x23\x03\xc6\x4e\x00\x33\x89\x0b\x28\xc5\xaf\x03\x8c\xc1\x96\x2f\x64\xd6\xc0\x68\x23\xbe\xca\xc7\x5f\xab\xc1\x79\x5a\x82\x3c\x32\xde\x1f\xba\xa5\x97\x22\xe7\x4e\xb9\xda\xa7\x15\x68\xbc\xa7\xfb\x0a\xac\xc4\xaf\x00\x00\x00\xff\xff\x9b\xad\x5b\x5d\x3b\x02\x00\x00")

func templatesVerifyHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesVerifyHtml,
		"templates/verify.html",
	)
}

func templatesVerifyHtml() (*asset, error) {
	bytes, err := templatesVerifyHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/verify.html", size: 571, mode: os.FileMode(436), modTime: time.Unix(1464039900, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVerify_emailHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\x41\x4e\xc4\x30\x0c\x45\xf7\x39\x85\x99\x35\xe4\x00\x6c\x07\x24\x56\x08\x01\xc3\xb6\x0a\xad\xdb\xb1\xea\x71\x90\xe3\x8e\xa8\x22\xdf\x9d\x76\x0a\x12\x42\x2c\x93\xaf\xf7\xbf\x5f\xad\x1d\xf6\x24\x08\xbb\x33\x2a\xf5\x73\x83\xa7\x44\xbc\x73\x0f\x0f\xc8\x9c\xa1\xd6\x78\x28\xa8\xf1\x8e\xca\x07\xa7\xf9\x31\x9d\xd0\xfd\x3a\x84\xd7\x63\x92\x11\xe6\x3c\x41\x9f\x15\x14\x07\x2a\xb6\x14\xc8\x00\xc9\x56\x68\x9f\xa5\xa7\x21\x16\x32\x6c\xe4\x02\xc5\x10\x9e\x18\x53\x41\xd8\x96\x56\x58\xe1\x32\x07\xa9\xeb\x14\x4b\x81\xf7\x19\x5a\xa6\x76\x5c\x7b\xec\x48\x05\x98\x64\xbc\xfd\xdb\x37\x29\xbb\x6f\x7f\x86\x9f\x16\xef\x85\xe9\x8c\xb2\xbc\x15\xe3\xc1\x88\xe3\x8b\xe9\x33\x2e\xf7\xb6\x08\x3f\xe4\xb4\x58\x34\xdf\x8e\x9a\x27\x5b\x8c\x6b\x9b\x3b\xf4\x1d\x6c\x86\x6f\x6b\x46\x6d\x32\xca\xb2\x5f\x03\xff\x65\x79\x15\x6e\xfe\xb7\x0a\xb5\xa2\x74\xee\x5f\x01\x00\x00\xff\xff\x7f\x37\x68\xbb\x48\x01\x00\x00")

func templatesVerify_emailHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesVerify_emailHtml,
		"templates/verify_email.html",
	)
}

func templatesVerify_emailHtml() (*asset, error) {
	bytes, err := templatesVerify_emailHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/verify_email.html", size: 328, mode: os.FileMode(436), modTime: time.Unix(1464039900, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"templates/login.html": templatesLoginHtml,
	"templates/password.html": templatesPasswordHtml,
	"templates/profile.html": templatesProfileHtml,
	"templates/register.html": templatesRegisterHtml,
	"templates/verify.html": templatesVerifyHtml,
	"templates/verify_email.html": templatesVerify_emailHtml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"templates": &bintree{nil, map[string]*bintree{
		"login.html": &bintree{templatesLoginHtml, map[string]*bintree{}},
		"password.html": &bintree{templatesPasswordHtml, map[string]*bintree{}},
		"profile.html": &bintree{templatesProfileHtml, map[string]*bintree{}},
		"register.html": &bintree{templatesRegisterHtml, map[string]*bintree{}},
		"verify.html": &bintree{templatesVerifyHtml, map[string]*bintree{}},
		"verify_email.html": &bintree{templatesVerify_emailHtml, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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
