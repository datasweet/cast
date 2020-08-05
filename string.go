package cast

import "strconv"

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

// AsStringSlice to convert as a slice of string
func AsStringSlice(values ...interface{}) ([]string, bool) {
	arr := make([]string, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsString(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}
