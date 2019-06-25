package cast

import (
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

func valueOf(v interface{}) reflect.Value {
	if v == nil {
		return reflect.Value{}
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return rv
	}
	return rv.Elem()
}

// AsBool to convert as bool
func AsBool(v interface{}) (bool, bool) {
	rv := valueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() > 0, true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() > 0, true
	case reflect.Float32, reflect.Float64:
		return rv.Float() > 0, true
	case reflect.String:
		if b, err := strconv.ParseBool(rv.String()); err == nil {
			return b, true
		}
		return false, false
	case reflect.Bool:
		return rv.Bool(), true
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
	rv := valueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10), true
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'f', -1, 64), true
	case reflect.String:
		return rv.String(), true
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool()), true
	case reflect.Invalid:
		return "", false
	default:
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
	rv := valueOf(v)

	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(rv.Uint()), true
	case reflect.Float32, reflect.Float64:
		return int64(rv.Float()), true
	case reflect.String:
		if n, err := strconv.ParseInt(rv.String(), 10, 64); err == nil {
			return n, true
		}
		return 0, false
	case reflect.Bool:
		if n := rv.Bool(); n {
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

// AsUint to convert as a uint64
func AsUint(v interface{}) (uint64, bool) {
	rv := valueOf(v)

	switch kind := rv.Kind(); kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n := rv.Int()
		if n < 0 {
			return 0, false
		}
		return uint64(n), true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint64(rv.Uint()), true

	case reflect.Float32, reflect.Float64:
		n := reflect.ValueOf(v).Float()
		if n < 0 {
			return 0, false
		}
		return uint64(n), true

	case reflect.String:
		if n, err := strconv.ParseUint(rv.String(), 10, 64); err == nil {
			return n, true
		}
		return 0, false

	case reflect.Bool:
		n := rv.Bool()
		if n {
			return 1, true
		}
		return 0, true

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
	rv := valueOf(v)

	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(rv.Uint()), true
	case reflect.Float32, reflect.Float64:
		return rv.Float(), true
	case reflect.String:
		if n, err := strconv.ParseFloat(rv.String(), 64); err == nil {
			return n, true
		}
		return 0, false
	case reflect.Bool:
		if n := rv.Bool(); n {
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

	rv := valueOf(v)

	switch rv.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		// timestamp compliant with javascript
		return time.Unix(0, rv.Int()*int64(time.Millisecond)).UTC(), true

	case reflect.String:
		str := rv.String()
		// Try formats
		for _, format := range DateFormats {
			if dt, err := time.Parse(format, str); err == nil {
				return dt.UTC(), true
			}
		}

		// Try to convert to int64
		if ts, err := strconv.ParseInt(str, 10, 64); err == nil {
			return time.Unix(0, ts*int64(time.Millisecond)).UTC(), true
		}

		return time.Time{}, false

	case reflect.Struct:
		if c, ok := v.(time.Time); ok {
			return c.UTC(), true
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
	rv := valueOf(v)

	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return time.Duration(rv.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return time.Duration(int64(rv.Uint())), true
	case reflect.Float32, reflect.Float64:
		return time.Duration(int64(rv.Float())), true

	case reflect.String:
		if du, err := time.ParseDuration(rv.String()); err == nil {
			return du, true
		}
		return time.Duration(0), false
	case reflect.Struct:
		if du, ok := v.(time.Duration); ok {
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
