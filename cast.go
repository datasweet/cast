package cast

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var (
	// DateFormats are our valid date formats
	DateFormats = []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		time.RFC3339,
		"02/01/2006",
		"02/01/2006 15:04:05",
		"20060102",
		"20060102150405",
	}
)

// AsBool to convert as bool
func AsBool(v interface{}) (bool, bool) {
	switch d := v.(type) {
	case bool:
		return d, true
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(d).Int() > 0, true
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(d).Uint() > 0, true
	case float32, float64:
		return reflect.ValueOf(d).Float() > 0, true
	case string:
		if b, err := strconv.ParseBool(d); err == nil {
			return b, true
		}
		return false, false

	default:
		return false, false
	}

}

// AsBoolArray to convert as an array of bool
func AsBoolArray(values ...interface{}) ([]bool, bool) {
	arr := make([]bool, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsBool(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}

// AsString to convert as string
func AsString(v interface{}) (string, bool) {
	switch d := v.(type) {
	case string:
		return d, true
	case int:
		return strconv.FormatInt(int64(d), 10), true
	case int8:
		return strconv.FormatInt(int64(d), 10), true
	case int16:
		return strconv.FormatInt(int64(d), 10), true
	case int32:
		return strconv.FormatInt(int64(d), 10), true
	case int64:
		return strconv.FormatInt(d, 10), true
	case uint:
		return strconv.FormatUint(uint64(d), 10), true
	case uint8:
		return strconv.FormatUint(uint64(d), 10), true
	case uint16:
		return strconv.FormatUint(uint64(d), 10), true
	case uint32:
		return strconv.FormatUint(uint64(d), 10), true
	case uint64:
		return strconv.FormatUint(uint64(d), 10), true
	case float32:
		return strconv.FormatFloat(float64(d), 'f', -1, 32), true
	case float64:
		return strconv.FormatFloat(d, 'f', -1, 64), true
	case json.Number:
		return d.String(), true
	case bool:
		return strconv.FormatBool(d), true
	default:
		if v == nil {
			return "", false
		}

		return fmt.Sprintf("%v", v), false
	}
}

// AsStringArray to convert as an array of string
func AsStringArray(values ...interface{}) ([]string, bool) {
	arr := make([]string, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsString(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}

// AsInt to convert as a int
func AsInt(v interface{}) (int64, bool) {
	switch d := v.(type) {
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(d).Int(), true
	case float32, float64:
		return int64(reflect.ValueOf(d).Float()), true
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(d).Uint()), true
	case json.Number:
		if f, err := d.Int64(); err == nil {
			return f, true
		}
		return 0, false
	case string:
		if i, err := strconv.ParseInt(d, 10, 64); err == nil {
			return i, true
		}
		return 0, false
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	default:
		return 0, false
	}
}

// AsIntArray to convert as an array of int64
func AsIntArray(values ...interface{}) ([]int64, bool) {
	arr := make([]int64, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsInt(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}

// AsFloat to convert as a float64
func AsFloat(v interface{}) (float64, bool) {
	switch d := v.(type) {
	case float32, float64:
		return reflect.ValueOf(d).Float(), true
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(d).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(d).Uint()), true
	case json.Number:
		if f, err := d.Float64(); err == nil {
			return f, true
		}
		return 0, false
	case string:
		if f, err := strconv.ParseFloat(d, 64); err == nil {
			return f, true
		}
		return 0, false
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	default:
		return 0, false
	}
}

// AsFloatArray to convert as an array of float64
func AsFloatArray(values ...interface{}) ([]float64, bool) {
	arr := make([]float64, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsFloat(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}

// AsDatetime to convert as datetime (time.Time)
func AsDatetime(v interface{}) (time.Time, bool) {
	switch d := v.(type) {
	case time.Time:
		return d.UTC(), true
	case int: // timestamp compliant with javascript
		return time.Unix(0, int64(d)*int64(time.Millisecond)).UTC(), true
	case int64:
		return time.Unix(0, d*int64(time.Millisecond)).UTC(), true
	case string:
		// Try formats
		for _, format := range DateFormats {
			if dt, err := time.Parse(format, d); err == nil {
				return dt.UTC(), true
			}
		}

		// Try to convert to int64
		if ts, err := strconv.ParseInt(d, 10, 64); err == nil {
			return time.Unix(0, ts*int64(time.Millisecond)).UTC(), true
		}

		return time.Time{}, false
	default:
		return time.Time{}, false
	}
}

// AsDatetimeArray to convert as an array of datetime
func AsDatetimeArray(values ...interface{}) ([]time.Time, bool) {
	arr := make([]time.Time, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsDatetime(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}
