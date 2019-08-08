package simpleutils

import (
	"encoding/json"
)

//ReadJSON f
func ReadJSON(filePath string, obj interface{}) (interface{}, error) {

	var data []byte
	var err error

	if data, err = ReadFileData(filePath); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	return obj, nil
}
