package parser

import (
	"strconv"
	"strings"
)

func coerce(value string) interface{} {
	// `coerceNull' needs to go last, since it returns a pointer (since
	// only pointers can be `nil'
	return coerceNull(coerceBool(coerceFloat(coerceInt(value))))
}

func createNestedHash(path string, value string) map[string]interface{} {
	parts := strings.Split(path, ".")
	document := make(map[string]interface{})
	currentDocument := &document
	levels := len(parts)

	for index, part := range parts {
		if _, ok := (*currentDocument)[part]; !ok {
			if index != (levels - 1) {
				newDocument := make(map[string]interface{})
				(*currentDocument)[part] = newDocument
				currentDocument = &newDocument
			} else {
				(*currentDocument)[part] = coerce(value)
			}
		}
	}
	return document
}

func coerceNull(value interface{}) interface{} {
	if !isString(value) {
		return value
	}

	strValue := value.(string)

	var result *string

	if value == "null" {
		result = nil
	} else {
		result = &strValue
	}
	return result
}

func isString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func coerceBool(value interface{}) interface{} {
	if !isString(value) {
		return value
	}

	var result interface{}

	if value == "true" {
		result = true
	} else if value == "false" {
		result = false
	} else {
		result = value
	}

	return result
}

func coerceInt(value interface{}) interface{} {
	if !isString(value) {
		return value
	}

	result, err := strconv.Atoi(value.(string))

	if err != nil {
		return value
	}

	return result
}

func coerceFloat(value interface{}) interface{} {
	if !isString(value) {
		return value
	}

	// 53-bit floating point to mimic Python
	result, err := strconv.ParseFloat(value.(string), 53)

	if err != nil {
		return value
	}

	return result
}

func Parse(value string) map[string]interface{} {
	parts := strings.SplitN(value, "=", 2)
	return createNestedHash(parts[0], parts[1])
}
