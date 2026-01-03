// Package convert provides type conversion utilities
package convert

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// ToString converts an interface{} value to string
func ToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case time.Time:
		t, _ := value.(time.Time)
		key = t.String()
		// Remove timezone suffix
		key = strings.Replace(key, " +0800 CST", "", 1)
		key = strings.Replace(key, " +0000 UTC", "", 1)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

// ToInt converts an interface{} value to int
func ToInt(v interface{}) int {
	var r int
	switch v.(type) {
	case uint:
		r = int(v.(uint))
	case int8:
		r = int(v.(int8))
	case uint8:
		r = int(v.(uint8))
	case int16:
		r = int(v.(int16))
	case uint16:
		r = int(v.(uint16))
	case int32:
		r = int(v.(int32))
	case uint32:
		r = int(v.(uint32))
	case int64:
		r = int(v.(int64))
	case uint64:
		r = int(v.(uint64))
	case float32:
		r = int(v.(float32))
	case float64:
		r = int(v.(float64))
	case string:
		r, _ = strconv.Atoi(v.(string))
		if r == 0 && len(v.(string)) > 0 {
			f, _ := strconv.ParseFloat(v.(string), 64)
			r = int(f)
		}
	case nil:
		r = 0
	case json.Number:
		t3, _ := v.(json.Number).Int64()
		r = int(t3)
	default:
		r = v.(int)
	}
	return r
}

// ToInt64 converts an interface{} value to int64
func ToInt64(v interface{}) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		i, _ := strconv.ParseInt(val, 10, 64)
		return i
	case json.Number:
		i, _ := val.Int64()
		return i
	default:
		return 0
	}
}

// ToFloat64 converts an interface{} value to float64
func ToFloat64(v interface{}) float64 {
	switch val := v.(type) {
	case float32:
		return float64(val)
	case float64:
		return val
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	case json.Number:
		f, _ := val.Float64()
		return f
	default:
		return 0
	}
}

// ToBool converts an interface{} value to bool
func ToBool(v interface{}) bool {
	switch val := v.(type) {
	case bool:
		return val
	case string:
		b, _ := strconv.ParseBool(val)
		return b
	case int:
		return val != 0
	case int8:
		return val != 0
	case int16:
		return val != 0
	case int32:
		return val != 0
	case int64:
		return val != 0
	default:
		return false
	}
}
