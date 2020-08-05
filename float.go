package cast

import "strconv"

// AsFloat64 to convert as a float64
func AsFloat64(v interface{}) (float64, bool) {
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
		return d, true
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

// AsFloat64Slice to convert as a slice of float64
func AsFloat64Slice(values ...interface{}) ([]float64, bool) {
	arr := make([]float64, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsFloat64(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}

// AsFloat32 to convert as a float32
func AsFloat32(v interface{}) (float32, bool) {
	switch d := indirect(v).(type) {
	case bool:
		if d {
			return 1, true
		}
		return 0, true
	case int:
		return float32(d), true
	case int64:
		return float32(d), true
	case int32:
		return float32(d), true
	case int16:
		return float32(d), true
	case int8:
		return float32(d), true
	case uint:
		return float32(d), true
	case uint64:
		return float32(d), true
	case uint32:
		return float32(d), true
	case uint16:
		return float32(d), true
	case uint8:
		return float32(d), true
	case float64:
		return float32(d), true
	case float32:
		return d, true
	case string:
		if n, err := strconv.ParseFloat(d, 32); err == nil {
			return float32(n), true
		}
		return 0, false
	default:
		return 0, false
	}
}

// AsFloat32Slice to convert as a slice of float32
func AsFloat32Slice(values ...interface{}) ([]float32, bool) {
	arr := make([]float32, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsFloat32(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}
