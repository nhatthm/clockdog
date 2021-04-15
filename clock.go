package clockdog

import (
	"errors"
	"sync"
	"time"

	clock "github.com/nhatthm/go-clock"
)

// ErrClockIsNotSet indicates that the clock must be set by either Clock.Set() or Clock.Freeze() before adding some
// time.Duration into it.
var ErrClockIsNotSet = errors.New("clock is not set")

var _ clock.Clock = (*Clock)(nil)

// Clock is a clock.Clock.
type Clock struct {
	timestamp *time.Time
	mu        sync.Mutex
}

// Now returns a fixed timestamp or time.Now().
func (c *Clock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.timestamp == nil {
		return time.Now()
	}

	return *c.timestamp
}

// Set fixes the clock at a time.
func (c *Clock) Set(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timestamp = timestamp(t)
}

// Add adds time to the clock.
func (c *Clock) Add(d time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.timestamp == nil {
		return ErrClockIsNotSet
	}

	c.timestamp = timestamp(c.timestamp.Add(d))

	return nil
}

// AddDate adds date to the clock.
func (c *Clock) AddDate(years, months, days int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.timestamp == nil {
		return ErrClockIsNotSet
	}

	c.timestamp = timestamp(c.timestamp.AddDate(years, months, days))

	return nil
}

// Freeze freezes the clock.
func (c *Clock) Freeze() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timestamp = timestamp(time.Now())
}

// Unfreeze unfreezes the clock.
func (c *Clock) Unfreeze() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.timestamp = nil
}

// Clock provides clock.Clock.
func (c *Clock) Clock() clock.Clock {
	return c
}

// New initiates a new Clock.
func New() *Clock {
	return &Clock{}
}

func timestamp(t time.Time) *time.Time {
	return &t
}
