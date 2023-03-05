package util

import (
	"errors"
	"reflect"
)

func ValidateQuery(castedQuery map[string]interface{}) error {
	// Mandatory fields checking
	if err := ValidateOp(castedQuery); err != nil {
		return err
	}
	if err := ValidatePath(castedQuery); err != nil {
		return err
	}

	queryType := castedQuery["op"]
	switch queryType {
	case READ:
		if err := ValidateRead(castedQuery); err != nil {
			return err
		}
	case WRITE:
		if err := ValidateWrite(castedQuery); err != nil {
			return err
		}
	case LIST:
		if err := ValidateList(castedQuery); err != nil {
			return err
		}
		// case MAKECOLLECTION:
		// 	if err := ValidateMake(castedQuery); err != nil {
		// 		return err
		// 	}
		// case MAKEDOCUMENT:
		// 	if err := ValidateMake(castedQuery); err != nil {
		// 		return err
		// 	}
	}

	return nil
}

func ValidateRead(query map[string]interface{}) error {
	op := query["op"]
	if op != READ {
		return errors.New("ERROR: Not a read operation")
	}

	// TODO: Check if path being read from exists
	return nil
}

func ValidateWrite(query map[string]interface{}) error {
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

func ValidateList(query map[string]interface{}) error {
	op := query["op"]
	if op != LIST {
		return errors.New("ERROR: Not a list operation")
	}

	return nil
}

func ValidateMake(query map[string]interface{}) error {
	// op := query["op"]
	// if op != MAKECOLLECTION && op != MAKEDOCUMENT {
	// 	return errors.New("ERROR: Not a make operation")
	// }

	return nil
}

// Helper functions
func ValidateOp(query map[string]interface{}) error {
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
	// validOps[MAKECOLLECTION] = true
	// validOps[MAKEDOCUMENT] = true
	validOps[MAKEDB] = true

	if !validOps[op.(string)] {
		return errors.New("ERROR: Invalid operand clause")
	}

	return nil
}

func ValidatePath(query map[string]interface{}) error {
	path, pathExists := query["path"]
	if !pathExists {
		return errors.New("ERROR: Query is missing path property")
	}
	if reflect.TypeOf(path).String() != "string" {
		return errors.New("ERROR: Path is not a string")
	}

	if !IsNormal(path.(string)) {
		return errors.New("ERROR: Invalid path string")
	}
	return nil
}

func IsNormal(input string) bool {
	for _, r := range input {
		if r < 33 || r > 127 {
			return false
		}
	}
	return true
}
