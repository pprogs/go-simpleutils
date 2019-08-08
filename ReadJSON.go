package simpleutils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//ReadJSON f
func ReadJSON(filePath string, obj interface{}) (interface{}, error) {

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

	if err = json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}
