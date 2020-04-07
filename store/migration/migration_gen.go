// Code generated by go-bindata.
// sources:
// sqlite3/001_init.sql
// sqlite3/002_org.sql
// sqlite3/003_drop_emails.sql
// sqlite3/004_limit_users.sql
// sqlite3/005_oauth_scope.sql
// sqlite3/006_add_orgs_table.sql
// sqlite3/007_add_slack_urls.sql
// mysql/001_init.sql
// mysql/002_org.sql
// mysql/003_drop_emails.sql
// mysql/004_limit_users.sql
// mysql/005_oauth_scope.sql
// mysql/006_add_orgs_table.sql
// mysql/007_add_slack_urls.sql
// postgres/001_init.sql
// postgres/002_org.sql
// postgres/003_drop_emails.sql
// postgres/004_limit_users.sql
// postgres/005_oauth_scope.sql
// postgres/006_add_orgs_table.sql
// postgres/007_add_slack_urls.sql
// DO NOT EDIT!

package migration

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

var _sqlite3001_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x8e\x82\x30\x10\x86\xef\x7d\x8a\x39\x6a\x56\x9f\x80\x13\xca\xec\xa6\x59\x6d\xdd\x5a\x12\x3c\x99\x66\xb7\x21\x8d\x40\x49\x41\xdd\xc7\xdf\xd0\x14\x94\x1e\x36\x72\x6a\xbe\xf0\x4f\xe7\x9b\xce\x7a\x0d\x6f\xb5\x29\x9d\xea\x35\xe4\x2d\x21\x5b\x81\xa9\x44\x90\xe9\x66\x87\x40\xdf\x81\x71\x09\x58\xd0\xa3\x3c\xc2\xb5\xd3\xae\x83\x05\xf1\x87\xb3\xf9\x01\xff\x51\x26\xf1\x03\x05\x1c\x04\xdd\xa7\xe2\x04\x9f\x78\x82\x34\x97\x9c\xb2\xad\xc0\x3d\x32\x49\x56\xfe\xff\xca\x96\xa6\x01\x00\x89\xc5\x88\x7a\x7b\xd1\x11\xd2\xb5\x32\xd5\x1c\xa9\x9b\xea\x95\x9b\xa1\x4e\x7f\x3b\xdd\x07\x44\x56\x39\xa3\x5f\x39\x2e\x1e\xd7\x2c\xc9\x32\xf9\x57\xc5\xe9\xd6\x7a\x95\xe1\x30\xa9\xbc\xe2\xe2\x03\xd3\x00\x42\x20\x60\x7b\x6f\xb4\x83\xa9\x7b\xcf\x1a\x55\x6b\x88\x58\x57\x5d\xcb\x98\x55\xa6\xb9\xc4\xac\x75\xe6\x36\xbc\x0b\x6c\x38\xdf\x61\xca\xc6\x78\xb0\x8f\xf4\xa7\xd2\x33\x7b\xca\x32\x2c\x22\x7b\xf3\x7b\x9e\xf5\xcb\xd9\x38\x90\x07\x5e\x26\xaf\x54\x18\x07\x11\x55\x08\x78\x68\xe3\x79\xbf\x32\x7b\x6f\x08\xc9\x04\x3f\x84\x47\xf1\x99\xe4\x99\xf8\x1d\x4b\xc8\x5f\x00\x00\x00\xff\xff\x54\x85\x0e\xca\x96\x02\x00\x00")

func sqlite3001_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite3001_initSql,
		"sqlite3/001_init.sql",
	)
}

func sqlite3001_initSql() (*asset, error) {
	bytes, err := sqlite3001_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/001_init.sql", size: 662, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite3002_orgSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\x2d\xc8\x2f\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x03\x0b\xc4\xe7\x17\xa5\x2b\x38\xf9\xfb\xfb\xb8\x3a\xfa\x29\xb8\xb8\xba\x39\x86\xfa\x84\x28\xb8\x39\xfa\x04\xbb\x5a\x73\x01\x02\x00\x00\xff\xff\x4d\x6f\x52\x6d\x4d\x00\x00\x00")

func sqlite3002_orgSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite3002_orgSql,
		"sqlite3/002_org.sql",
	)
}

func sqlite3002_orgSql() (*asset, error) {
	bytes, err := sqlite3002_orgSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/002_org.sql", size: 77, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite3003_drop_emailsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xc1\x4e\xc3\x30\x10\x44\xef\xfb\x15\x7b\x6b\x2b\xd2\x2f\xc8\xc9\x24\x1b\x64\x91\xd8\x65\xb3\x96\xda\x53\x15\x81\x55\x45\x34\x4d\x95\x04\xf8\x7d\x84\x63\x28\x41\xf8\x34\x7e\xda\xf5\x8c\x67\xbb\xc5\xbb\xae\x3d\x0d\xcd\xe4\xd1\x5d\x01\x54\x29\xc4\x28\xea\xbe\x24\x7c\x1b\xfd\x30\x22\x93\x51\x15\xa1\x58\x9c\x7c\x77\x3d\x06\x98\x02\x64\x4c\x4a\x28\x4e\xea\x02\x8d\x15\xa4\xbd\xae\xa5\x8e\x7b\x6b\x08\xe2\xd8\xbe\x60\x38\xda\x08\x3d\x10\xe3\x8e\x75\xa5\xf8\x80\x8f\x74\x40\xe5\xc4\x6a\x93\x31\x55\x64\x04\x92\x30\x7f\xee\x4f\xed\x05\x11\x85\xf6\xdf\x68\xea\x5f\xfd\x1f\xd4\xbc\x37\x53\x33\x2c\xd0\xe8\x9f\x07\x3f\x45\x04\x89\x33\xfa\xc9\xd1\xfa\xf6\xe6\x06\x36\x29\x80\x36\x35\xb1\x7c\xa5\xb1\x73\x50\xa8\xa9\xa4\x4c\x7e\xd2\x26\x78\x5b\x89\x3a\xf8\x47\x3d\x1b\xc7\xcb\x6c\x09\x05\xdb\x0a\x96\xf5\xe4\x6c\x77\xb1\x9c\x05\xff\xdd\x77\xde\x7f\x5c\xfe\x6b\x5c\xe5\x39\x66\xb6\x74\x95\x99\x5d\x7c\xd7\xb4\xe7\xf0\x2d\xcc\xa9\x50\xae\x14\x5c\xad\xd2\xcf\x00\x00\x00\xff\xff\x71\xe7\x23\x15\xba\x01\x00\x00")

func sqlite3003_drop_emailsSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite3003_drop_emailsSql,
		"sqlite3/003_drop_emails.sql",
	)
}

func sqlite3003_drop_emailsSql() (*asset, error) {
	bytes, err := sqlite3003_drop_emailsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/003_drop_emails.sql", size: 442, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite3004_limit_usersSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\xb1\x0a\x83\x30\x10\xc6\xf1\xfd\x9e\xe2\x1b\x95\xea\x13\x64\xb2\xf5\x0a\x81\xa2\xad\x5e\xc0\xad\x74\x90\x10\x50\x53\xa2\xa5\xaf\x5f\x1a\x1d\x1c\x3a\xf4\xc6\x3b\xf8\xdd\x3f\xcf\x71\x18\x9d\x0d\x8f\xa5\x87\x79\x12\x9d\x1a\x2e\x84\x21\xc5\xf1\xc2\xd0\x67\x54\xb5\x80\x3b\xdd\x4a\x8b\xc1\x8d\x6e\xb9\xbf\xe6\x3e\xcc\x48\x08\x83\xb7\x6e\x42\x1c\xe1\x4e\x28\x33\x95\xbe\x19\x4e\xe2\x3e\xa5\x54\xfd\x81\xf9\x60\xa3\x05\x1f\xec\xaa\x00\x19\x36\xc8\x07\xbb\x32\xfb\xc6\xd2\xbf\x27\xa2\xb2\xa9\xaf\x1b\xbb\xab\x52\x3f\x0e\xdf\x0f\x8a\x3e\x01\x00\x00\xff\xff\x23\xb0\x80\x6b\xe6\x00\x00\x00")

func sqlite3004_limit_usersSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite3004_limit_usersSql,
		"sqlite3/004_limit_users.sql",
	)
}

func sqlite3004_limit_usersSql() (*asset, error) {
	bytes, err := sqlite3004_limit_usersSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/004_limit_users.sql", size: 230, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite3005_oauth_scopeSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x03\x0b\xc4\x17\x27\xe7\x17\xa4\x16\x2b\x84\xb8\x46\x84\x28\xb8\xb8\xba\x39\x86\xfa\x84\x28\xa8\xab\x5b\x73\x01\x02\x00\x00\xff\xff\x49\xec\x90\x98\x4a\x00\x00\x00")

func sqlite3005_oauth_scopeSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite3005_oauth_scopeSql,
		"sqlite3/005_oauth_scope.sql",
	)
}

func sqlite3005_oauth_scopeSql() (*asset, error) {
	bytes, err := sqlite3005_oauth_scopeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/005_oauth_scope.sql", size: 74, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite3006_add_orgs_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x4e\xc3\x30\x10\x44\xef\xfe\x8a\x39\x26\x82\x4a\x50\xd1\x53\x4e\x6e\xb2\x80\x45\x6b\x97\xad\x83\xda\x53\x85\xc0\xaa\x2c\x20\xa9\x9c\x42\xf9\x7c\x64\xe4\xa4\x05\x2e\xf8\x64\xcd\xe8\xed\xee\xcc\x68\x84\xb3\x37\xbf\x0d\x8f\x7b\x87\x7a\x27\x44\xc9\x24\x2d\xc1\xca\xe9\x8c\xa0\xae\xa1\x8d\x05\xad\xd4\xd2\x2e\xd1\x86\x6d\x87\x4c\x20\x7e\x36\xfe\x19\xdf\x4f\x69\x4b\x37\xc4\x58\xb0\x9a\x4b\x5e\xe3\x8e\xd6\x90\xb5\x35\x4a\x97\x4c\x73\xd2\xf6\x3c\x01\xef\x9d\x0b\x91\x4a\x40\x2f\xb7\x87\xc6\x05\x00\x0f\x92\xcb\x5b\xc9\xd9\x78\x32\xc9\x7b\xef\xd5\x37\x2f\x38\xf1\x2e\x2f\xc6\x57\x83\xb9\x0b\xfe\x23\x1e\x3d\x35\x66\x46\x52\xf7\x72\xe7\x9e\x82\xdb\xff\x9d\x57\x6b\x75\x5f\x13\xb2\x61\x67\x2e\xf2\x62\x88\xab\x74\x45\xab\x5f\x71\xfd\xe7\xe6\x78\xa0\xd1\x29\xfe\x91\x2f\xfe\x01\xf7\xa1\x7f\xe0\x49\x8c\xeb\x4f\xdb\xaf\xda\x43\x23\x44\xc5\x66\x91\xda\x8f\x44\x21\xbe\x02\x00\x00\xff\xff\x7b\x6c\xa2\x6a\xa1\x01\x00\x00")

func sqlite3006_add_orgs_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite3006_add_orgs_tableSql,
		"sqlite3/006_add_orgs_table.sql",
	)
}

func sqlite3006_add_orgs_tableSql() (*asset, error) {
	bytes, err := sqlite3006_add_orgs_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/006_add_orgs_table.sql", size: 417, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sqlite3007_add_slack_urlsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x31\x4e\xc5\x40\x0c\x44\x7b\x9f\x62\xca\xff\xc5\xcf\x09\x52\x81\x42\x87\x04\x42\x50\x47\x4b\x30\xc1\xca\xae\x37\x78\xbd\x22\xb9\x3d\x4a\x0a\x94\x22\x9d\xf5\xac\x79\x9a\x69\x1a\xdc\x25\x19\x2d\x38\xe3\x7d\x26\x1a\x8c\xb7\xd3\xc3\x47\x64\xc8\x17\x34\x3b\x78\x91\xe2\x05\x25\x86\x61\xea\xab\xc5\x72\x21\x40\x3e\x21\xea\x3c\xb2\x61\x36\x49\xc1\x56\x4c\xbc\x22\x54\xcf\xa2\x83\x71\x62\xf5\x1b\x01\xdf\xb9\x78\xaf\x21\x31\x9c\x17\xdf\x7d\x5a\x63\xdc\x5e\xb5\xb0\x9d\x50\x8b\x27\x50\xe5\xa7\xf2\xe5\x5f\x76\xdb\xc3\x57\xba\xb6\x44\xc7\x09\x5d\xfe\x55\xa2\xee\xf5\xf9\x05\x6f\xf7\x0f\x4f\x8f\x87\xd2\x2d\xfd\x05\x00\x00\xff\xff\x7c\x1e\xe0\x1f\xec\x00\x00\x00")

func sqlite3007_add_slack_urlsSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlite3007_add_slack_urlsSql,
		"sqlite3/007_add_slack_urls.sql",
	)
}

func sqlite3007_add_slack_urlsSql() (*asset, error) {
	bytes, err := sqlite3007_add_slack_urlsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sqlite3/007_add_slack_urls.sql", size: 236, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql001_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x6f\x82\x30\x18\x86\xef\xfd\x15\xdf\x11\xb2\x99\x6c\x66\x9e\x38\x55\xf9\xb6\x35\xd3\xe2\x6a\x59\xf4\x64\x9a\xad\x31\x8d\x08\xa6\xa0\xfe\xfd\x85\x5a\x81\x6d\xb2\xc8\xa9\xe9\xc3\x5b\x78\x9f\x7e\x83\x01\xdc\xed\xcc\xc6\xaa\x4a\x43\xba\x27\x64\x22\x90\x4a\x04\x49\xc7\x53\x04\xf6\x0c\x3c\x91\x80\x4b\xb6\x90\x0b\x38\x94\xda\x96\x10\x10\xb7\x58\x9b\x2f\x70\x0f\xe3\x12\x5f\x50\xc0\x5c\xb0\x19\x15\x2b\x78\xc3\x15\xd0\x54\x26\x6b\xc6\x27\x02\x67\xc8\x25\xb9\x77\x81\xac\xd8\x98\x1c\x00\x3e\xa8\x98\xbc\x52\x11\x0c\x47\xa3\xd0\xa3\xaa\xd8\xea\x1e\xa4\x77\xca\x64\xd7\x91\x3a\xaa\x4a\xd9\x16\x3d\x3e\x0c\x9f\x2e\xac\xd4\x9f\x56\x57\xbf\x63\x29\x67\xef\x29\x06\xed\xef\x84\x24\x8c\xfe\xed\x6c\xf5\xbe\x70\x9d\xeb\x45\xd3\xf9\xa6\xd2\x2e\xd1\xa8\xf2\x09\xbf\x5d\x9c\x72\x6d\xe1\x4f\x2d\xc7\x72\xb5\xd3\xd0\xc3\xca\xec\xb0\xe9\x63\x99\xc9\xb7\x3f\x98\xf7\xe1\xe0\xde\x9a\x63\x7d\xc5\x30\x4e\x92\x29\x52\x7e\x39\xcf\x6b\xba\xee\xa9\xf9\xe4\x59\x13\x9d\x4a\x14\xde\xd2\xd9\x0b\x8d\x63\x60\x3c\xc6\x25\x04\x6d\xad\x30\xba\xe1\x4d\xef\xa5\x3e\xb6\x3b\x81\x71\x71\xca\x09\x89\x45\x32\xef\xa6\xa3\xee\x8e\x9b\xc2\x88\x7c\x07\x00\x00\xff\xff\x77\x0d\xa2\x03\xb8\x02\x00\x00")

func mysql001_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql001_initSql,
		"mysql/001_init.sql",
	)
}

func mysql001_initSql() (*asset, error) {
	bytes, err := mysql001_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/001_init.sql", size: 696, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql002_orgSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\x2d\xc8\x2f\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x03\x0b\xc4\xe7\x17\xa5\x2b\x38\xf9\xfb\xfb\xb8\x3a\xfa\x29\xb8\xb8\xba\x39\x86\xfa\x84\x28\xb8\x39\xfa\x04\xbb\x5a\x73\x71\x21\x9b\xe6\x92\x5f\x9e\x87\xcd\x3c\x97\x20\xff\x00\x74\x03\xad\x01\x01\x00\x00\xff\xff\x25\x2b\x07\x70\x87\x00\x00\x00")

func mysql002_orgSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql002_orgSql,
		"mysql/002_org.sql",
	)
}

func mysql002_orgSql() (*asset, error) {
	bytes, err := mysql002_orgSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/002_org.sql", size: 135, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql003_drop_emailsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x03\x8b\xc4\xa7\xe6\x26\x66\xe6\x58\x73\x71\x21\x6b\x75\xc9\x2f\xcf\xc3\xa6\xd9\xd1\xc5\x05\x53\xaf\x42\x98\x63\x90\xb3\x87\x63\x90\x86\x91\xa9\xa9\xa6\x82\x8b\xab\x9b\x63\xa8\x4f\x88\x82\xba\xba\x35\x20\x00\x00\xff\xff\xe3\x75\x50\xf8\x8d\x00\x00\x00")

func mysql003_drop_emailsSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql003_drop_emailsSql,
		"mysql/003_drop_emails.sql",
	)
}

func mysql003_drop_emailsSql() (*asset, error) {
	bytes, err := mysql003_drop_emailsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/003_drop_emails.sql", size: 141, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql004_limit_usersSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\xc9\xcc\xcd\x2c\x89\x2f\x2d\x4e\x2d\x2a\x56\xd0\xe0\x52\x50\xc8\xc9\x4f\xcf\xcc\x53\x00\x83\x30\xc7\x20\x67\x0f\xc7\x20\x0d\x23\x53\x53\x4d\xb0\x16\xbf\x50\x1f\x1f\x2e\x9d\x50\x3f\xcf\xc0\x50\x57\x0d\xb0\x42\x4d\x2e\x4d\x6b\x22\x8c\xcf\x2f\x4a\x87\x98\x9e\x5f\x94\x8e\xc3\x58\x05\x05\x1d\x05\xa8\xc9\xf9\x45\xe9\x10\x73\x91\xbd\xe1\x92\x5f\x9e\xc7\xc5\xe5\x12\xe4\x1f\x00\xb5\x07\xc9\xe1\xd6\x58\x24\x40\x56\x5a\x73\x01\x02\x00\x00\xff\xff\x21\x06\x25\x0f\x09\x01\x00\x00")

func mysql004_limit_usersSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql004_limit_usersSql,
		"mysql/004_limit_users.sql",
	)
}

func mysql004_limit_usersSql() (*asset, error) {
	bytes, err := mysql004_limit_usersSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/004_limit_users.sql", size: 265, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql005_oauth_scopeSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x03\x0b\xc4\x17\x27\xe7\x17\xa4\x16\x2b\x84\x39\x06\x39\x7b\x38\x06\x69\x18\x99\x9a\x6a\x2a\xb8\xb8\xba\x39\x86\xfa\x84\x28\xa8\xab\x5b\x73\x71\x21\x1b\xe9\x92\x5f\x9e\x87\xcd\x50\x97\x20\xff\x00\x2c\xa6\x5a\x03\x02\x00\x00\xff\xff\x8c\x08\x6c\xbb\x8f\x00\x00\x00")

func mysql005_oauth_scopeSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql005_oauth_scopeSql,
		"mysql/005_oauth_scope.sql",
	)
}

func mysql005_oauth_scopeSql() (*asset, error) {
	bytes, err := mysql005_oauth_scopeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/005_oauth_scope.sql", size: 143, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql006_add_orgs_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xc1\x4e\xc3\x30\x10\x44\xef\xfb\x15\x73\x4c\x04\x95\xa0\xa2\xa7\x9c\xdc\x64\x01\x8b\xd4\x2e\x5b\x07\xb5\xa7\x0a\x81\x55\x45\x40\x52\x39\x85\xfe\x3e\x72\x49\xaa\x1c\x90\xea\x93\xb5\x33\x6f\xec\xd9\xc9\x04\x57\x5f\xf5\x2e\xbc\x1e\x3c\xaa\x3d\x51\x2e\xac\x1c\xc3\xa9\x79\xc9\xd0\xf7\x30\xd6\x81\xd7\x7a\xe5\x56\x68\xc3\xae\x43\x42\x40\xbc\x6d\xeb\x77\xfc\x1d\x6d\x1c\x3f\xb0\x60\x29\x7a\xa1\x64\x83\x27\xde\x40\x55\xce\x6e\xb5\xc9\x85\x17\x6c\x1c\x01\xd7\x11\xf9\xee\x7c\x38\x71\x3d\x32\xcc\xdb\x63\xe3\x43\x8c\x7a\x51\x92\x3f\x2a\x49\xa6\xb3\x59\x3a\x88\x9f\x75\xf3\x81\xb1\x78\x7b\x33\xbd\x3b\xab\xfb\x50\xff\xc4\xaf\x63\x6e\x6d\xc9\xca\x0c\xf3\xce\xbf\x05\x7f\xf8\x27\xb2\x32\xfa\xb9\xe2\xe4\xfc\x6c\x4a\x69\x46\xa4\x4a\xc7\xd2\x97\x3e\xd5\x54\x45\x01\x6d\x0a\x5e\x63\x64\xcd\x2e\xfa\xfa\x86\x31\x72\xbc\xd8\xa2\x3d\x36\x44\x85\xd8\xe5\x88\xcd\xe8\x37\x00\x00\xff\xff\x43\x18\x08\xd4\x7c\x01\x00\x00")

func mysql006_add_orgs_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql006_add_orgs_tableSql,
		"mysql/006_add_orgs_table.sql",
	)
}

func mysql006_add_orgs_tableSql() (*asset, error) {
	bytes, err := mysql006_add_orgs_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/006_add_orgs_table.sql", size: 380, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mysql007_add_slack_urlsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x41\x4b\xc3\x40\x10\x46\xef\xf3\x2b\xbe\x63\x82\x2d\x68\xb1\xa7\x9e\xd6\x36\xa0\xa0\xad\x2c\xa9\xd7\xb0\xd6\x31\x2e\xd9\x6c\xe2\xec\x2e\x9a\x7f\x2f\xc9\x21\x04\x0f\xbd\x0d\xbc\xe1\x7d\xbc\xf5\x1a\x37\xad\xad\xc5\x44\xc6\xb9\x27\xba\x08\x8f\x67\x34\xef\x8e\x61\x3f\xe1\xbb\x08\xfe\xb5\x21\x06\x04\x67\x2e\x4d\x95\xc4\x85\x8c\x00\xfb\x01\xeb\x23\xd7\x2c\xe8\xc5\xb6\x46\x06\x34\x3c\x40\x9d\xcb\x53\xf5\x74\xdc\xeb\xe2\xa5\x38\x96\x2b\x02\xbe\xba\x10\x2b\x6f\x5a\xc6\x9b\xd2\xfb\x47\xa5\xb3\xcd\x76\x9b\x4f\x62\x9f\x9c\x1b\x5f\x52\x60\xb9\x42\xc5\xcd\xf0\xee\x76\x73\xff\x8f\x7a\xfb\x9d\x38\x9b\x57\x56\x93\x2d\xa7\x7c\x47\xb4\x8c\x3b\x74\x3f\x9e\xe8\xa0\x4f\xaf\x28\xd5\xc3\x73\xb1\xc8\xd9\xd1\x5f\x00\x00\x00\xff\xff\x5f\x31\x16\x1b\x06\x01\x00\x00")

func mysql007_add_slack_urlsSqlBytes() ([]byte, error) {
	return bindataRead(
		_mysql007_add_slack_urlsSql,
		"mysql/007_add_slack_urls.sql",
	)
}

func mysql007_add_slack_urlsSql() (*asset, error) {
	bytes, err := mysql007_add_slack_urlsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mysql/007_add_slack_urls.sql", size: 262, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres001_initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xdf\x6a\x83\x30\x14\x87\xef\xf3\x14\xe7\xb2\xb2\xf6\x09\xbc\xd2\x79\x56\xc2\x5c\xec\x62\x04\x7b\x55\xc2\x16\x24\xd4\x7f\x44\xdb\xee\xf1\x87\x21\xda\x1a\xd8\xa8\x57\xe1\x23\xe7\xf8\xfd\x4e\xce\x6e\x07\x2f\x8d\xae\x8c\x1c\x15\x14\x3d\x21\xaf\x1c\x23\x81\x20\xa2\x38\x45\xa0\x6f\xc0\x32\x01\x58\xd2\x5c\xe4\x70\x19\x94\x19\x60\x43\xec\xe1\xa4\xbf\xc1\x7e\x31\xdd\xe7\xc8\x69\x94\xc2\x81\xd3\x8f\x88\x1f\xe1\x1d\x8f\x64\x6b\xef\xd4\x5d\xa5\x5b\x00\x10\x58\x0a\x87\xc6\xee\xac\x3c\xa4\x1a\xa9\xeb\x35\x92\x57\x39\x4a\xb3\x42\x83\xfa\x32\x6a\x74\x88\x6c\x0b\x46\x3f\x0b\xdc\xdc\x7f\x13\x90\x20\xfc\x57\xdf\xa8\xbe\xb3\xfa\xd3\x61\xd1\xff\xcb\xdf\x5e\x5a\x82\x52\x26\x70\x8f\xdc\xe1\xee\xd6\x2a\x03\x8b\xb1\x65\xad\x6c\x14\x78\x6c\xa8\x2f\x95\xcf\x6a\xdd\x9e\x7d\xd6\x1b\x7d\x9d\xe6\x0f\x71\x96\xa5\x18\xb1\xb9\xdc\x25\xf6\x22\x2f\xad\x57\x89\x29\x4b\xb0\xf4\x12\xeb\x9f\xd3\xca\x37\x63\xf3\x10\xee\x38\x08\x9f\xe9\x30\x0f\xc2\xeb\xe0\xf0\xa4\xf1\xb8\x47\x49\x77\x6b\x09\x49\x78\x76\x70\x0f\x61\x6b\xc2\x47\x62\x77\x29\x24\xbf\x01\x00\x00\xff\xff\x1e\xfd\x38\xa0\x7e\x02\x00\x00")

func postgres001_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres001_initSql,
		"postgres/001_init.sql",
	)
}

func postgres001_initSql() (*asset, error) {
	bytes, err := postgres001_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/001_init.sql", size: 638, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres002_orgSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\x2d\xc8\x2f\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x03\x0b\xc4\xe7\x17\xa5\x2b\x38\xf9\xfb\xfb\xb8\x3a\xfa\x29\xb8\xb8\xba\x39\x86\xfa\x84\x28\xb8\x39\xfa\x04\xbb\x5a\x73\x71\x21\x9b\xe6\x92\x5f\x9e\x87\xcd\x3c\x97\x20\xff\x00\x74\x03\xad\x01\x01\x00\x00\xff\xff\x25\x2b\x07\x70\x87\x00\x00\x00")

func postgres002_orgSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres002_orgSql,
		"postgres/002_org.sql",
	)
}

func postgres002_orgSql() (*asset, error) {
	bytes, err := postgres002_orgSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/002_org.sql", size: 135, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres003_drop_emailsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x03\x8b\xc4\xa7\xe6\x26\x66\xe6\x58\x73\x71\x21\x6b\x75\xc9\x2f\xcf\xc3\xa6\xd9\xd1\xc5\x05\x53\xaf\x42\x98\x63\x90\xb3\x87\x63\x90\x86\x91\xa9\xa9\xa6\x82\x8b\xab\x9b\x63\xa8\x4f\x88\x82\xba\xba\x35\x20\x00\x00\xff\xff\xe3\x75\x50\xf8\x8d\x00\x00\x00")

func postgres003_drop_emailsSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres003_drop_emailsSql,
		"postgres/003_drop_emails.sql",
	)
}

func postgres003_drop_emailsSql() (*asset, error) {
	bytes, err := postgres003_drop_emailsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/003_drop_emails.sql", size: 141, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres004_limit_usersSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\xc1\x0a\x82\x40\x10\xc6\xf1\xfb\x3c\xc5\x77\x54\xca\x27\xf0\x64\x39\xc1\x42\x68\xe9\x2c\x78\x8b\x0e\xb2\x2c\xa8\x1b\xa3\xd1\xeb\x47\x2a\xe1\xa1\x43\x73\xfd\xe0\x37\xff\x24\xc1\xae\xf7\x4e\xef\x53\x0b\xfb\x20\x3a\x56\x9c\x09\x43\xb2\xc3\x99\x61\x4e\x28\x4a\x01\x37\xa6\x96\x1a\x9d\xef\xfd\x74\x7b\x8e\xad\x8e\x88\x08\xe8\x82\xf3\x03\xe6\x13\x6e\x84\x80\xbd\x2d\xcc\xd5\x72\x34\x2f\x31\xc5\xe9\x1f\x5e\x50\xb7\x70\x41\xdd\xd7\xc1\x0a\x05\x75\x0b\xb3\xcd\xcc\xc3\x6b\x20\xca\xab\xf2\xb2\xb2\x9b\xb0\xf4\xc7\xf0\xf9\x90\xd2\x3b\x00\x00\xff\xff\x2f\x48\xf5\xe6\xe9\x00\x00\x00")

func postgres004_limit_usersSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres004_limit_usersSql,
		"postgres/004_limit_users.sql",
	)
}

func postgres004_limit_usersSql() (*asset, error) {
	bytes, err := postgres004_limit_usersSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/004_limit_users.sql", size: 233, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres005_oauth_scopeSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x03\x0b\xc4\x17\x27\xe7\x17\xa4\x16\x2b\x84\x39\x06\x39\x7b\x38\x06\x69\x18\x99\x9a\x6a\x2a\xb8\xb8\xba\x39\x86\xfa\x84\x28\xa8\xab\x5b\x73\x71\x21\x1b\xe9\x92\x5f\x9e\x87\xcd\x50\x97\x20\xff\x00\x2c\xa6\x5a\x03\x02\x00\x00\xff\xff\x8c\x08\x6c\xbb\x8f\x00\x00\x00")

func postgres005_oauth_scopeSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres005_oauth_scopeSql,
		"postgres/005_oauth_scope.sql",
	)
}

func postgres005_oauth_scopeSql() (*asset, error) {
	bytes, err := postgres005_oauth_scopeSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/005_oauth_scope.sql", size: 143, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres006_add_orgs_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x50\xb1\x4e\xc3\x30\x14\xdc\xfd\x15\x37\x26\x82\x4a\x50\xd1\x29\x93\xd3\x3c\x8a\x45\xb0\x8b\xe3\xa0\x76\xaa\x10\x58\x95\x05\x24\x95\x53\x08\x9f\x8f\x0c\x89\x09\x88\x01\x4f\xd6\xdd\xbb\xf7\xee\x6e\x36\xc3\xc9\x8b\xdb\xfb\xfb\xa3\x45\x7d\x60\x6c\xa9\x89\x1b\x82\xe1\x79\x49\x10\x97\x90\xca\x80\x36\xa2\x32\x15\x5a\xbf\xef\x90\x30\x20\xfc\x76\xee\x11\x5f\x2f\x17\xab\x8a\xb4\xe0\x25\xd6\x5a\xdc\x70\xbd\xc5\x35\x6d\x19\x70\x1a\xa6\x5e\x3b\xeb\x3f\x47\x85\x34\xb4\x22\x3d\xe2\x6d\xdf\x58\x1f\xd4\x77\x5c\x2f\xaf\xb8\x4e\xe6\x8b\x45\x3a\x92\xcf\xae\x79\xc2\x94\x3c\x3f\x9b\x5f\x44\xf6\xe0\xdd\x5b\x70\x8b\x5c\xa9\x92\xb8\x1c\xf1\xce\x3e\x78\x7b\xfc\x63\x65\x2d\xc5\x6d\x4d\x49\x3c\x9b\xb2\x34\x8b\x41\x85\x2c\x68\xf3\x2b\xa8\x7b\xdf\x4d\x3d\x2a\x39\x44\xff\xde\x90\xfd\x43\x1e\xa3\xff\xd0\x0f\x68\x70\x30\xad\xbe\x68\xfb\x86\xb1\x42\xab\xf5\x50\x7d\x50\x64\xec\x23\x00\x00\xff\xff\xd2\x89\x99\x08\x9e\x01\x00\x00")

func postgres006_add_orgs_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres006_add_orgs_tableSql,
		"postgres/006_add_orgs_table.sql",
	)
}

func postgres006_add_orgs_tableSql() (*asset, error) {
	bytes, err := postgres006_add_orgs_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/006_add_orgs_table.sql", size: 414, mode: os.FileMode(420), modTime: time.Unix(1585151454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgres007_add_slack_urlsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xcd\x6a\x83\x40\x14\x46\xf7\xf7\x29\x3e\x5c\x29\xd5\x27\x70\xa5\x28\x45\x6a\xa9\x4c\xdb\x85\x2b\x99\xda\x9b\x64\xc8\x38\x26\xf3\x43\x7c\xfc\x60\x08\x41\x48\x76\x97\x73\xf9\x0e\x27\xcb\xf0\x36\xa9\xbd\x95\x9e\xf1\x7b\x22\x1a\x2d\xaf\xa7\x97\x7f\x9a\xa1\x76\x30\xb3\x07\x2f\xca\x79\x07\xa7\xe5\x78\x1c\x82\xd5\x2e\x26\x40\xfd\xa3\x6c\xde\xbf\x6b\xd1\x14\x2d\x3a\xd1\x7c\x16\xa2\xc7\x47\xdd\xa7\x04\x1c\x66\xe7\x07\x23\x27\x86\xe7\xc5\xdf\x1c\x26\x68\xbd\xbe\xa2\xe0\xd8\x46\xcf\x3c\x58\xfd\x02\x1a\x75\x0e\x1c\x3f\x74\xe9\x7d\x9e\x50\x92\x13\x6d\xd3\xab\xf9\x62\x88\x2a\xf1\xd5\xe1\xa7\x28\xdb\x7a\x13\x9b\xd3\x35\x00\x00\xff\xff\x28\x08\x6b\x88\xe4\x00\x00\x00")

func postgres007_add_slack_urlsSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres007_add_slack_urlsSql,
		"postgres/007_add_slack_urls.sql",
	)
}

func postgres007_add_slack_urlsSql() (*asset, error) {
	bytes, err := postgres007_add_slack_urlsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/007_add_slack_urls.sql", size: 228, mode: os.FileMode(420), modTime: time.Unix(1586188331, 0)}
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
	"sqlite3/001_init.sql":            sqlite3001_initSql,
	"sqlite3/002_org.sql":             sqlite3002_orgSql,
	"sqlite3/003_drop_emails.sql":     sqlite3003_drop_emailsSql,
	"sqlite3/004_limit_users.sql":     sqlite3004_limit_usersSql,
	"sqlite3/005_oauth_scope.sql":     sqlite3005_oauth_scopeSql,
	"sqlite3/006_add_orgs_table.sql":  sqlite3006_add_orgs_tableSql,
	"sqlite3/007_add_slack_urls.sql":  sqlite3007_add_slack_urlsSql,
	"mysql/001_init.sql":              mysql001_initSql,
	"mysql/002_org.sql":               mysql002_orgSql,
	"mysql/003_drop_emails.sql":       mysql003_drop_emailsSql,
	"mysql/004_limit_users.sql":       mysql004_limit_usersSql,
	"mysql/005_oauth_scope.sql":       mysql005_oauth_scopeSql,
	"mysql/006_add_orgs_table.sql":    mysql006_add_orgs_tableSql,
	"mysql/007_add_slack_urls.sql":    mysql007_add_slack_urlsSql,
	"postgres/001_init.sql":           postgres001_initSql,
	"postgres/002_org.sql":            postgres002_orgSql,
	"postgres/003_drop_emails.sql":    postgres003_drop_emailsSql,
	"postgres/004_limit_users.sql":    postgres004_limit_usersSql,
	"postgres/005_oauth_scope.sql":    postgres005_oauth_scopeSql,
	"postgres/006_add_orgs_table.sql": postgres006_add_orgs_tableSql,
	"postgres/007_add_slack_urls.sql": postgres007_add_slack_urlsSql,
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
		"001_init.sql":           &bintree{mysql001_initSql, map[string]*bintree{}},
		"002_org.sql":            &bintree{mysql002_orgSql, map[string]*bintree{}},
		"003_drop_emails.sql":    &bintree{mysql003_drop_emailsSql, map[string]*bintree{}},
		"004_limit_users.sql":    &bintree{mysql004_limit_usersSql, map[string]*bintree{}},
		"005_oauth_scope.sql":    &bintree{mysql005_oauth_scopeSql, map[string]*bintree{}},
		"006_add_orgs_table.sql": &bintree{mysql006_add_orgs_tableSql, map[string]*bintree{}},
		"007_add_slack_urls.sql": &bintree{mysql007_add_slack_urlsSql, map[string]*bintree{}},
	}},
	"postgres": &bintree{nil, map[string]*bintree{
		"001_init.sql":           &bintree{postgres001_initSql, map[string]*bintree{}},
		"002_org.sql":            &bintree{postgres002_orgSql, map[string]*bintree{}},
		"003_drop_emails.sql":    &bintree{postgres003_drop_emailsSql, map[string]*bintree{}},
		"004_limit_users.sql":    &bintree{postgres004_limit_usersSql, map[string]*bintree{}},
		"005_oauth_scope.sql":    &bintree{postgres005_oauth_scopeSql, map[string]*bintree{}},
		"006_add_orgs_table.sql": &bintree{postgres006_add_orgs_tableSql, map[string]*bintree{}},
		"007_add_slack_urls.sql": &bintree{postgres007_add_slack_urlsSql, map[string]*bintree{}},
	}},
	"sqlite3": &bintree{nil, map[string]*bintree{
		"001_init.sql":           &bintree{sqlite3001_initSql, map[string]*bintree{}},
		"002_org.sql":            &bintree{sqlite3002_orgSql, map[string]*bintree{}},
		"003_drop_emails.sql":    &bintree{sqlite3003_drop_emailsSql, map[string]*bintree{}},
		"004_limit_users.sql":    &bintree{sqlite3004_limit_usersSql, map[string]*bintree{}},
		"005_oauth_scope.sql":    &bintree{sqlite3005_oauth_scopeSql, map[string]*bintree{}},
		"006_add_orgs_table.sql": &bintree{sqlite3006_add_orgs_tableSql, map[string]*bintree{}},
		"007_add_slack_urls.sql": &bintree{sqlite3007_add_slack_urlsSql, map[string]*bintree{}},
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
