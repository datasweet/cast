package cast_test

import (
	"testing"

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
	testBool(t, "yes", true, true)
	testBool(t, "wrong", false, false)
}

func TestAsBoolSlice(t *testing.T) {
	arr, ok := cast.AsBoolSlice(true, 1, "false")
	assert.True(t, ok)
	assert.Equal(t, []bool{true, true, false}, arr)

	arr, ok = cast.AsBoolSlice(true, 1, "false", "wrong")
	assert.False(t, ok)
	assert.Equal(t, []bool{true, true, false, false}, arr)
}

func testBool(t *testing.T, value interface{}, expected bool, ok bool) {
	b, o := cast.AsBool(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}
