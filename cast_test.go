package cast_test

import (
	"testing"
	"time"

	"github.com/datasweet/cast"
	"github.com/stretchr/testify/assert"
)

func TestAsBool(t *testing.T) {
	testBool(t, nil, false, false)
	testBool(t, false, false, true)
	testBool(t, true, true, true)
	testBool(t, 1, true, true)
	testBool(t, int8(1), true, true)
	testBool(t, int16(1), true, true)
	testBool(t, int32(1), true, true)
	testBool(t, int64(1), true, true)
	testBool(t, uint(1), true, true)
	testBool(t, uint8(1), true, true)
	testBool(t, uint16(1), true, true)
	testBool(t, uint32(1), true, true)
	testBool(t, uint64(1), true, true)
	testBool(t, float32(1), true, true)
	testBool(t, float64(1), true, true)
	testBool(t, "true", true, true)
	testBool(t, "wrong", false, false)
}

func TestAsBoolArray(t *testing.T) {
	arr, ok := cast.AsBoolArray(true, 1, "false")
	assert.True(t, ok)
	assert.Equal(t, []bool{true, true, false}, arr)

	arr, ok = cast.AsBoolArray(true, 1, "false", "wrong")
	assert.False(t, ok)
	assert.Equal(t, []bool{true, true, false, false}, arr)
}

func TestAsString(t *testing.T) {
	testString(t, nil, "", false)
	testString(t, "hello", "hello", true)
	testString(t, 123, "123", true)
	testString(t, int8(123), "123", true)
	testString(t, int16(123), "123", true)
	testString(t, int32(123), "123", true)
	testString(t, int64(123), "123", true)
	testString(t, uint(123), "123", true)
	testString(t, uint8(123), "123", true)
	testString(t, uint16(123), "123", true)
	testString(t, uint32(123), "123", true)
	testString(t, uint64(123), "123", true)
	testString(t, float32(3.1), "3.0999999046325684", true)
	testString(t, float64(3.14), "3.14", true)
	testString(t, true, "true", true)
}

func TestAsStringArray(t *testing.T) {
	arr, ok := cast.AsStringArray("hello", 123, false)
	assert.True(t, ok)
	assert.Equal(t, []string{"hello", "123", "false"}, arr)

	arr, ok = cast.AsStringArray("hello", 123, false, nil)
	assert.False(t, ok)
	assert.Equal(t, []string{"hello", "123", "false", ""}, arr)
}

func TestAsInt(t *testing.T) {
	testInt(t, "123", 123, true)
	testInt(t, "wrong", 0, false)
	testInt(t, true, 1, true)
	testInt(t, false, 0, true)
	testInt(t, 123, 123, true)
	testInt(t, int8(123), 123, true)
	testInt(t, int16(123), 123, true)
	testInt(t, int32(123), 123, true)
	testInt(t, int64(123), 123, true)
	testInt(t, uint(123), 123, true)
	testInt(t, uint8(123), 123, true)
	testInt(t, uint16(123), 123, true)
	testInt(t, uint32(123), 123, true)
	testInt(t, uint64(123), 123, true)
	testInt(t, float32(123), 123, true)
	testInt(t, float64(123), 123, true)
}

func TestAsUintArray(t *testing.T) {
	arr, ok := cast.AsUintArray(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []uint64{1, 2, 1}, arr)

	arr, ok = cast.AsUintArray(1, 2, true, "no", -2)
	assert.False(t, ok)
	assert.Equal(t, []uint64{1, 2, 1, 0, 0}, arr)
}

func TestAsUint(t *testing.T) {
	testUint(t, "123", 123, true)
	testUint(t, "-123", 0, false)
	testUint(t, "wrong", 0, false)
	testUint(t, true, 1, true)
	testUint(t, false, 0, true)
	testUint(t, 123, 123, true)
	testUint(t, int8(123), 123, true)
	testUint(t, int8(-123), 0, false)
	testUint(t, int16(123), 123, true)
	testUint(t, int16(-123), 0, false)
	testUint(t, int32(123), 123, true)
	testUint(t, int32(-123), 0, false)
	testUint(t, int64(123), 123, true)
	testUint(t, int64(-123), 0, false)
	testUint(t, uint(123), 123, true)
	testUint(t, uint8(123), 123, true)
	testUint(t, uint16(123), 123, true)
	testUint(t, uint32(123), 123, true)
	testUint(t, uint64(123), 123, true)
	testUint(t, float32(123), 123, true)
	testUint(t, float32(-123), 0, false)
	testUint(t, float64(123), 123, true)
	testUint(t, float64(-123), 0, false)
}

func TestAsIntArray(t *testing.T) {
	arr, ok := cast.AsIntArray(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []int64{1, 2, 1}, arr)

	arr, ok = cast.AsIntArray(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []int64{1, 2, 1, 0}, arr)
}

func TestAsFloat(t *testing.T) {
	testFloat(t, "123", 123, true)
	testFloat(t, "wrong", 0, false)
	testFloat(t, true, 1, true)
	testFloat(t, false, 0, true)
	testFloat(t, 123, 123, true)
	testFloat(t, int8(123), 123, true)
	testFloat(t, int16(123), 123, true)
	testFloat(t, int32(123), 123, true)
	testFloat(t, int64(123), 123, true)
	testFloat(t, uint(123), 123, true)
	testFloat(t, uint8(123), 123, true)
	testFloat(t, uint16(123), 123, true)
	testFloat(t, uint32(123), 123, true)
	testFloat(t, uint64(123), 123, true)
	testFloat(t, float32(123), 123, true)
	testFloat(t, float64(123), 123, true)
}

func TestAsFloatArray(t *testing.T) {
	arr, ok := cast.AsFloatArray(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []float64{1, 2, 1}, arr)

	arr, ok = cast.AsFloatArray(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []float64{1, 2, 1, 0}, arr)
}

func TestAsDatetime(t *testing.T) {
	date := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)           // only date
	datetime := time.Date(2019, time.March, 1, 10, 13, 40, 0, time.UTC)    // date + time
	timestamp := time.Unix(0, 1551435220270*int64(time.Millisecond)).UTC() // date + time + ns

	testDatetime(t, timestamp, timestamp, true)
	testDatetime(t, 1551435220270, timestamp, true)
	testDatetime(t, int64(1551435220270), timestamp, true)
	testDatetime(t, "1551435220270", timestamp, true)
	testDatetime(t, "2019-03-01", date, true)
	testDatetime(t, "2019-03-01 10:13:40", datetime, true)
	testDatetime(t, "2019-03-01T10:13:40.27Z", timestamp, true)
	testDatetime(t, "01/03/2019", date, true)
	testDatetime(t, "01/03/2019 10:13:40", datetime, true)
	testDatetime(t, "20190301", date, true)
	testDatetime(t, "20190301101340", datetime, true)
	testDatetime(t, "wrong", time.Time{}, false)
}

func TestAsDatetimeArray(t *testing.T) {
	date := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)           // only date
	datetime := time.Date(2019, time.March, 1, 10, 13, 40, 0, time.UTC)    // date + time
	timestamp := time.Unix(0, 1551435220270*int64(time.Millisecond)).UTC() // date + time + ns

	arr, ok := cast.AsDatetimeArray(1551435220270, "01/03/2019", "20190301101340")
	assert.True(t, ok)
	assert.Equal(t, []time.Time{timestamp, date, datetime}, arr)

	arr, ok = cast.AsDatetimeArray(1551435220270, "01/03/2019", "20190301101340", "wrong")
	assert.False(t, ok)
	assert.Equal(t, []time.Time{timestamp, date, datetime, {}}, arr)
}

func TestAsDuration(t *testing.T) {
	duration := time.Second * 10

	testDuration(t, duration, duration, true)
	testDuration(t, 10000000000, duration, true)
	testDuration(t, "10s", duration, true)
	testDuration(t, "wrong", time.Duration(0), false)
}

func TestAsDurationArray(t *testing.T) {
	duration := time.Second * 10

	arr, ok := cast.AsDurationArray(duration, 10000000000, "10s")
	assert.True(t, ok)
	assert.Equal(t, []time.Duration{duration, duration, duration}, arr)

	arr, ok = cast.AsDurationArray(duration, 10000000000, "10s", "wrong")
	assert.False(t, ok)
	assert.Equal(t, []time.Duration{duration, duration, duration, time.Duration(0)}, arr)
}

func testBool(t *testing.T, value interface{}, expected bool, ok bool) {
	b, o := cast.AsBool(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testString(t *testing.T, value interface{}, expected string, ok bool) {
	b, o := cast.AsString(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testInt(t *testing.T, value interface{}, expected int64, ok bool) {
	b, o := cast.AsInt(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testUint(t *testing.T, value interface{}, expected uint64, ok bool) {
	b, o := cast.AsUint(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testFloat(t *testing.T, value interface{}, expected float64, ok bool) {
	b, o := cast.AsFloat(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testDatetime(t *testing.T, value interface{}, expected time.Time, ok bool) {
	b, o := cast.AsDatetime(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testDuration(t *testing.T, value interface{}, expected time.Duration, ok bool) {
	b, o := cast.AsDuration(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}
