package util

import (
	"encoding/json"
	"errors"
)

// StrToMap takes a JSON formatted string and returns it as a Go struct.
// It returns an error if the operation fails.
func StrToMap(input string) (map[string]interface{}, error) {
	var buffer interface{}
	err := json.Unmarshal([]byte(input), &buffer)
	if err != nil {
		return nil, err
	}
	casted, ok := buffer.(map[string]interface{})
	if !ok {
		return nil, errors.New("ERROR: Query is not a JSON object")
	}

	return casted, nil
}
