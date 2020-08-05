package cast

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
		switch d {
		case "1", "t", "T", "y", "Y", "on", "true", "yes", "ON", "TRUE", "YES", "On", "True", "Yes":
			return true, true
		case "0", "f", "F", "n", "N", "false", "no", "FALSE", "NO", "False", "No":
			return false, true
		default:
			return false, false
		}
	default:
		return false, false
	}
}

// AsBoolSlice to convert as a slice of bool
func AsBoolSlice(values ...interface{}) ([]bool, bool) {
	arr := make([]bool, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsBool(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}
