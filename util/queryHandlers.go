package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
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
	// Check not writing to root
	pathArr := strings.Split(path, "/")
	if pathArr[0] == "" {
		pathArr = pathArr[1:]
	}
	if pathArr[len(pathArr)-1] == "" {
		pathArr = pathArr[:len(pathArr)-1]
	}
	if len(pathArr) == 0 {
		return errors.New("ERROR: Cannot write to root")
	}

	// Ensure data directory exists
	os.Mkdir(DBROOT, 0644)
	finalPath := d.root + ".json"
	tempPath := finalPath + ".tmp"

	// Get the contents
	original := map[string]interface{}{}
	b1, err1 := ioutil.ReadFile(finalPath)
	if err1 == nil {
		json.Unmarshal(b1, &original)
	}
	toWrite := InjectData(original, path, value)

	// Write the contents
	b2, err2 := json.MarshalIndent(toWrite, "", "\t")
	if err2 != nil {
		return err2
	}
	b2 = append(b2, byte('\n'))

	if err := ioutil.WriteFile(tempPath, b2, 0644); err != nil {
		return err
	}

	return os.Rename(tempPath, finalPath)
}

func (d *Driver) HandleList(path string) error {
	return nil
}
