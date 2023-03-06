package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Returns data at given node or nil if it doesn't exist
func GetData(target map[string]interface{}, path string) interface{} {
	pathArr := strings.Split(path, "/")
	if pathArr[0] == "" {
		pathArr = pathArr[1:]
	}
	if pathArr[len(pathArr)-1] == "" {
		pathArr = pathArr[:len(pathArr)-1]
	}

	var traverse func(node interface{}, pathArr []string) interface{}
	traverse = func(node interface{}, pathArr []string) interface{} {
		if len(pathArr) == 0 {
			return node
		}
		if _, ok := node.(map[string]interface{}); !ok {
			return nil
		}

		return traverse(node.(map[string]interface{})[pathArr[0]], pathArr[1:])
	}

	return traverse(target, pathArr)
}

// Takes a struct and injects the given data at the specified location.
func InjectData(target map[string]interface{}, path string, value interface{}) map[string]interface{} {
	pathArr := strings.Split(path, "/")
	if pathArr[0] == "" {
		pathArr = pathArr[1:]
	}
	if pathArr[len(pathArr)-1] == "" {
		pathArr = pathArr[:len(pathArr)-1]
	}

	// Takes a node and overwrites value at specified path with given value
	var traverse func(node interface{}, pathArr []string, value interface{}) interface{}
	traverse = func(node interface{}, pathArr []string, value interface{}) interface{} {
		if len(pathArr) == 0 {
			return value
		}
		target := pathArr[0]

		if _, ok := node.(map[string]interface{}); !ok {
			// Primitive
			if len(pathArr) != 0 {
				return makeObj(pathArr, value)
			} else {
				return value
			}
		}

		// Object
		result := map[string]interface{}{}
		castedNode := node.(map[string]interface{})
		targetFound := false
		for key := range castedNode {
			if key == target {
				result[key] = traverse(castedNode[key], pathArr[1:], value)
				targetFound = true
			} else {
				result[key] = castedNode[key]
			}
		}

		if !targetFound {
			result[target] = traverse(castedNode[target], pathArr[1:], value)
		}

		return result
	}

	return traverse(target, pathArr, value).(map[string]interface{})
}

// Returns a nested object structure according to the path array, with the given value attached at the end
func makeObj(pathArr []string, value interface{}) interface{} {
	if len(pathArr) == 0 {
		return value
	}

	result := map[string]interface{}{}
	if len(pathArr) == 1 {
		result[pathArr[0]] = value
	} else {
		result[pathArr[0]] = makeObj(pathArr[1:], value)
	}

	return result
}

func PrintJSON(input interface{}) {
	jsonString, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Print the JSON string to the console
	fmt.Println(string(jsonString))
}
