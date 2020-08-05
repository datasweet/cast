package cast_test

import (
	"testing"

	"github.com/datasweet/cast"
	"github.com/stretchr/testify/assert"
)

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

func TestAsStringSlice(t *testing.T) {
	arr, ok := cast.AsStringSlice("hello", 123, false)
	assert.True(t, ok)
	assert.Equal(t, []string{"hello", "123", "false"}, arr)

	arr, ok = cast.AsStringSlice("hello", 123, false, nil)
	assert.False(t, ok)
	assert.Equal(t, []string{"hello", "123", "false", ""}, arr)
}

func testString(t *testing.T, value interface{}, expected string, ok bool) {
	b, o := cast.AsString(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}
