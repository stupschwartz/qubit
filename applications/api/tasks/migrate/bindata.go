// Code generated by go-bindata.
// sources:
// sql/1497233022_initial_schema.down.sql
// sql/1497233022_initial_schema.up.sql
// DO NOT EDIT!

package migrate

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

var __1497233022_initial_schemaDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\x2f\x48\x2d\x4a\x2c\xc9\x2f\x8a\xcf\xcc\x2b\x28\x2d\x29\xb6\xe6\xe2\xc2\x22\x89\x26\x9c\x99\x9b\x98\x9e\x8a\x4d\x2c\xbe\x38\xb5\xb0\x34\x35\x2f\x19\x5d\xb2\x38\x39\x35\x0f\x5d\xac\xa0\x28\x3f\x2b\x35\x19\xc3\xc6\xa2\xf4\xc4\xbc\xcc\xaa\xc4\x92\xcc\xfc\xbc\x62\x6b\x2e\x40\x00\x00\x00\xff\xff\x34\x65\x70\xe8\xa9\x00\x00\x00")

func _1497233022_initial_schemaDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1497233022_initial_schemaDownSql,
		"1497233022_initial_schema.down.sql",
	)
}

func _1497233022_initial_schemaDownSql() (*asset, error) {
	bytes, err := _1497233022_initial_schemaDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1497233022_initial_schema.down.sql", size: 169, mode: os.FileMode(420), modTime: time.Unix(1497709628, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1497233022_initial_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x93\x3d\x6f\xc2\x30\x10\x86\xf7\xfc\x8a\x13\x53\x22\x75\xa0\x43\xa7\x4e\x06\x19\x9a\x36\x35\x95\x49\x2b\x31\x21\x2b\x9c\x90\x2b\xc5\x4e\x1d\x23\xb5\xfd\xf5\x55\x12\x12\x70\x3e\x20\x03\x4b\xe7\x8b\x1f\x3f\xef\x1b\xdf\x9c\x53\x12\x53\x88\xc9\x2c\xa2\xa0\xcd\x5e\x28\xf9\x2b\xac\xd4\x2a\x07\xdf\x03\x00\x90\x3b\x98\x85\xcb\x35\xe5\x21\x89\xe0\x8d\x87\xaf\x84\x6f\xe0\x85\x6e\xee\xca\xe9\x44\x89\x14\x27\xf0\x41\xf8\xfc\x89\x70\xff\x7e\x3a\x0d\x80\xad\x62\x60\xef\x51\xe4\x05\x8f\x9e\xe7\xf0\x33\xa3\x3f\x31\xb1\xe3\xd0\xe7\x32\xdb\xea\xd3\x90\xc5\x0d\xfd\xfa\xfd\xd5\x17\x8b\x15\xa7\xe1\x92\x15\x5c\xf0\x5b\xcc\x00\x38\x5d\x50\x4e\xd9\x9c\xae\xdd\xf0\xbe\xdc\x05\x5d\xff\x3c\x41\x85\xe3\xec\x8f\x51\x6f\x25\x7e\xc2\x39\xce\x75\xa1\xfd\xba\x32\x15\x7b\xdc\xe6\xf8\x75\x40\x95\xfc\x37\xef\x71\xba\x6e\xc4\x5b\x59\x77\xa8\x8e\x7c\xab\xd6\xfe\x0c\x3a\x43\x23\xac\x36\xe3\x62\x94\x0f\x6b\xd0\x3e\xd1\xca\xe2\xb7\x3d\xd3\x6f\xdb\xdb\x9f\x0c\x9b\xf1\x43\x67\x7c\x3d\x7e\x26\x8c\x48\xd1\xa2\xc9\xe1\x79\xbd\x62\x97\xca\xa9\x5d\x9d\x4e\xaa\xcd\xb8\x5c\xc5\x56\xaa\xec\x30\x76\xf9\x9b\x33\x03\x9d\x94\xac\xc1\xa9\xbb\xf4\x27\x96\xbb\xf0\xf5\x2f\x2a\xb4\xfb\x5e\xc1\xf1\x8a\xe1\x43\x45\xd6\xbf\x00\x00\x00\xff\xff\xa5\x22\xf3\x1c\x41\x05\x00\x00")

func _1497233022_initial_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1497233022_initial_schemaUpSql,
		"1497233022_initial_schema.up.sql",
	)
}

func _1497233022_initial_schemaUpSql() (*asset, error) {
	bytes, err := _1497233022_initial_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1497233022_initial_schema.up.sql", size: 1345, mode: os.FileMode(420), modTime: time.Unix(1497830467, 0)}
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
	"1497233022_initial_schema.down.sql": _1497233022_initial_schemaDownSql,
	"1497233022_initial_schema.up.sql":   _1497233022_initial_schemaUpSql,
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
	"1497233022_initial_schema.down.sql": &bintree{_1497233022_initial_schemaDownSql, map[string]*bintree{}},
	"1497233022_initial_schema.up.sql":   &bintree{_1497233022_initial_schemaUpSql, map[string]*bintree{}},
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
