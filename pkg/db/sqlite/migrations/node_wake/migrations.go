// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../db/sqlite/migrations/node_wake/1616878455601.sql

package node_wake


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

var _bindataDbSqliteMigrationsNodewake1616878455601Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x41\x0a\xc2\x40\x0c\x45\xf7\x73\x8a\xbf\x54\xb4\x27\xe8\xd6\x2b\xb8\x1e\x62\x13\xca\xd0\x36\x19\xd2\x48\xed\xed\xa5\x16\x41\xec\x2e\xf0\x7e\x1e\xbc\xa6\xc1\x65\x2a\xbd\x53\x08\xee\x35\x75\x2e\xdb\x15\xf4\x18\x05\x6a\x2c\x79\xa1\x41\x66\x9c\x12\x00\x14\x46\xd1\x90\x5e\x1c\x6a\x01\x7d\x8e\x23\xaa\x97\x89\x7c\xc5\x20\xeb\xf5\x33\xda\x15\x9c\x29\xc0\x9b\xeb\xbb\xdc\x29\x9b\xca\x41\xb2\xa3\x89\xba\x4c\xcc\x2e\xf3\x8c\x90\x57\xfc\xe1\x6a\x8b\xb8\x70\x36\x3d\xfc\xa7\x73\x9b\x7e\x43\x6e\xb6\x68\x62\xb7\x7a\x08\x69\xdf\x01\x00\x00\xff\xff\x1a\xcc\x56\xf4\xf0\x00\x00\x00")

func bindataDbSqliteMigrationsNodewake1616878455601SqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataDbSqliteMigrationsNodewake1616878455601Sql,
		"../../db/sqlite/migrations/node_wake/1616878455601.sql",
	)
}



func bindataDbSqliteMigrationsNodewake1616878455601Sql() (*asset, error) {
	bytes, err := bindataDbSqliteMigrationsNodewake1616878455601SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "../../db/sqlite/migrations/node_wake/1616878455601.sql",
		size: 240,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1617034594, 0),
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
	"../../db/sqlite/migrations/node_wake/1616878455601.sql": bindataDbSqliteMigrationsNodewake1616878455601Sql,
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
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
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
	"..": {Func: nil, Children: map[string]*bintree{
		"..": {Func: nil, Children: map[string]*bintree{
			"db": {Func: nil, Children: map[string]*bintree{
				"sqlite": {Func: nil, Children: map[string]*bintree{
					"migrations": {Func: nil, Children: map[string]*bintree{
						"node_wake": {Func: nil, Children: map[string]*bintree{
							"1616878455601.sql": {Func: bindataDbSqliteMigrationsNodewake1616878455601Sql, Children: map[string]*bintree{}},
						}},
					}},
				}},
			}},
		}},
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
