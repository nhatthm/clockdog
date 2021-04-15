package clockdog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClock_setError(t *testing.T) {
	t.Parallel()

	c := New()

	err := c.set("foobar")
	expectedError := `parsing time "foobar" as "2006-01-02": cannot parse "foobar" as "2006"`

	assert.EqualError(t, err, expectedError)
}

func TestClock_addDateError(t *testing.T) {
	t.Parallel()

	c := New()

	t.Run("invalid year", func(t *testing.T) {
		t.Parallel()

		err := c.addDate("foobar", "0", "0")
		expectedError := `strconv.Atoi: parsing "foobar": invalid syntax`

		assert.EqualError(t, err, expectedError)
	})

	t.Run("invalid month", func(t *testing.T) {
		t.Parallel()

		err := c.addDate("0", "foobar", "0")
		expectedError := `strconv.Atoi: parsing "foobar": invalid syntax`

		assert.EqualError(t, err, expectedError)
	})

	t.Run("invalid year", func(t *testing.T) {
		t.Parallel()

		err := c.addDate("0", "0", "foobar")
		expectedError := `strconv.Atoi: parsing "foobar": invalid syntax`

		assert.EqualError(t, err, expectedError)
	})
}
