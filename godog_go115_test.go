// +build go1.15

package clockdog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClock_addError(t *testing.T) {
	t.Parallel()

	c := New()

	err := c.add("foobar")
	expectedError := `time: invalid duration "foobar"`

	assert.EqualError(t, err, expectedError)
}
