package cast_test

import (
	"testing"

	"github.com/datasweet/cast"
	"github.com/stretchr/testify/assert"
)

func TestAsInt(t *testing.T) {
	testInt(t, nil, 0, false)
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

func TestAsIntSlice(t *testing.T) {
	arr, ok := cast.AsIntSlice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []int{1, 2, 1}, arr)

	arr, ok = cast.AsIntSlice(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []int{1, 2, 1, 0}, arr)
}

func TestAsInt64(t *testing.T) {
	testInt64(t, nil, 0, false)
	testInt64(t, "123", 123, true)
	testInt64(t, "wrong", 0, false)
	testInt64(t, true, 1, true)
	testInt64(t, false, 0, true)
	testInt64(t, 123, 123, true)
	testInt64(t, int8(123), 123, true)
	testInt64(t, int16(123), 123, true)
	testInt64(t, int32(123), 123, true)
	testInt64(t, int64(123), 123, true)
	testInt64(t, uint(123), 123, true)
	testInt64(t, uint8(123), 123, true)
	testInt64(t, uint16(123), 123, true)
	testInt64(t, uint32(123), 123, true)
	testInt64(t, uint64(123), 123, true)
	testInt64(t, float32(123), 123, true)
	testInt64(t, float64(123), 123, true)
}

func TestAsInt64Slice(t *testing.T) {
	arr, ok := cast.AsInt64Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []int64{1, 2, 1}, arr)

	arr, ok = cast.AsInt64Slice(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []int64{1, 2, 1, 0}, arr)
}

func TestAsInt32(t *testing.T) {
	testInt32(t, nil, 0, false)
	testInt32(t, "123", 123, true)
	testInt32(t, "wrong", 0, false)
	testInt32(t, true, 1, true)
	testInt32(t, false, 0, true)
	testInt32(t, 123, 123, true)
	testInt32(t, int8(123), 123, true)
	testInt32(t, int16(123), 123, true)
	testInt32(t, int32(123), 123, true)
	testInt32(t, int64(123), 123, true)
	testInt32(t, uint(123), 123, true)
	testInt32(t, uint8(123), 123, true)
	testInt32(t, uint16(123), 123, true)
	testInt32(t, uint32(123), 123, true)
	testInt32(t, uint64(123), 123, true)
	testInt32(t, float32(123), 123, true)
	testInt32(t, float64(123), 123, true)
}

func TestAsInt32Slice(t *testing.T) {
	arr, ok := cast.AsInt32Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []int32{1, 2, 1}, arr)

	arr, ok = cast.AsInt32Slice(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []int32{1, 2, 1, 0}, arr)
}

func TestAsInt16(t *testing.T) {
	testInt16(t, nil, 0, false)
	testInt16(t, "123", 123, true)
	testInt16(t, "wrong", 0, false)
	testInt16(t, true, 1, true)
	testInt16(t, false, 0, true)
	testInt16(t, 123, 123, true)
	testInt16(t, int8(123), 123, true)
	testInt16(t, int16(123), 123, true)
	testInt16(t, int32(123), 123, true)
	testInt16(t, int64(123), 123, true)
	testInt16(t, uint(123), 123, true)
	testInt16(t, uint8(123), 123, true)
	testInt16(t, uint16(123), 123, true)
	testInt16(t, uint32(123), 123, true)
	testInt16(t, uint64(123), 123, true)
	testInt16(t, float32(123), 123, true)
	testInt16(t, float64(123), 123, true)
}

func TestAsInt16Slice(t *testing.T) {
	arr, ok := cast.AsInt16Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []int16{1, 2, 1}, arr)

	arr, ok = cast.AsInt16Slice(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []int16{1, 2, 1, 0}, arr)
}

func TestAsInt8(t *testing.T) {
	testInt8(t, nil, 0, false)
	testInt8(t, "123", 123, true)
	testInt8(t, "wrong", 0, false)
	testInt8(t, true, 1, true)
	testInt8(t, false, 0, true)
	testInt8(t, 123, 123, true)
	testInt8(t, int8(123), 123, true)
	testInt8(t, int16(123), 123, true)
	testInt8(t, int32(123), 123, true)
	testInt8(t, int64(123), 123, true)
	testInt8(t, uint(123), 123, true)
	testInt8(t, uint8(123), 123, true)
	testInt8(t, uint16(123), 123, true)
	testInt8(t, uint32(123), 123, true)
	testInt8(t, uint64(123), 123, true)
	testInt8(t, float32(123), 123, true)
	testInt8(t, float64(123), 123, true)
}

func TestAsInt8Slice(t *testing.T) {
	arr, ok := cast.AsInt8Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []int8{1, 2, 1}, arr)

	arr, ok = cast.AsInt8Slice(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []int8{1, 2, 1, 0}, arr)
}

func testInt(t *testing.T, value interface{}, expected int, ok bool) {
	b, o := cast.AsInt(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testInt64(t *testing.T, value interface{}, expected int64, ok bool) {
	b, o := cast.AsInt64(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testInt32(t *testing.T, value interface{}, expected int32, ok bool) {
	b, o := cast.AsInt32(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testInt16(t *testing.T, value interface{}, expected int16, ok bool) {
	b, o := cast.AsInt16(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testInt8(t *testing.T, value interface{}, expected int8, ok bool) {
	b, o := cast.AsInt8(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}
