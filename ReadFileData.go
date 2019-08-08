package simpleutils

import (
	"io/ioutil"
	"os"
)

//ReadFileData f
func ReadFileData(filePath string) ([]byte, error) {

	var data []byte
	var err error
	var f *os.File

	if f, err = os.OpenFile(filePath, os.O_RDWR|os.O_SYNC, 0755); err != nil {
		return nil, err
	}
	defer f.Close()

	if data, err = ioutil.ReadAll(f); err != nil {
		return nil, err
	}

	return data, nil
}
