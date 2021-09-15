// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// mysql/000001_create_teams.down.sql
// mysql/000001_create_teams.up.sql
// mysql/000002_create_team_members.down.sql
// mysql/000002_create_team_members.up.sql
// postgres/000001_create_teams.down.sql
// postgres/000001_create_teams.up.sql
// postgres/000002_create_team_members.down.sql
// postgres/000002_create_team_members.up.sql
package migrations

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _mysql000001_create_teamsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x91\x51\x6b\xbb\x30\x14\xc5\xdf\xf3\x29\xee\x9b\xfa\xe7\xcf\xd8\x9e\xa5\x63\xa9\x5e\xd7\x80\x9a\x62\x22\xeb\x5b\xc9\x6c\xc6\x04\x8d\x62\xd2\x6d\x1f\x7f\xb4\x4a\xd9\xda\xbd\x14\x0a\x83\xd1\x3c\x84\x40\xce\x39\x39\x37\x3f\x81\x12\x1e\xfa\x41\xf7\x6a\xd0\x1b\xe1\x94\xd3\xad\x36\x0e\x66\xe0\x0b\x4c\x31\x92\xc0\x12\x9f\x00\x00\x8c\xfb\x6e\x4d\x17\x11\x2f\x73\xe9\xff\x0b\x20\x29\x78\x06\x2c\x4f\x78\x91\x51\xc9\x78\xbe\x16\xd1\x02\x33\x7a\x13\xf1\xb4\xcc\x72\x71\xf0\x3d\x2d\xb0\x40\x70\xea\xb9\xd1\x6b\xa3\x5a\x0d\x33\xf0\xa4\x56\xad\xf5\x0e\x12\x9a\xc7\x93\xc0\x56\xaf\xba\x55\x30\x83\x98\x4a\x3a\xa7\x02\xfd\xe0\x9b\xaa\xea\x9a\x6d\x6b\x0e\x39\xb4\x69\xba\x77\xde\x6b\xc3\xcc\x5b\xed\xf4\x98\x18\xc0\x3d\xdc\xfe\xdf\x1f\x3d\x9a\x4a\x2c\x40\xd2\x79\x8a\xb0\x7f\x14\xe2\x82\x2f\x61\xec\x08\x47\xf6\xd0\x9b\x5c\xd3\xa8\x77\x1e\x09\x82\x90\x90\x65\x81\x4b\x5a\x20\xa8\xc6\xe9\x81\xbd\xe0\x47\x6d\x9d\x1d\xe7\x3f\xfd\xc3\x90\xe0\x0a\xa3\x52\x1e\xc9\x43\x12\x23\x4d\x53\x1e\x51\x89\xf0\x63\x60\x48\xc8\x1f\xa1\x92\x2a\xeb\x76\x59\xac\xea\x4c\xd9\x6f\xd4\xf9\x60\x4e\x13\xae\x6c\x2e\xc3\x26\xd6\xb6\x1a\xea\xde\xd5\x9d\x39\x17\xca\x17\xeb\x95\xc6\x65\x68\x3c\x0e\xdd\xb6\x8f\x3a\x63\xdd\xa0\x6a\xa3\x37\xe7\x22\x39\xf6\xff\x3a\x97\x7d\xb9\xb1\x2d\x4b\x00\x57\x4c\x48\x31\xf6\x0e\xc9\x67\x00\x00\x00\xff\xff\xf1\x20\x25\x99\x74\x06\x00\x00")

func mysql000001_create_teamsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000001_create_teamsDownSql,
		"mysql/000001_create_teams.down.sql",
	)
}

func mysql000001_create_teamsDownSql() (*asset, error) {
	bytes, err := mysql000001_create_teamsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000001_create_teams.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000001_create_teamsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\x51\x6f\x9b\x3e\x14\xc5\xdf\xf9\x14\xf7\x2d\xf0\xd7\x5f\x53\x9b\xb5\x55\xa5\x28\xd3\x5c\xb8\x69\xad\x11\xe8\xc0\x68\xed\x53\xe4\x04\x77\xb5\x04\x36\x02\xa7\x6b\xbf\xfd\x84\x21\x59\xd7\x50\x2d\x4a\x1f\x26\x4d\x7d\x41\x0a\xfe\xf9\xe8\xfa\x9e\xe3\x4b\xfc\x04\x09\x43\x60\xe4\x22\x44\xa0\x33\x88\x62\x06\x78\x43\x53\x96\x02\x13\xbc\x6c\xc0\x75\x00\x00\x68\x0e\x0f\xbc\x5e\xdd\xf3\xda\x1d\x9f\x79\x96\x8a\xb2\x30\xfc\xdf\x2e\xfa\xb5\xe0\x46\x10\x03\x4b\xf9\x5d\x2a\xe3\x8e\x8f\x3c\x08\x70\x46\xb2\xf0\x39\x95\x55\xf9\x1e\x54\x20\x0a\xb1\x07\x25\x9b\xaa\xe0\x4f\x11\x2f\xc5\xb6\xae\xb3\x93\x21\x72\x0f\x24\x10\xcd\xaa\x96\x95\x91\x5a\xfd\x3a\xe4\xe9\xe9\x10\x8a\x25\x97\xc5\x16\x3a\x1e\x9f\x0f\x41\xec\xa9\x12\x7f\x12\xf2\x75\x59\x71\xb5\xcf\x01\x48\x51\xe8\x1f\x22\x0f\x74\xc9\xa5\x6a\xc0\x88\x47\xd3\x2d\x50\xf5\x20\x8d\x78\x66\xcc\xc7\xf1\xd0\xfe\x74\x75\x2f\x4a\xf1\xc2\xbf\x5d\xec\x3a\xa1\x73\x92\xdc\xc2\x17\xbc\x05\x97\xe6\x5e\xef\x59\x44\xbf\x66\x68\x5f\xda\x52\xdd\xf6\xd9\xaf\xb5\x2f\x65\xfe\xb8\x30\x6d\x4c\x16\xd2\x56\xb3\x90\x39\xb8\x9b\xc2\x06\xb9\xb5\x4d\xc1\x82\x1b\x70\x37\x81\x18\xe4\x56\x36\x53\x96\xdb\xc4\x6b\x90\xcb\x6d\x5e\x2c\xb7\x89\xce\x20\xd7\xd8\x36\xd8\xfa\x36\x1d\xf1\x1c\x0f\x30\xba\xa4\x11\x4e\xa9\x52\x3a\xb8\xd8\xb6\xc5\xbf\x22\x49\x8a\x6c\xba\x36\x77\xe7\xe5\xf2\x64\xe2\x38\x29\x32\xf8\x5c\xd5\xa2\xe2\xb5\xc8\x53\xc3\x8d\x28\x85\x32\x30\x05\x37\xc5\x10\x7d\x06\x74\xd6\x5d\x94\xee\x69\x1b\xdf\x2d\xf8\x71\x16\x31\xf7\x3f\x0f\x66\x49\x3c\x07\x1a\xcd\xe2\x64\x4e\x18\x8d\xa3\x45\xea\x5f\xe1\x9c\x7c\xf0\xe3\x30\x9b\x47\xe9\x76\xdf\xb7\x2b\x4c\x10\x0c\x5f\x16\x62\xa1\xda\xa6\x4f\x61\x64\xaf\xe2\x68\x8b\x90\x28\xe8\x01\x7b\x2c\x0e\x53\x08\x08\x23\x17\x24\x45\xd7\xfb\x8d\x5a\xe9\x62\x5d\xaa\xad\x8e\x0d\x53\x5c\x09\xd5\x79\xd4\x29\x7a\xf0\x09\x8e\xba\x9e\x8d\xfa\xa2\x8f\x47\xfd\x6f\x12\x32\x4c\xfa\xe9\xd0\xcd\x03\x12\x04\xf0\x42\x06\x96\x5a\x17\x93\x91\xe3\x79\x13\xc7\xb9\x4e\xf0\x9a\x24\x08\xbc\x30\xa2\xa6\x77\x91\x36\xf8\x28\x1b\xd3\x74\xe7\xdf\xed\xe1\xc4\xc1\x1b\xf4\x33\xb6\xbb\x63\xe2\x04\x48\xc2\x30\xf6\xdb\x01\xf5\x9a\xec\xbf\xe3\x4d\xc8\x1b\xd3\x6a\xd1\x95\x56\xdd\xd5\x38\xd4\x9e\x5d\xa5\x7e\x9a\xbe\x7b\xf4\x46\x8f\x9e\x7d\x2a\x0e\x35\xe7\xb5\xaf\xcd\xbb\x37\x6f\xf4\xe6\xb2\xd6\xeb\xca\xd7\xaa\x31\x35\x97\x4a\xe4\x87\x1a\xf4\x52\x07\x8c\x54\x4f\xed\x5f\x91\xe3\xbf\xe7\xd1\xcf\x00\x00\x00\xff\xff\x18\x00\xd6\x8a\xa6\x09\x00\x00")

func mysql000001_create_teamsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000001_create_teamsUpSql,
		"mysql/000001_create_teams.up.sql",
	)
}

func mysql000001_create_teamsUpSql() (*asset, error) {
	bytes, err := mysql000001_create_teamsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000001_create_teams.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000002_create_team_membersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x91\xdf\x6b\xbb\x30\x14\xc5\xdf\xf3\x57\xdc\x37\xf5\xcb\x97\xb1\x3d\x4b\xc7\x52\xbd\xae\x82\x9a\x92\x44\xd6\xb7\x92\xb6\x77\xac\x60\x5c\xd1\x14\xf6\xe7\x8f\xaa\x73\xdd\x8f\xa7\xb5\x30\x18\xcd\x43\x08\xe4\x9c\x93\x93\xfb\x51\xa8\xe1\x6e\xd7\xd0\xce\x34\xb4\x51\xce\x38\xb2\x54\x3b\x98\x80\xaf\x30\xc3\x48\x43\x9a\xf8\x0c\x00\xa0\xdf\x0f\x6b\xb8\x88\x44\x59\x68\xff\x5f\x00\x89\x14\x39\xa4\x45\x22\x64\xce\x75\x2a\x8a\xa5\x8a\x66\x98\xf3\xab\x48\x64\x65\x5e\xa8\xd1\xf7\x30\x43\x89\xe0\xcc\xaa\xa2\x65\x6d\x2c\xc1\x04\x3c\x4d\xc6\xe6\x64\x57\xd4\xb4\xde\x28\xe4\x45\x3c\xc8\xda\xf5\x13\x59\x03\x13\x88\xb9\xe6\x53\xae\xd0\x0f\x3e\xa8\xd6\xcf\xd5\xde\xd6\x63\x9a\x3a\xc8\xa9\x6c\xa9\xe9\xc3\x02\xb8\x85\xeb\xff\xdd\xd1\xe3\x99\x46\x09\x9a\x4f\x33\x84\xa3\x57\x21\x96\x62\x0e\x7d\x55\x78\xf7\x87\xde\x60\x1b\x3e\x7b\xe3\xb1\x20\x08\x19\x9b\x4b\x9c\x73\x89\x60\x2a\x47\x4d\xfa\x88\x2f\xdb\xd6\xb5\xfd\x04\xbe\x4e\x31\x64\xb8\xc0\xa8\xd4\x9f\xe4\x21\x8b\x91\x67\x99\x88\xb8\x46\xf8\x36\x30\x64\xec\x0f\x72\xe1\x1b\xbb\xad\x4f\x01\xd3\x05\x5c\xc8\x9c\x9f\xcc\xfd\x9e\x5a\x77\x0a\x99\x2e\xe0\x42\xe6\x9c\x64\x62\xaa\xc8\x11\xff\x21\x96\x37\xf7\xaf\x33\xe9\x4a\xf5\x5d\xd3\x04\x70\x91\x2a\xad\x8e\x5b\x87\xec\x35\x00\x00\xff\xff\x11\xda\xf6\x5b\x82\x06\x00\x00")

func mysql000002_create_team_membersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000002_create_team_membersDownSql,
		"mysql/000002_create_team_members.down.sql",
	)
}

func mysql000002_create_team_membersDownSql() (*asset, error) {
	bytes, err := mysql000002_create_team_membersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000002_create_team_members.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql000002_create_team_membersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x5f\x6b\xdb\x30\x14\xc5\xdf\xfd\x29\xee\x5b\xac\x51\x46\x57\xca\x18\x84\x8c\xa9\xf6\x4d\x23\x66\xcb\x41\x92\x59\xfb\x64\x94\xf8\x76\x35\x58\x4e\xb1\x95\x91\x7d\xfb\xe1\x3f\xcd\x3a\xba\x3f\x0f\xeb\xd8\x28\xf3\x83\x41\xbe\x47\x47\x97\xf3\xbb\x58\x91\x42\x6e\x10\x0c\xbf\x48\x10\xc4\x12\x64\x66\x00\xaf\x84\x36\x1a\x0c\x59\x97\x92\xdb\x50\xdb\x41\x18\x00\xc0\xf0\x45\x94\xf0\xc9\xb6\xdb\x5b\xdb\x86\x67\xaf\xd9\xa0\x97\x79\x92\x9c\x0c\x82\xbc\xa3\xf6\xa7\x02\xb5\xab\xa9\x03\x4f\x07\x3f\xae\x63\xaa\xc9\x13\xf7\xb0\xa9\x3e\x56\x8d\x0f\xcf\x4e\xd9\x58\x58\x2b\x91\x72\x75\x0d\xef\xf1\x1a\xc2\xf1\xdc\x93\xc9\x7e\x52\xf4\x95\xaa\x3c\x14\x9e\xac\x73\x63\x9b\xc5\xbe\xa3\xb6\xa8\x4a\x08\x7f\xa9\x2c\x87\x83\x0b\xeb\x21\xbc\xef\x81\x05\x0c\x50\x5e\x0a\x89\x0b\xd1\x34\xbb\xf8\x02\x62\x5c\xf2\x3c\x31\x10\xad\xb8\xd2\x68\x16\x7b\x7f\xf3\xc6\x6d\xce\xe7\x41\xa0\xd1\xc0\xbb\xbb\x96\xee\x6c\x4b\xa5\xf6\xd6\x93\xa3\xc6\xc3\x02\x42\x8d\x09\x46\x06\xc4\x72\x8c\x6c\x7c\xf7\xcf\x54\x88\xb2\x5c\x9a\xf0\x05\x83\xa5\xca\x52\x10\x72\x99\xa9\x94\x1b\x91\xc9\x42\x47\x2b\x4c\xf9\xcb\x28\x4b\xf2\x54\xea\xe3\xbe\x0f\x2b\x54\x08\xde\x6e\x6a\x2a\x1a\xeb\x08\x16\x30\x7b\x80\x66\x76\x14\x72\x19\x4f\xb2\x6e\x7b\x4b\xce\xc2\x02\x62\x6e\xf8\x05\xd7\x18\xb2\x6f\x54\xdb\x5d\xbd\x77\xcd\xd1\x4d\xf7\x72\xea\x23\x1b\xcd\x18\xbc\x85\xd3\x31\xba\xd9\xd4\xf5\xab\xd9\xb4\xe6\x89\x41\x35\x8d\xcb\xc3\x01\xe1\x71\x0c\x5f\x7d\xc0\x57\xcd\xe7\x9e\xe7\x39\x9b\xcf\x02\xc6\xe6\x41\xb0\x56\xb8\xe6\x0a\xc1\xd6\x9e\x5a\x71\x23\x77\x1e\x0f\x55\xe7\xbb\x31\x87\xc7\x59\xce\x03\xbc\xc2\x28\x37\x8f\x77\xcc\x83\x18\x79\x92\x64\x51\x3f\xb8\x3f\xb2\x7d\x9e\x8c\x78\xe9\xaa\xe6\x29\x20\x0d\x46\xff\x29\xfd\x19\x4a\x97\x7b\xea\xfc\x53\x50\x1a\x8c\xfe\x09\x4a\xcf\x0b\xd3\xfd\x1f\xff\xf7\x18\x7d\xe7\xee\xfa\x6b\x80\xbe\x04\x00\x00\xff\xff\x0d\x78\xb7\x3b\xc7\x07\x00\x00")

func mysql000002_create_team_membersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql000002_create_team_membersUpSql,
		"mysql/000002_create_team_members.up.sql",
	)
}

func mysql000002_create_team_membersUpSql() (*asset, error) {
	bytes, err := mysql000002_create_team_membersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/000002_create_team_members.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000001_create_teamsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xd0\xb1\x0e\xc2\x20\x10\x06\xe0\x9d\xa7\xe0\x3d\x98\xaa\xc5\xa4\x49\x6d\x4d\x5b\x93\x6e\x84\xc0\x45\x2f\xa1\x40\x80\xaa\x8f\x6f\x2c\x83\x6e\xc5\xf9\xbe\xff\x72\xff\x55\xed\xc4\x07\x3a\x55\x87\x96\xd3\x04\x72\x89\xb4\x1e\xfa\x0b\x3d\xf6\xed\xf5\xdc\xd1\xe6\x44\xf9\xdc\x8c\xd3\x48\xa5\x31\xee\xe9\x3c\x58\xb4\x0f\x4c\xc0\x48\x69\xd0\xc8\x98\x3e\x73\x54\xce\xae\x5e\xcb\x7f\xb2\x1a\xa2\x0a\xe8\x13\x3a\x5b\x1e\xba\x05\xb7\x7a\xe5\x6c\x4c\x41\xa2\x05\xcd\x08\xd9\x60\xd3\xd5\x7c\xfe\x71\xa8\x5f\x62\xdb\x23\x72\x25\x81\x9a\xed\xc9\x5c\x40\xc8\xb4\x2b\x55\x80\x42\xa9\xc1\x40\x99\x8c\xea\x0e\x4b\xbe\x33\xd3\xfc\x8c\x2f\xdd\x18\x23\xef\x00\x00\x00\xff\xff\x38\xaf\xff\x77\xd4\x01\x00\x00")

func postgres000001_create_teamsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000001_create_teamsDownSql,
		"postgres/000001_create_teams.down.sql",
	)
}

func postgres000001_create_teamsDownSql() (*asset, error) {
	bytes, err := postgres000001_create_teamsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000001_create_teams.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000001_create_teamsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x4f\x6f\xe2\x30\x14\xc4\xef\xf9\x14\xef\x48\xa4\x3d\xb0\xec\x82\x56\xe2\x94\x05\x57\x8d\x0a\xa1\x0d\xa1\x82\x53\x64\xe2\x27\xfa\xa4\xc4\x8e\x62\xd3\x96\x6f\x5f\x81\x1d\x20\xfc\x91\xc2\x75\xc6\xfa\xbd\x64\x66\x46\x31\x0b\x12\x06\x49\xf0\x7f\xc2\x20\x7c\x82\x68\x96\x00\x5b\x86\xf3\x64\x0e\x06\x79\xa1\xa1\xe3\x01\x00\x90\x80\xf7\x20\x1e\x3d\x07\x71\xa7\x37\xf0\xe1\x35\x0e\xa7\x41\xbc\x82\x17\xb6\xfa\x75\xf0\xb3\x0a\xb9\x41\x6e\x60\x4d\x1b\x92\xc6\x8a\xdb\x52\x5c\x8b\x02\x73\xbc\x16\x49\x97\x39\xdf\x49\x5e\xe0\xf1\xce\xe0\xaf\x6f\xcd\xdb\xaa\x40\x9d\x55\x54\x1a\x52\xf2\xf4\x69\xfd\xbe\x73\xb1\xe0\x94\x1f\xf5\xdf\xbd\x7f\x4e\x37\xbb\x12\x6f\x3c\xcf\x54\x51\x72\x79\xe7\x3e\xcf\x73\xf5\x85\x42\xa8\x82\x93\xd4\x27\x68\xb7\xdb\x75\x2f\x48\x7e\x92\xc1\xb3\x8c\xfe\xf4\x9c\xa3\xb3\x0f\x2c\xb0\x99\x9e\x75\x16\x51\xf8\xb6\x60\x9d\xfd\x49\xdf\xf3\x87\x9e\xe7\x9a\x08\xa3\x31\x5b\x5e\x34\x41\xe2\x3b\x3d\xb4\x91\xda\x4b\x29\x09\x98\x45\x75\x41\xf5\x75\x7f\xd8\x8e\x61\x6b\x49\xb9\x39\x63\xd4\x55\xb5\x65\xd8\xbe\x9b\x8c\x7a\x03\x6d\x19\x76\x09\x4d\x46\xbd\x8e\xb6\x0c\x9b\x6f\x33\x8f\x3a\xf3\x7d\xa8\xc1\x24\x61\xb1\x5b\xb7\xb5\x83\xf1\x18\x46\xb3\xc9\x62\x1a\x5d\x40\x0f\x35\xab\x12\xa5\xcd\x13\xd6\x4a\xe5\xc8\xe5\xf0\x11\x46\xce\xb5\xd9\x3f\xa1\x4c\x49\x1b\xa9\x9b\xf9\x43\x94\x7b\xd3\x7e\x08\xb2\xa9\xd4\xb6\xcc\x94\xd4\xa6\xe2\x24\x51\x9c\xfe\xe7\x27\x00\x00\xff\xff\x07\xd8\x47\xb0\xf3\x03\x00\x00")

func postgres000001_create_teamsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000001_create_teamsUpSql,
		"postgres/000001_create_teams.up.sql",
	)
}

func postgres000001_create_teamsUpSql() (*asset, error) {
	bytes, err := postgres000001_create_teamsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000001_create_teams.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000002_create_team_membersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\xb1\x0a\xc2\x30\x14\x45\xf7\x7c\xc5\xfb\x8f\x4c\xd5\x46\x08\xa4\xad\xb4\x11\xba\x85\x68\x2e\x1a\xf0\x39\x34\x29\xf8\xf9\xa2\x19\xda\xd1\x76\x7e\xe7\x1c\xde\xad\x8c\x55\x3d\xd9\xea\x60\x14\x65\x78\x66\xf0\x15\x53\xa2\xba\xef\xce\x74\xec\xcc\xa5\x69\x49\x9f\x48\x8d\x7a\xb0\x03\xa5\xdb\x03\x8c\x39\x61\x92\x62\x8f\xe9\x03\xc7\xd7\x3e\xf5\x3e\x23\xe5\xad\x6a\xc0\x13\x19\x3e\x4b\x21\x7e\x80\x6e\x6b\x35\xae\xee\x31\xbc\xdd\xaa\xe2\xbe\xcb\x5c\x0c\xf2\x3f\xba\xd4\xdd\x92\x2f\x6f\x2d\x82\x85\xe7\xa6\xc0\x52\x7c\x02\x00\x00\xff\xff\xc2\x6e\xba\x19\x6a\x01\x00\x00")

func postgres000002_create_team_membersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000002_create_team_membersDownSql,
		"postgres/000002_create_team_members.down.sql",
	)
}

func postgres000002_create_team_membersDownSql() (*asset, error) {
	bytes, err := postgres000002_create_team_membersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000002_create_team_members.down.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres000002_create_team_membersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\xc1\x6b\x83\x30\x18\xc5\xef\xf9\x2b\xbe\xa3\x42\x4f\x83\xf5\xe2\x29\xd3\x8c\xc9\x62\x1c\x69\x3a\xda\x93\xc4\xe5\xa3\x0b\x18\x05\x93\xc2\xfe\xfc\xa1\xa6\xac\x38\x18\xac\xbd\x3e\x9f\xbf\xf7\xde\x97\x5c\x32\xaa\x18\x28\xfa\xc4\x19\x94\xcf\x20\x6a\x05\xec\x50\xee\xd4\x0e\x02\x6a\xe7\xd0\xb5\x38\x7a\x48\x08\x00\xcc\x8a\x35\xf0\x4e\x65\xfe\x42\x65\xf2\xb0\x4d\x67\xbf\xd8\x73\xbe\x99\x0d\x67\x8f\xe3\x9f\x86\x71\xe8\xd0\xff\x7c\x7f\xdc\xa6\x8b\x6e\xb0\xc3\x80\x3a\x40\x6b\x4f\xb6\x0f\x8b\xf8\x26\xcb\x8a\xca\x23\xbc\xb2\x23\x24\x4b\xf6\x26\x46\xa4\x24\xcd\x08\x89\xe5\x4b\x51\xb0\xc3\xaa\xbc\x35\x5f\xcd\xd5\x80\x66\xfa\xad\xb1\x06\x6a\x01\x0a\xb5\xab\x2e\xbb\x22\x2e\xfb\x0f\x6b\x29\xdb\xe8\xf0\x8b\x76\x99\x31\x95\xa3\x5c\x31\x19\x0f\x7b\x7d\x4a\x5a\x14\x90\xd7\x7c\x5f\x89\x55\x8c\xff\xf8\x44\x87\x53\x21\x68\x87\xa1\x43\xdd\x67\xb7\x42\xb4\x71\xb6\xbf\x9b\x72\x3a\xa3\x0f\x77\x50\x56\x8f\x9a\x91\xef\x00\x00\x00\xff\xff\xa5\xab\x10\xe7\x6d\x02\x00\x00")

func postgres000002_create_team_membersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres000002_create_team_membersUpSql,
		"postgres/000002_create_team_members.up.sql",
	)
}

func postgres000002_create_team_membersUpSql() (*asset, error) {
	bytes, err := postgres000002_create_team_membersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/000002_create_team_members.up.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"mysql/000001_create_teams.down.sql":           mysql000001_create_teamsDownSql,
	"mysql/000001_create_teams.up.sql":             mysql000001_create_teamsUpSql,
	"mysql/000002_create_team_members.down.sql":    mysql000002_create_team_membersDownSql,
	"mysql/000002_create_team_members.up.sql":      mysql000002_create_team_membersUpSql,
	"postgres/000001_create_teams.down.sql":        postgres000001_create_teamsDownSql,
	"postgres/000001_create_teams.up.sql":          postgres000001_create_teamsUpSql,
	"postgres/000002_create_team_members.down.sql": postgres000002_create_team_membersDownSql,
	"postgres/000002_create_team_members.up.sql":   postgres000002_create_team_membersUpSql,
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
	"mysql": &bintree{nil, map[string]*bintree{
		"000001_create_teams.down.sql":        &bintree{mysql000001_create_teamsDownSql, map[string]*bintree{}},
		"000001_create_teams.up.sql":          &bintree{mysql000001_create_teamsUpSql, map[string]*bintree{}},
		"000002_create_team_members.down.sql": &bintree{mysql000002_create_team_membersDownSql, map[string]*bintree{}},
		"000002_create_team_members.up.sql":   &bintree{mysql000002_create_team_membersUpSql, map[string]*bintree{}},
	}},
	"postgres": &bintree{nil, map[string]*bintree{
		"000001_create_teams.down.sql":        &bintree{postgres000001_create_teamsDownSql, map[string]*bintree{}},
		"000001_create_teams.up.sql":          &bintree{postgres000001_create_teamsUpSql, map[string]*bintree{}},
		"000002_create_team_members.down.sql": &bintree{postgres000002_create_team_membersDownSql, map[string]*bintree{}},
		"000002_create_team_members.up.sql":   &bintree{postgres000002_create_team_membersUpSql, map[string]*bintree{}},
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
