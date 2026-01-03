// Package jsonutil provides JSON manipulation utilities
package jsonutil

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ConvertValuesToString converts all values in a JSON object to strings
func ConvertValuesToString(data interface{}) (interface{}, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, val := range v {
			converted, err := ConvertValuesToString(val)
			if err != nil {
				return nil, err
			}
			result[key] = converted
		}
		return result, nil
	case []interface{}:
		result := make([]interface{}, len(v))
		for i, val := range v {
			converted, err := ConvertValuesToString(val)
			if err != nil {
				return nil, err
			}
			result[i] = converted
		}
		return result, nil
	case string:
		return v, nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(v), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case bool:
		return strconv.FormatBool(v), nil
	case nil:
		return "", nil
	default:
		// For other types, try to marshal to JSON and then convert
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal value: %w", err)
		}
		return string(jsonBytes), nil
	}
}

// ConvertJSONStringValuesToString converts all values in a JSON string to strings
func ConvertJSONStringValuesToString(jsonStr string) (string, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	converted, err := ConvertValuesToString(data)
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(converted)
	if err != nil {
		return "", fmt.Errorf("failed to marshal converted JSON: %w", err)
	}

	return string(jsonBytes), nil
}

// GetValueByPath gets a value from JSON data using a path (e.g., "user.name" or "items[0].name")
func GetValueByPath(data interface{}, path string) (interface{}, error) {
	if path == "" {
		return data, nil
	}

	parts := parsePath(path)
	current := data

	for i, part := range parts {
		if current == nil {
			return nil, fmt.Errorf("path segment '%s' at index %d: value is nil", part, i)
		}

		switch v := current.(type) {
		case map[string]interface{}:
			key, _, isArray := parsePart(part)
			if isArray {
				return nil, fmt.Errorf("path segment '%s' at index %d: cannot use array index on map", part, i)
			}
			var exists bool
			current, exists = v[key]
			if !exists {
				return nil, fmt.Errorf("path segment '%s' at index %d: key '%s' not found", part, i, key)
			}
		case []interface{}:
			key, index, isArray := parsePart(part)
			if isArray {
				if index < 0 || index >= len(v) {
					return nil, fmt.Errorf("path segment '%s' at index %d: array index %d out of range", part, i, index)
				}
				current = v[index]
			} else {
				// Try to find in array elements if they are maps
				var found bool
				for _, item := range v {
					if itemMap, ok := item.(map[string]interface{}); ok {
						if val, exists := itemMap[key]; exists {
							current = val
							found = true
							break
						}
					}
				}
				if !found {
					return nil, fmt.Errorf("path segment '%s' at index %d: key '%s' not found in array elements", part, i, key)
				}
			}
		default:
			return nil, fmt.Errorf("path segment '%s' at index %d: cannot traverse type %T", part, i, current)
		}
	}

	return current, nil
}

// GetStringByPath gets a string value from JSON data using a path
func GetStringByPath(data interface{}, path string) (string, error) {
	value, err := GetValueByPath(data, path)
	if err != nil {
		return "", err
	}
	return convertToString(value), nil
}

// GetIntByPath gets an int value from JSON data using a path
func GetIntByPath(data interface{}, path string) (int, error) {
	value, err := GetValueByPath(data, path)
	if err != nil {
		return 0, err
	}
	return convertToInt(value), nil
}

// GetFloat64ByPath gets a float64 value from JSON data using a path
func GetFloat64ByPath(data interface{}, path string) (float64, error) {
	value, err := GetValueByPath(data, path)
	if err != nil {
		return 0, err
	}
	return convertToFloat64(value), nil
}

// GetBoolByPath gets a bool value from JSON data using a path
func GetBoolByPath(data interface{}, path string) (bool, error) {
	value, err := GetValueByPath(data, path)
	if err != nil {
		return false, err
	}
	return convertToBool(value), nil
}

// SetValueByPath sets a value in JSON data using a path
func SetValueByPath(data interface{}, path string, value interface{}) error {
	if path == "" {
		return fmt.Errorf("path cannot be empty")
	}

	parts := parsePath(path)
	if len(parts) == 0 {
		return fmt.Errorf("invalid path")
	}

	// Navigate to the parent and set the value
	if len(parts) == 1 {
		return setValueAtPath(data, parts[0], value)
	}

	parentPath := strings.Join(parts[:len(parts)-1], ".")
	lastPart := parts[len(parts)-1]

	parent, err := GetValueByPath(data, parentPath)
	if err != nil {
		return fmt.Errorf("failed to get parent path '%s': %w", parentPath, err)
	}

	return setValueAtPath(parent, lastPart, value)
}

// FindPaths finds all paths in JSON data that match a pattern or contain a specific value
func FindPaths(data interface{}, options *FindOptions) ([]string, error) {
	if options == nil {
		options = &FindOptions{}
	}

	var paths []string
	err := findPathsRecursive(data, "", options, &paths)
	return paths, err
}

// FindOptions provides options for finding paths
type FindOptions struct {
	// KeyPattern is a regex pattern to match keys
	KeyPattern string
	// ValuePattern is a regex pattern to match values (converted to string)
	ValuePattern string
	// ExactValue matches exact value (after conversion to string)
	ExactValue string
	// ValueType filters by value type (e.g., "string", "number", "bool", "object", "array")
	ValueType string
}

// HasPath checks if a path exists in JSON data
func HasPath(data interface{}, path string) bool {
	_, err := GetValueByPath(data, path)
	return err == nil
}

// GetAllPaths returns all possible paths in JSON data
func GetAllPaths(data interface{}) []string {
	var paths []string
	getAllPathsRecursive(data, "", &paths)
	return paths
}

// parsePath parses a path string into parts
func parsePath(path string) []string {
	var parts []string
	var current strings.Builder
	inBrackets := false

	for _, char := range path {
		switch char {
		case '.':
			if !inBrackets {
				if current.Len() > 0 {
					parts = append(parts, current.String())
					current.Reset()
				}
			} else {
				current.WriteRune(char)
			}
		case '[':
			if current.Len() > 0 {
				parts = append(parts, current.String())
				current.Reset()
			}
			inBrackets = true
		case ']':
			if inBrackets {
				parts = append(parts, "["+current.String()+"]")
				current.Reset()
				inBrackets = false
			}
		default:
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}

// parsePart parses a path part (e.g., "key", "[0]", "key[0]")
func parsePart(part string) (key string, index int, isArray bool) {
	if strings.HasPrefix(part, "[") && strings.HasSuffix(part, "]") {
		// Array index only: [0]
		indexStr := part[1 : len(part)-1]
		idx, err := strconv.Atoi(indexStr)
		if err != nil {
			return "", -1, false
		}
		return "", idx, true
	}

	if idx := strings.Index(part, "["); idx != -1 {
		// Key with array index: key[0]
		key = part[:idx]
		indexStr := part[idx+1 : strings.Index(part, "]")]
		idx, err := strconv.Atoi(indexStr)
		if err != nil {
			return key, -1, false
		}
		return key, idx, true
	}

	// Just a key
	return part, -1, false
}

// setValueAtPath sets a value at a specific path part
func setValueAtPath(data interface{}, part string, value interface{}) error {
	key, index, isArray := parsePart(part)

	switch v := data.(type) {
	case map[string]interface{}:
		if isArray {
			return fmt.Errorf("cannot use array index on map")
		}
		v[key] = value
		return nil
	case []interface{}:
		if isArray {
			if index < 0 || index >= len(v) {
				return fmt.Errorf("array index %d out of range", index)
			}
			v[index] = value
			return nil
		}
		return fmt.Errorf("cannot set key '%s' on array without index", key)
	default:
		return fmt.Errorf("cannot set value on type %T", data)
	}
}

// findPathsRecursive recursively finds paths matching the options
func findPathsRecursive(data interface{}, currentPath string, options *FindOptions, paths *[]string) error {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, val := range v {
			newPath := currentPath
			if newPath != "" {
				newPath += "."
			}
			newPath += key

			if matchesOptions(val, key, options) {
				*paths = append(*paths, newPath)
			}

			if err := findPathsRecursive(val, newPath, options, paths); err != nil {
				return err
			}
		}
	case []interface{}:
		for i, val := range v {
			newPath := fmt.Sprintf("%s[%d]", currentPath, i)

			if matchesOptions(val, "", options) {
				*paths = append(*paths, newPath)
			}

			if err := findPathsRecursive(val, newPath, options, paths); err != nil {
				return err
			}
		}
	}
	return nil
}

// matchesOptions checks if a value matches the find options
func matchesOptions(value interface{}, key string, options *FindOptions) bool {
	if options == nil {
		return false
	}

	// Check key pattern
	if options.KeyPattern != "" && key != "" {
		matched, _ := regexp.MatchString(options.KeyPattern, key)
		if !matched {
			return false
		}
	}

	// Check value type
	if options.ValueType != "" {
		valueType := getValueType(value)
		if valueType != options.ValueType {
			return false
		}
	}

	// Check exact value
	if options.ExactValue != "" {
		valueStr := convertToString(value)
		if valueStr != options.ExactValue {
			return false
		}
	}

	// Check value pattern
	if options.ValuePattern != "" {
		valueStr := convertToString(value)
		matched, _ := regexp.MatchString(options.ValuePattern, valueStr)
		if !matched {
			return false
		}
	}

	return true
}

// getAllPathsRecursive recursively gets all paths
func getAllPathsRecursive(data interface{}, currentPath string, paths *[]string) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, val := range v {
			newPath := currentPath
			if newPath != "" {
				newPath += "."
			}
			newPath += key
			*paths = append(*paths, newPath)
			getAllPathsRecursive(val, newPath, paths)
		}
	case []interface{}:
		for i, val := range v {
			newPath := fmt.Sprintf("%s[%d]", currentPath, i)
			*paths = append(*paths, newPath)
			getAllPathsRecursive(val, newPath, paths)
		}
	}
}

// getValueType returns the type of a value as a string
func getValueType(value interface{}) string {
	if value == nil {
		return "null"
	}

	switch value.(type) {
	case string:
		return "string"
	case float64, float32, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return "number"
	case bool:
		return "bool"
	case map[string]interface{}:
		return "object"
	case []interface{}:
		return "array"
	default:
		return "unknown"
	}
}

// Helper conversion functions
func convertToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case bool:
		return strconv.FormatBool(v)
	default:
		jsonBytes, _ := json.Marshal(v)
		return string(jsonBytes)
	}
}

func convertToInt(value interface{}) int {
	if value == nil {
		return 0
	}

	switch v := value.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case float32:
		return int(v)
	case int64:
		return int(v)
	case string:
		i, _ := strconv.Atoi(v)
		return i
	default:
		return 0
	}
}

func convertToFloat64(value interface{}) float64 {
	if value == nil {
		return 0
	}

	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		f, _ := strconv.ParseFloat(v, 64)
		return f
	default:
		return 0
	}
}

func convertToBool(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case bool:
		return v
	case string:
		b, _ := strconv.ParseBool(v)
		return b
	case int:
		return v != 0
	case float64:
		return v != 0
	default:
		return false
	}
}
