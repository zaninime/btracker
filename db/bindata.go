// Code generated by go-bindata.
// sources:
// schema_files/init_v1.sql
// schema_files/init_v2.sql
// schema_files/v1_to_v2.sql
// DO NOT EDIT!

package db

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _init_v1Sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x51\x4f\xc2\x30\x14\x85\xdf\xf7\x2b\x6e\xfa\xc2\x48\x86\xd1\x67\x9f\xa6\xd4\x84\x38\x37\x33\x8a\x4a\x42\x42\xca\x76\x75\x8d\xa3\x6d\xba\x0e\xc4\x5f\x6f\x37\x21\x11\xc7\x08\x8f\xdb\xe9\x3d\xe7\xf4\xeb\xbd\x4f\x69\xc8\x28\xb0\xf0\x2e\xa2\x40\x74\xbd\x2a\x45\x46\xae\x88\x55\xc6\xa0\xb4\x04\x7c\x0f\x80\x14\xbc\x2a\x08\xac\x76\x16\x39\x3c\xa7\x93\xa7\x30\x9d\xc3\x23\x9d\x07\x8d\x96\xab\xad\x2c\x15\xcf\x31\x27\x20\xa4\xc5\x0f\x34\x10\x27\x0c\xe2\x59\x14\xc1\x98\x3e\x84\xb3\x88\xc1\xb5\x37\xbc\xf5\xbc\xd1\x08\x16\xa6\x96\x0b\xcf\xeb\x49\xd5\x88\x66\x1f\x29\xf2\x43\xe0\xc1\xac\x4d\xdb\xf7\x5a\xf6\xc8\x95\xe5\x16\xbb\x3d\x5a\x4d\xe8\x46\x40\x7b\xfc\x57\x2b\x63\x7b\x06\xce\xdd\xac\x3d\x50\xeb\xb3\x72\x89\xef\x7d\xd6\x25\xaf\xec\xb2\xd6\xb9\x6b\xeb\xa6\xad\x58\xa3\x6b\xbe\xd6\xb0\x15\xb6\x68\x3f\xe1\x5b\x49\xec\x72\x8c\x93\x57\x7f\xd8\x38\xfc\x79\x05\xf0\x1b\x58\xc1\x11\x9b\xe1\x65\xbc\x33\x25\x25\x66\x56\x28\x79\x9e\xfa\x69\x74\xf8\xa5\x85\xd9\x5d\x50\xbf\xaf\xb0\xb3\xbd\xb0\x68\x95\x15\xb8\xe6\xbf\x25\x3f\x71\x07\x1b\x6e\xb2\x82\x9b\xff\xcb\xb8\xe1\x65\x8d\x1d\xe2\xa7\x23\x26\xf1\x98\xbe\x41\x12\x77\xd7\xaf\x8b\x33\x38\x6c\xd6\xb1\xd3\x24\x9e\xd2\x94\x39\x27\x96\x9c\xe8\xfa\x12\x46\x33\x3a\x05\x7f\xb0\x41\x53\x39\xc4\x83\x00\x6e\xdc\xfc\x4f\x00\x00\x00\xff\xff\x35\xc1\x50\xce\x71\x03\x00\x00"

func init_v1SqlBytes() ([]byte, error) {
	return bindataRead(
		_init_v1Sql,
		"init_v1.sql",
	)
}

func init_v1Sql() (*asset, error) {
	bytes, err := init_v1SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init_v1.sql", size: 881, mode: os.FileMode(420), modTime: time.Unix(1457466849, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _init_v2Sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x51\x6f\x82\x30\x14\x85\xdf\xf9\x15\x37\x7d\x11\x13\x5c\x96\xbd\xee\x89\xcd\x2e\x31\x63\xb0\x20\x6e\x33\x31\x31\x15\xee\xa4\x19\xb6\x4d\x29\x3a\xf7\xeb\x57\x9c\x26\x32\xc4\xf0\x08\x87\x7b\xce\xc7\xb9\xf7\x31\xa6\x7e\x42\x21\xf1\x1f\x02\x0a\x44\x55\xab\x82\xa7\xe4\x86\x18\xa9\x35\x0a\x43\xc0\x75\x00\x48\xce\xca\x9c\xc0\x6a\x6f\x90\xc1\x6b\x3c\x79\xf1\xe3\x39\x3c\xd3\xb9\x57\x6b\x99\xdc\x89\x42\xb2\x0c\x33\x02\x5c\x18\x5c\xa3\x86\x30\x4a\x20\x9c\x05\x01\x8c\xe9\x93\x3f\x0b\x12\xb8\x75\x86\xf7\x8e\x33\x1a\xc1\x42\x57\x62\xe1\x38\x1d\xa9\x0a\x51\x1f\x23\x79\x76\x0a\x3c\x99\x1d\xd2\x8e\x5c\xcb\x0e\xb9\x34\xcc\x60\x9b\xe3\xa0\x71\x55\x0b\x68\x9a\x6f\x95\xd4\xa6\x63\xe0\xfc\xcf\x56\x7c\x6d\xbf\x69\xea\x95\xba\xa6\x16\xf8\x69\x3a\x14\x56\x9a\x65\xa5\x32\x8b\x6a\x67\x0d\xdf\xa0\xc5\xde\x28\xd8\x71\x93\x1f\x1e\xe1\x47\x0a\x6c\x97\x18\x46\xef\xee\xb0\x76\x38\x5b\x01\xb8\x75\x53\x5e\xa3\x98\x61\xbf\xb2\x53\x29\x04\xa6\x86\x4b\x71\xbd\xf2\xcb\xbd\xe1\xb7\xe2\x7a\xdf\x03\xbf\x0b\xd8\xda\xf6\x04\x2d\xd3\x1c\x37\xec\x0f\xf2\x0b\xf7\xb0\x65\x3a\xcd\x99\xfe\x7f\x89\x5b\x56\x54\xd8\xda\xe4\xe5\x88\x49\x38\xa6\x1f\x10\x85\xed\xdb\x6b\xd7\xe9\x9d\xce\xaa\xe9\x34\x09\xa7\x34\x4e\xac\x53\x12\x5d\x60\x7d\xf3\x83\x19\x9d\x82\x3b\xd8\xa2\x2e\x6d\xc5\x03\x0f\xee\xec\xfc\x6f\x00\x00\x00\xff\xff\xde\x02\x9e\x2c\x6e\x03\x00\x00"

func init_v2SqlBytes() ([]byte, error) {
	return bindataRead(
		_init_v2Sql,
		"init_v2.sql",
	)
}

func init_v2Sql() (*asset, error) {
	bytes, err := init_v2SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init_v2.sql", size: 878, mode: os.FileMode(420), modTime: time.Unix(1457475452, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _v1_to_v2Sql = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x50\x2a\x28\x4d\xca\xc9\x4c\x56\xd2\x53\x2a\x48\x4d\x2d\x52\x52\x80\xc8\x39\xfb\xfb\x84\xfa\xfa\x29\x28\xa5\xe4\x97\xe7\xe5\xe4\x27\xa6\xa4\xa6\x28\x29\x84\x44\x06\xb8\x2a\x24\x65\xa6\x67\xe6\x95\xe8\xa0\x29\x2b\x2d\x20\x42\x51\x4e\x6a\x5a\x09\x8a\x02\x6b\x2e\x2e\x5d\x5d\x85\x98\xa2\xd2\xbc\x18\x2e\xae\xd0\x00\x17\xc7\x10\x64\xe7\x14\x27\x67\xa4\xe6\x26\x2a\x29\x04\xbb\x86\x28\x28\x95\x25\xe6\x94\xa6\x2a\xd9\x1a\x29\x84\x7b\xb8\x06\x01\x55\x65\xa7\x56\x2a\xd9\xaa\x97\xa5\x16\x15\x67\xe6\xe7\xa9\x5b\x73\x01\x02\x00\x00\xff\xff\xaf\x47\xac\xc7\xd2\x00\x00\x00"

func v1_to_v2SqlBytes() ([]byte, error) {
	return bindataRead(
		_v1_to_v2Sql,
		"v1_to_v2.sql",
	)
}

func v1_to_v2Sql() (*asset, error) {
	bytes, err := v1_to_v2SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1_to_v2.sql", size: 210, mode: os.FileMode(420), modTime: time.Unix(1457473616, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"init_v1.sql": init_v1Sql,
	"init_v2.sql": init_v2Sql,
	"v1_to_v2.sql": v1_to_v2Sql,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"init_v1.sql": &bintree{init_v1Sql, map[string]*bintree{
	}},
	"init_v2.sql": &bintree{init_v2Sql, map[string]*bintree{
	}},
	"v1_to_v2.sql": &bintree{v1_to_v2Sql, map[string]*bintree{
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

