package cast_test

import (
	"testing"

	"github.com/datasweet/cast"
	"github.com/stretchr/testify/assert"
)

func TestAsUint(t *testing.T) {
	testUint(t, nil, 0, false)
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

func TestAsUintSlice(t *testing.T) {
	arr, ok := cast.AsUintSlice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []uint{1, 2, 1}, arr)

	arr, ok = cast.AsUintSlice(1, 2, true, "no", -2)
	assert.False(t, ok)
	assert.Equal(t, []uint{1, 2, 1, 0, 0}, arr)
}

func TestAsUint64(t *testing.T) {
	testUint64(t, nil, 0, false)
	testUint64(t, "123", 123, true)
	testUint64(t, "-123", 0, false)
	testUint64(t, "wrong", 0, false)
	testUint64(t, true, 1, true)
	testUint64(t, false, 0, true)
	testUint64(t, 123, 123, true)
	testUint64(t, int8(123), 123, true)
	testUint64(t, int8(-123), 0, false)
	testUint64(t, int16(123), 123, true)
	testUint64(t, int16(-123), 0, false)
	testUint64(t, int32(123), 123, true)
	testUint64(t, int32(-123), 0, false)
	testUint64(t, int64(123), 123, true)
	testUint64(t, int64(-123), 0, false)
	testUint64(t, uint(123), 123, true)
	testUint64(t, uint8(123), 123, true)
	testUint64(t, uint16(123), 123, true)
	testUint64(t, uint32(123), 123, true)
	testUint64(t, uint64(123), 123, true)
	testUint64(t, float32(123), 123, true)
	testUint64(t, float32(-123), 0, false)
	testUint64(t, float64(123), 123, true)
	testUint64(t, float64(-123), 0, false)
}

func TestAsUint64Slice(t *testing.T) {
	arr, ok := cast.AsUint64Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []uint64{1, 2, 1}, arr)

	arr, ok = cast.AsUint64Slice(1, 2, true, "no", -2)
	assert.False(t, ok)
	assert.Equal(t, []uint64{1, 2, 1, 0, 0}, arr)
}

func TestAsUint32(t *testing.T) {
	testUint32(t, nil, 0, false)
	testUint32(t, "123", 123, true)
	testUint32(t, "-123", 0, false)
	testUint32(t, "wrong", 0, false)
	testUint32(t, true, 1, true)
	testUint32(t, false, 0, true)
	testUint32(t, 123, 123, true)
	testUint32(t, int8(123), 123, true)
	testUint32(t, int8(-123), 0, false)
	testUint32(t, int16(123), 123, true)
	testUint32(t, int16(-123), 0, false)
	testUint32(t, int32(123), 123, true)
	testUint32(t, int32(-123), 0, false)
	testUint32(t, int64(123), 123, true)
	testUint32(t, int64(-123), 0, false)
	testUint32(t, uint(123), 123, true)
	testUint32(t, uint8(123), 123, true)
	testUint32(t, uint16(123), 123, true)
	testUint32(t, uint32(123), 123, true)
	testUint32(t, uint64(123), 123, true)
	testUint32(t, float32(123), 123, true)
	testUint32(t, float32(-123), 0, false)
	testUint32(t, float64(123), 123, true)
	testUint32(t, float64(-123), 0, false)
}

func TestAsUint32Slice(t *testing.T) {
	arr, ok := cast.AsUint32Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []uint32{1, 2, 1}, arr)

	arr, ok = cast.AsUint32Slice(1, 2, true, "no", -2)
	assert.False(t, ok)
	assert.Equal(t, []uint32{1, 2, 1, 0, 0}, arr)
}

func TestAsUint16(t *testing.T) {
	testUint16(t, nil, 0, false)
	testUint16(t, "123", 123, true)
	testUint16(t, "-123", 0, false)
	testUint16(t, "wrong", 0, false)
	testUint16(t, true, 1, true)
	testUint16(t, false, 0, true)
	testUint16(t, 123, 123, true)
	testUint16(t, int8(123), 123, true)
	testUint16(t, int8(-123), 0, false)
	testUint16(t, int16(123), 123, true)
	testUint16(t, int16(-123), 0, false)
	testUint16(t, int32(123), 123, true)
	testUint16(t, int32(-123), 0, false)
	testUint16(t, int64(123), 123, true)
	testUint16(t, int64(-123), 0, false)
	testUint16(t, uint(123), 123, true)
	testUint16(t, uint8(123), 123, true)
	testUint16(t, uint16(123), 123, true)
	testUint16(t, uint32(123), 123, true)
	testUint16(t, uint64(123), 123, true)
	testUint16(t, float32(123), 123, true)
	testUint16(t, float32(-123), 0, false)
	testUint16(t, float64(123), 123, true)
	testUint16(t, float64(-123), 0, false)
}

func TestAsUint16Slice(t *testing.T) {
	arr, ok := cast.AsUint16Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []uint16{1, 2, 1}, arr)

	arr, ok = cast.AsUint16Slice(1, 2, true, "no", -2)
	assert.False(t, ok)
	assert.Equal(t, []uint16{1, 2, 1, 0, 0}, arr)
}

func TestAsUint8(t *testing.T) {
	testUint8(t, nil, 0, false)
	testUint8(t, "123", 123, true)
	testUint8(t, "-123", 0, false)
	testUint8(t, "wrong", 0, false)
	testUint8(t, true, 1, true)
	testUint8(t, false, 0, true)
	testUint8(t, 123, 123, true)
	testUint8(t, int8(123), 123, true)
	testUint8(t, int8(-123), 0, false)
	testUint8(t, int16(123), 123, true)
	testUint8(t, int16(-123), 0, false)
	testUint8(t, int32(123), 123, true)
	testUint8(t, int32(-123), 0, false)
	testUint8(t, int64(123), 123, true)
	testUint8(t, int64(-123), 0, false)
	testUint8(t, uint(123), 123, true)
	testUint8(t, uint8(123), 123, true)
	testUint8(t, uint16(123), 123, true)
	testUint8(t, uint32(123), 123, true)
	testUint8(t, uint64(123), 123, true)
	testUint8(t, float32(123), 123, true)
	testUint8(t, float32(-123), 0, false)
	testUint8(t, float64(123), 123, true)
	testUint8(t, float64(-123), 0, false)
}

func TestAsUint8Slice(t *testing.T) {
	arr, ok := cast.AsUint8Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []uint8{1, 2, 1}, arr)

	arr, ok = cast.AsUint8Slice(1, 2, true, "no", -2)
	assert.False(t, ok)
	assert.Equal(t, []uint8{1, 2, 1, 0, 0}, arr)
}

func testUint(t *testing.T, value interface{}, expected uint, ok bool) {
	b, o := cast.AsUint(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testUint64(t *testing.T, value interface{}, expected uint64, ok bool) {
	b, o := cast.AsUint64(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testUint32(t *testing.T, value interface{}, expected uint32, ok bool) {
	b, o := cast.AsUint32(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testUint16(t *testing.T, value interface{}, expected uint16, ok bool) {
	b, o := cast.AsUint16(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testUint8(t *testing.T, value interface{}, expected uint8, ok bool) {
	b, o := cast.AsUint8(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}
