package cast_test

import (
	"testing"

	"github.com/datasweet/cast"
	"github.com/stretchr/testify/assert"
)

func TestAsFloat64(t *testing.T) {
	testFloat64(t, nil, 0, false)
	testFloat64(t, "123", 123, true)
	testFloat64(t, "wrong", 0, false)
	testFloat64(t, true, 1, true)
	testFloat64(t, false, 0, true)
	testFloat64(t, 123, 123, true)
	testFloat64(t, int8(123), 123, true)
	testFloat64(t, int16(123), 123, true)
	testFloat64(t, int32(123), 123, true)
	testFloat64(t, int64(123), 123, true)
	testFloat64(t, uint(123), 123, true)
	testFloat64(t, uint8(123), 123, true)
	testFloat64(t, uint16(123), 123, true)
	testFloat64(t, uint32(123), 123, true)
	testFloat64(t, uint64(123), 123, true)
	testFloat64(t, float32(123), 123, true)
	testFloat64(t, float64(123), 123, true)
}

func TestAsFloat64Slice(t *testing.T) {
	arr, ok := cast.AsFloat64Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []float64{1, 2, 1}, arr)

	arr, ok = cast.AsFloat64Slice(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []float64{1, 2, 1, 0}, arr)
}

func TestAsFloat32(t *testing.T) {
	testFloat32(t, nil, 0, false)
	testFloat32(t, "123", 123, true)
	testFloat32(t, "wrong", 0, false)
	testFloat32(t, true, 1, true)
	testFloat32(t, false, 0, true)
	testFloat32(t, 123, 123, true)
	testFloat32(t, int8(123), 123, true)
	testFloat32(t, int16(123), 123, true)
	testFloat32(t, int32(123), 123, true)
	testFloat32(t, int64(123), 123, true)
	testFloat32(t, uint(123), 123, true)
	testFloat32(t, uint8(123), 123, true)
	testFloat32(t, uint16(123), 123, true)
	testFloat32(t, uint32(123), 123, true)
	testFloat32(t, uint64(123), 123, true)
	testFloat32(t, float32(123), 123, true)
	testFloat32(t, float64(123), 123, true)
}

func TestAsFloat32Slice(t *testing.T) {
	arr, ok := cast.AsFloat32Slice(1, 2, true)
	assert.True(t, ok)
	assert.Equal(t, []float32{1, 2, 1}, arr)

	arr, ok = cast.AsFloat32Slice(1, 2, true, "no")
	assert.False(t, ok)
	assert.Equal(t, []float32{1, 2, 1, 0}, arr)
}

func testFloat64(t *testing.T, value interface{}, expected float64, ok bool) {
	b, o := cast.AsFloat64(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testFloat32(t *testing.T, value interface{}, expected float32, ok bool) {
	b, o := cast.AsFloat32(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}
