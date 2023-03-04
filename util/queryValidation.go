package util

import (
	"encoding/json"
	"errors"
	"reflect"
)

func ValidateQuery(queryString string) error {
	var query interface{}
	err := json.Unmarshal([]byte(queryString), &query)
	if err != nil {
		return err
	}
	castedQuery, ok := query.(map[string]interface{})
	if !ok {
		return errors.New("ERROR: Query is not a JSON object")
	}

	// Mandatory fields checking
	if err := validateOp(castedQuery); err != nil {
		return err
	}
	if err := validatePath(castedQuery); err != nil {
		return err
	}

	queryType := castedQuery["op"]
	switch queryType {
	case READ:
		if err := validateRead(castedQuery); err != nil {
			return err
		}
	case WRITE:
		if err := validateWrite(castedQuery); err != nil {
			return err
		}
	case LIST:
		if err := validateList(castedQuery); err != nil {
			return err
		}
	case MAKECOLLECTION:
		if err := validateMake(castedQuery); err != nil {
			return err
		}
	case MAKEDOCUMENT:
		if err := validateMake(castedQuery); err != nil {
			return err
		}
	default:
		return errors.New("ERROR: Unexpected operation type")
	}

	return nil
}

func validateRead(query map[string]interface{}) error {
	op := query["op"]
	if op != READ {
		return errors.New("ERROR: Not a read operation")
	}

	// TODO: Check if path being read from exists
	return nil
}

func validateWrite(query map[string]interface{}) error {
	op := query["op"]
	_, valueExists := query["value"]
	if op != WRITE {
		return errors.New("ERROR: Not a write operation")
	}

	if !valueExists {
		return errors.New("ERROR: No value found")
	}

	return nil
}

func validateList(query map[string]interface{}) error {
	op := query["op"]
	if op != LIST {
		return errors.New("ERROR: Not a list operation")
	}

	return nil
}

func validateMake(query map[string]interface{}) error {
	op := query["op"]
	if op != MAKECOLLECTION && op != MAKEDOCUMENT {
		return errors.New("ERROR: Not a make operation")
	}

	return nil
}

// Helper functions
func validateOp(query map[string]interface{}) error {
	op, opExists := query["op"]
	if !opExists {
		return errors.New("ERROR: Query is missing op property")
	}
	if reflect.TypeOf(op).String() != "string" {
		return errors.New("ERROR: Operation is not a string")
	}

	validOps := make(map[string]bool)
	validOps[READ] = true
	validOps[WRITE] = true
	validOps[LIST] = true
	validOps[MAKECOLLECTION] = true
	validOps[MAKEDOCUMENT] = true

	if !validOps[op.(string)] {
		return errors.New("ERROR: Invalid operand clause")
	}

	return nil
}

func validatePath(query map[string]interface{}) error {
	path, pathExists := query["path"]
	if !pathExists {
		return errors.New("ERROR: Query is missing path property")
	}
	if reflect.TypeOf(path).String() != "string" {
		return errors.New("ERROR: Path is not a string")
	}

	if !isNormal(path.(string)) {
		return errors.New("ERROR: Invalid path string")
	}
	return nil
}

func isNormal(input string) bool {
	for _, r := range input {
		if r < 33 || r > 127 {
			return false
		}
	}
	return true
}
