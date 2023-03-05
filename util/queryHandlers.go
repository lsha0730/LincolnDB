package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var DBROOT string = "./data"

func NewDB(name string) (*Driver, error) {
	root := DBROOT + "/" + name

	driver := Driver{
		name: name,
		root: root,
	}

	// if _, err := os.Stat(root); err == nil {
	// 	return nil, errors.New("Database w id '" + name + "' already exists\n")
	// }

	return &driver, nil
}

func (d *Driver) HandleRead(path string) error {
	return nil
}

func (d *Driver) HandleWrite(path string, value interface{}) error {
	os.Mkdir(DBROOT, 0644)

	finalPath := d.root
	tempPath := finalPath + ".tmp"

	b, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		return err
	}
	b = append(b, byte('\n'))

	if err := ioutil.WriteFile(tempPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tempPath, finalPath)
}

func (d *Driver) HandleList(path string) error {
	return nil
}
