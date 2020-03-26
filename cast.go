package cast

import (
	"reflect"
	"strconv"
	"time"
)

var (
	// dateFormats are our default date formats
	dateFormats = []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
		time.RFC3339,
		"02/01/2006",
		"02/01/2006 15:04:05",
		"20060102",
		"20060102150405",
	}
)

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// AsBool to convert as bool
func AsBool(v interface{}) (bool, bool) {
	switch d := indirect(v).(type) {
	case bool:
		return d, true
	case int:
		return d > 0, true
	case int64:
		return d > 0, true
	case int32:
		return d > 0, true
	case int16:
		return d > 0, true
	case int8:
		return d > 0, true
	case uint:
		return d > 0, true
	case uint64:
		return d > 0, true
	case uint32:
		return d > 0, true
	case uint16:
		return d > 0, true
	case uint8:
		return d > 0, true
	case float64:
		return d > 0, true
	case float32:
		return d > 0, true
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
	switch d := indirect(v).(type) {
	case string:
		return d, true
	case bool:
		return strconv.FormatBool(d), true
	case int:
		return strconv.FormatInt(int64(d), 10), true
	case int64:
		return strconv.FormatInt(d, 10), true
	case int32:
		return strconv.FormatInt(int64(d), 10), true
	case int16:
		return strconv.FormatInt(int64(d), 10), true
	case int8:
		return strconv.FormatInt(int64(d), 10), true
	case uint:
		return strconv.FormatUint(uint64(d), 10), true
	case uint64:
		return strconv.FormatUint(d, 10), true
	case uint32:
		return strconv.FormatUint(uint64(d), 10), true
	case uint16:
		return strconv.FormatUint(uint64(d), 10), true
	case uint8:
		return strconv.FormatUint(uint64(d), 10), true
	case float64:
		return strconv.FormatFloat(d, 'f', -1, 64), true
	case float32:
		return strconv.FormatFloat(float64(d), 'f', -1, 64), true
	case []byte:
		return string(d), true
	default:
		return "", false
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

// AsInt to convert as a int64
// <!> for now, no overflow detection
func AsInt(v interface{}) (int64, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		return int64(d), true
	case int64:
		return d, true
	case int32:
		return int64(d), true
	case int16:
		return int64(d), true
	case int8:
		return int64(d), true
	case uint:
		return int64(d), true
	case uint64:
		return int64(d), true
	case uint32:
		return int64(d), true
	case uint16:
		return int64(d), true
	case uint8:
		return int64(d), true
	case float64:
		return int64(d), true
	case float32:
		return int64(d), true
	case string:
		if n, err := strconv.ParseInt(d, 10, 64); err == nil {
			return n, true
		}
		return 0, false
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

// AsUint to convert as a uint64
// <!> for now, no overflow detection
func AsUint(v interface{}) (uint64, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		if d >= 0 {
			return uint64(d), true
		}
		return 0, false
	case int64:
		if d >= 0 {
			return uint64(d), true
		}
		return 0, false
	case int32:
		if d >= 0 {
			return uint64(d), true
		}
		return 0, false
	case int16:
		if d >= 0 {
			return uint64(d), true
		}
		return 0, false
	case int8:
		if d >= 0 {
			return uint64(d), true
		}
		return 0, false
	case uint:
		return uint64(d), true
	case uint64:
		return d, true
	case uint32:
		return uint64(d), true
	case uint16:
		return uint64(d), true
	case uint8:
		return uint64(d), true
	case float64:
		if d >= 0 {
			return uint64(d), true
		}
		return 0, false
	case float32:
		if d >= 0 {
			return uint64(d), true
		}
		return 0, false
	case string:
		if n, err := strconv.ParseUint(d, 10, 64); err == nil {
			return n, true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsUintArray to convert as an array of uint64
func AsUintArray(values ...interface{}) ([]uint64, bool) {
	arr := make([]uint64, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsUint(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}

// AsFloat to convert as a float64
func AsFloat(v interface{}) (float64, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		return float64(d), true
	case int64:
		return float64(d), true
	case int32:
		return float64(d), true
	case int16:
		return float64(d), true
	case int8:
		return float64(d), true
	case uint:
		return float64(d), true
	case uint64:
		return float64(d), true
	case uint32:
		return float64(d), true
	case uint16:
		return float64(d), true
	case uint8:
		return float64(d), true
	case float64:
		return float64(d), true
	case float32:
		return float64(d), true
	case string:
		if n, err := strconv.ParseFloat(d, 64); err == nil {
			return n, true
		}
		return 0, false
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
func AsDatetime(v interface{}, format ...string) (time.Time, bool) {
	switch d := indirect(v).(type) {
	case time.Time:
		return d.UTC(), true

	// timestamp compliant with javascript
	case int:
		return time.Unix(0, int64(d)*int64(time.Millisecond)).UTC(), true
	case int64:
		return time.Unix(0, d*int64(time.Millisecond)).UTC(), true
	case int32:
		return time.Unix(0, int64(d)*int64(time.Millisecond)).UTC(), true
	case float64:
		return time.Unix(0, int64(d)*int64(time.Millisecond)).UTC(), true

	case string:
		// Try formats
		format = append(format, dateFormats...)
		for _, f := range format {
			if dt, err := time.Parse(f, d); err == nil {
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

// AsDuration to convert as a duration
func AsDuration(v interface{}) (time.Duration, bool) {
	switch d := indirect(v).(type) {
	case time.Duration:
		return d, true
	case int:
		return time.Duration(int64(d)), true
	case int64:
		return time.Duration(d), true
	case int32:
		return time.Duration(int64(d)), true
	case int16:
		return time.Duration(int64(d)), true
	case int8:
		return time.Duration(int64(d)), true
	case uint:
		return time.Duration(int64(d)), true
	case uint64:
		return time.Duration(int64(d)), true
	case uint32:
		return time.Duration(int64(d)), true
	case uint16:
		return time.Duration(int64(d)), true
	case uint8:
		return time.Duration(int64(d)), true
	case float64:
		return time.Duration(int64(d)), true
	case float32:
		return time.Duration(int64(d)), true
	case string:
		if du, err := time.ParseDuration(d); err == nil {
			return du, true
		}
		return time.Duration(0), false
	default:
		return time.Duration(0), false
	}
}

// AsDurationArray to convert as an array of duration
func AsDurationArray(values ...interface{}) ([]time.Duration, bool) {
	arr := make([]time.Duration, len(values))
	b := true
	for i, v := range values {
		if cv, ok := AsDuration(v); ok {
			arr[i] = cv
			continue
		}
		b = false
	}
	return arr, b
}
