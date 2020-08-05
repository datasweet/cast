package cast

import (
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

// AsTime to convert as time.Time
// You can pass golang format for a string conversion
func AsTime(v interface{}, format ...string) (time.Time, bool) {
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
		if len(format) == 0 {
			format = dateFormats
		}
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

// AsTimeSlice to convert as a slice of time
func AsTimeSlice(values ...interface{}) ([]time.Time, bool) {
	arr := make([]time.Time, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsTime(v)
		b = b && ok
		arr = append(arr, cv)
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

// AsDurationSlice to convert as a slice of duration
func AsDurationSlice(values ...interface{}) ([]time.Duration, bool) {
	arr := make([]time.Duration, 0, len(values))
	b := true
	for _, v := range values {
		cv, ok := AsDuration(v)
		b = b && ok
		arr = append(arr, cv)
	}
	return arr, b
}
