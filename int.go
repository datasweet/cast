package cast

import "strconv"

// <!> for now, no overflow detection

// AsInt to convert as a int
func AsInt(v interface{}) (int, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		return d, true
	case int64:
		return int(d), true
	case int32:
		return int(d), true
	case int16:
		return int(d), true
	case int8:
		return int(d), true
	case uint:
		return int(d), true
	case uint64:
		return int(d), true
	case uint32:
		return int(d), true
	case uint16:
		return int(d), true
	case uint8:
		return int(d), true
	case float64:
		return int(d), true
	case float32:
		return int(d), true
	case string:
		if n, err := strconv.ParseInt(d, 10, 64); err == nil {
			return int(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsIntSlice to convert as a slice of int
func AsIntSlice(values ...interface{}) ([]int, bool) {
	arr := make([]int, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsInt(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsInt64 to convert as a int64
func AsInt64(v interface{}) (int64, bool) {
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

// AsInt64Slice to convert as a slice of int64
func AsInt64Slice(values ...interface{}) ([]int64, bool) {
	arr := make([]int64, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsInt64(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsInt32 to convert as a int32
func AsInt32(v interface{}) (int32, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		return int32(d), true
	case int64:
		return int32(d), true
	case int32:
		return d, true
	case int16:
		return int32(d), true
	case int8:
		return int32(d), true
	case uint:
		return int32(d), true
	case uint64:
		return int32(d), true
	case uint32:
		return int32(d), true
	case uint16:
		return int32(d), true
	case uint8:
		return int32(d), true
	case float64:
		return int32(d), true
	case float32:
		return int32(d), true
	case string:
		if n, err := strconv.ParseInt(d, 10, 32); err == nil {
			return int32(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsInt32Slice to convert as a slice of int32
func AsInt32Slice(values ...interface{}) ([]int32, bool) {
	arr := make([]int32, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsInt32(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsInt16 to convert as a int16
func AsInt16(v interface{}) (int16, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		return int16(d), true
	case int64:
		return int16(d), true
	case int32:
		return int16(d), true
	case int16:
		return d, true
	case int8:
		return int16(d), true
	case uint:
		return int16(d), true
	case uint64:
		return int16(d), true
	case uint32:
		return int16(d), true
	case uint16:
		return int16(d), true
	case uint8:
		return int16(d), true
	case float64:
		return int16(d), true
	case float32:
		return int16(d), true
	case string:
		if n, err := strconv.ParseInt(d, 10, 16); err == nil {
			return int16(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsInt16Slice to convert as a slice of int16
func AsInt16Slice(values ...interface{}) ([]int16, bool) {
	arr := make([]int16, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsInt16(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsInt8 to convert as a int8
func AsInt8(v interface{}) (int8, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		return int8(d), true
	case int64:
		return int8(d), true
	case int32:
		return int8(d), true
	case int16:
		return int8(d), true
	case int8:
		return d, true
	case uint:
		return int8(d), true
	case uint64:
		return int8(d), true
	case uint32:
		return int8(d), true
	case uint16:
		return int8(d), true
	case uint8:
		return int8(d), true
	case float64:
		return int8(d), true
	case float32:
		return int8(d), true
	case string:
		if n, err := strconv.ParseInt(d, 10, 8); err == nil {
			return int8(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsInt8Slice to convert as a slice of int8
func AsInt8Slice(values ...interface{}) ([]int8, bool) {
	arr := make([]int8, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsInt8(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}
