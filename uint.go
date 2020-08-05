package cast

import "strconv"

// <!> for now, no overflow detection

// AsUint to convert as a uint
func AsUint(v interface{}) (uint, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		if d >= 0 {
			return uint(d), true
		}
		return 0, false
	case int64:
		if d >= 0 {
			return uint(d), true
		}
		return 0, false
	case int32:
		if d >= 0 {
			return uint(d), true
		}
		return 0, false
	case int16:
		if d >= 0 {
			return uint(d), true
		}
		return 0, false
	case int8:
		if d >= 0 {
			return uint(d), true
		}
		return 0, false
	case uint:
		return d, true
	case uint64:
		return uint(d), true
	case uint32:
		return uint(d), true
	case uint16:
		return uint(d), true
	case uint8:
		return uint(d), true
	case float64:
		if d >= 0 {
			return uint(d), true
		}
		return 0, false
	case float32:
		if d >= 0 {
			return uint(d), true
		}
		return 0, false
	case string:
		if n, err := strconv.ParseUint(d, 10, 64); err == nil {
			return uint(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsUintSlice to convert as a slice of uint
func AsUintSlice(values ...interface{}) ([]uint, bool) {
	arr := make([]uint, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsUint(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsUint64 to convert as a int64
func AsUint64(v interface{}) (uint64, bool) {
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

// AsUint64Slice to convert as a slice of uint64
func AsUint64Slice(values ...interface{}) ([]uint64, bool) {
	arr := make([]uint64, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsUint64(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsUint32 to convert as a uint32
func AsUint32(v interface{}) (uint32, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		if d >= 0 {
			return uint32(d), true
		}
		return 0, false
	case int64:
		if d >= 0 {
			return uint32(d), true
		}
		return 0, false
	case int32:
		if d >= 0 {
			return uint32(d), true
		}
		return 0, false
	case int16:
		if d >= 0 {
			return uint32(d), true
		}
		return 0, false
	case int8:
		if d >= 0 {
			return uint32(d), true
		}
		return 0, false
	case uint:
		return uint32(d), true
	case uint64:
		return uint32(d), true
	case uint32:
		return d, true
	case uint16:
		return uint32(d), true
	case uint8:
		return uint32(d), true
	case float64:
		if d >= 0 {
			return uint32(d), true
		}
		return 0, false
	case float32:
		if d >= 0 {
			return uint32(d), true
		}
		return 0, false
	case string:
		if n, err := strconv.ParseUint(d, 10, 32); err == nil {
			return uint32(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsUint32Slice to convert as a slice of uint32
func AsUint32Slice(values ...interface{}) ([]uint32, bool) {
	arr := make([]uint32, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsUint32(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsUint16 to convert as a int16
func AsUint16(v interface{}) (uint16, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		if d >= 0 {
			return uint16(d), true
		}
		return 0, false
	case int64:
		if d >= 0 {
			return uint16(d), true
		}
		return 0, false
	case int32:
		if d >= 0 {
			return uint16(d), true
		}
		return 0, false
	case int16:
		if d >= 0 {
			return uint16(d), true
		}
		return 0, false
	case int8:
		if d >= 0 {
			return uint16(d), true
		}
		return 0, false
	case uint:
		return uint16(d), true
	case uint64:
		return uint16(d), true
	case uint32:
		return uint16(d), true
	case uint16:
		return d, true
	case uint8:
		return uint16(d), true
	case float64:
		if d >= 0 {
			return uint16(d), true
		}
		return 0, false
	case float32:
		if d >= 0 {
			return uint16(d), true
		}
		return 0, false
	case string:
		if n, err := strconv.ParseUint(d, 10, 16); err == nil {
			return uint16(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsUint16Slice to convert as a slice of uint16
func AsUint16Slice(values ...interface{}) ([]uint16, bool) {
	arr := make([]uint16, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsUint16(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsUint8 to convert as a int8
func AsUint8(v interface{}) (uint8, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		if d >= 0 {
			return uint8(d), true
		}
		return 0, false
	case int64:
		if d >= 0 {
			return uint8(d), true
		}
		return 0, false
	case int32:
		if d >= 0 {
			return uint8(d), true
		}
		return 0, false
	case int16:
		if d >= 0 {
			return uint8(d), true
		}
		return 0, false
	case int8:
		if d >= 0 {
			return uint8(d), true
		}
		return 0, false
	case uint:
		return uint8(d), true
	case uint64:
		return uint8(d), true
	case uint32:
		return uint8(d), true
	case uint16:
		return uint8(d), true
	case uint8:
		return d, true
	case float64:
		if d >= 0 {
			return uint8(d), true
		}
		return 0, false
	case float32:
		if d >= 0 {
			return uint8(d), true
		}
		return 0, false
	case string:
		if n, err := strconv.ParseUint(d, 10, 8); err == nil {
			return uint8(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsUint8Slice to convert as a slice of uint8
func AsUint8Slice(values ...interface{}) ([]uint8, bool) {
	arr := make([]uint8, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsUint8(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}
