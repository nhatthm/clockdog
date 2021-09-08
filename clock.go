package clockdog

import (
	"github.com/godogx/clocksteps"
)

// ErrClockIsNotSet indicates that the clock must be set by either Clock.Set() or Clock.Freeze() before adding some
// time.Duration into it.
//
// Deprecated: Use clocksteps.ErrClockIsNotSet instead.
var ErrClockIsNotSet = clocksteps.ErrClockIsNotSet

// Clock is a clock.Clock.
//
// Deprecated: Use clocksteps.Clock instead.
type Clock = clocksteps.Clock

// New initiates a new Clock.
//
// Deprecated: Use clocksteps.New instead.
func New() *Clock {
	return clocksteps.New()
}
