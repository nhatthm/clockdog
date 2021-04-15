package clockdog_test

import (
	"testing"
	"time"

	"github.com/nhatthm/clockdog"
	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	t.Parallel()

	c := clockdog.New()

	now := time.Now()

	assert.True(t, now.Before(c.Now()))

	// Errors while adding time to a live clock.
	assert.Equal(t, clockdog.ErrClockIsNotSet, c.Add(time.Hour))
	assert.Equal(t, clockdog.ErrClockIsNotSet, c.AddDate(0, 0, 1))

	// Freeze the clock.
	c.Freeze()

	ts := c.Now()

	<-time.After(50 * time.Millisecond)

	assert.Equal(t, ts, c.Now())

	// Set to another time.
	ts = time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)

	c.Set(ts)

	<-time.After(50 * time.Millisecond)

	assert.Equal(t, ts, c.Now())

	// Change the time.
	ts = ts.Add(2 * time.Hour)
	err := c.Add(2 * time.Hour)
	assert.NoError(t, err)

	<-time.After(50 * time.Millisecond)

	assert.Equal(t, ts, c.Now())

	// Change the date.
	ts = ts.AddDate(2, 1, 3)
	err = c.AddDate(2, 1, 3)
	assert.NoError(t, err)

	<-time.After(50 * time.Millisecond)

	assert.Equal(t, ts, c.Now())

	// Unfreeze the clock.
	c.Unfreeze()

	now = time.Now()

	assert.True(t, now.Before(c.Now()))
}

func TestClock_Clock(t *testing.T) {
	t.Parallel()

	c := clockdog.New()

	assert.Equal(t, c, c.Clock())
}
