package cast_test

import (
	"testing"
	"time"

	"github.com/datasweet/cast"
	"github.com/stretchr/testify/assert"
)

func TestAsTime(t *testing.T) {
	date := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)           // only date
	datetime := time.Date(2019, time.March, 1, 10, 13, 40, 0, time.UTC)    // date + time
	timestamp := time.Unix(0, 1551435220270*int64(time.Millisecond)).UTC() // date + time + ns

	testTime(t, timestamp, timestamp, true)
	testTime(t, 1551435220270, timestamp, true)
	testTime(t, int64(1551435220270), timestamp, true)
	testTime(t, "1551435220270", timestamp, true)
	testTime(t, "2019-03-01", date, true)
	testTime(t, "2019-03-01 10:13:40", datetime, true)
	testTime(t, "2019-03-01T10:13:40.27Z", timestamp, true)
	testTime(t, "01/03/2019", date, true)
	testTime(t, "01/03/2019 10:13:40", datetime, true)
	testTime(t, "20190301", date, true)
	testTime(t, "20190301101340", datetime, true)
	testTime(t, "wrong", time.Time{}, false)
}

func TestAsTimeSlice(t *testing.T) {
	date := time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC)           // only date
	datetime := time.Date(2019, time.March, 1, 10, 13, 40, 0, time.UTC)    // date + time
	timestamp := time.Unix(0, 1551435220270*int64(time.Millisecond)).UTC() // date + time + ns

	arr, ok := cast.AsTimeSlice(1551435220270, "01/03/2019", "20190301101340")
	assert.True(t, ok)
	assert.Equal(t, []time.Time{timestamp, date, datetime}, arr)

	arr, ok = cast.AsTimeSlice(1551435220270, "01/03/2019", "20190301101340", "wrong")
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

func TestAsDurationSlice(t *testing.T) {
	duration := time.Second * 10

	arr, ok := cast.AsDurationSlice(duration, 10000000000, "10s")
	assert.True(t, ok)
	assert.Equal(t, []time.Duration{duration, duration, duration}, arr)

	arr, ok = cast.AsDurationSlice(duration, 10000000000, "10s", "wrong")
	assert.False(t, ok)
	assert.Equal(t, []time.Duration{duration, duration, duration, time.Duration(0)}, arr)
}

func testTime(t *testing.T, value interface{}, expected time.Time, ok bool) {
	b, o := cast.AsTime(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}

func testDuration(t *testing.T, value interface{}, expected time.Duration, ok bool) {
	b, o := cast.AsDuration(value)
	assert.Equal(t, expected, b)
	assert.Equal(t, ok, o)
}
