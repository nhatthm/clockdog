package bootstrap

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/nhatthm/clockdog"
	"github.com/nhatthm/timeparser"
	"github.com/stretchr/testify/assert"
)

// Used by init().
//
//nolint:gochecknoglobals
var (
	runGoDogTests bool

	out = new(bytes.Buffer)
	opt = godog.Options{
		Strict: true,
		Output: out,
	}
)

// This has to run on init to define -godog flag, otherwise "undefined flag" error happens.
//
//nolint:gochecknoinits
func init() {
	flag.BoolVar(&runGoDogTests, "godog", false, "Set this flag is you want to run godog BDD tests")
	godog.BindCommandLineFlags("", &opt)
}

func TestIntegration(t *testing.T) {
	if !runGoDogTests {
		t.Skip(`Missing "-godog" flag, skipping integration test.`)
	}

	c := clockdog.New()

	RunSuite(t, "..", func(_ *testing.T, ctx *godog.ScenarioContext) {
		c.RegisterContext(ctx)
		registerClock(c, ctx)
	})
}

func RunSuite(t *testing.T, path string, featureContext func(t *testing.T, ctx *godog.ScenarioContext)) {
	t.Helper()

	flag.Parse()

	if opt.Randomize == 0 {
		opt.Randomize = rand.Int63() // nolint: gosec
	}

	var paths []string

	files, err := ioutil.ReadDir(filepath.Clean(path))
	assert.NoError(t, err)

	paths = make([]string, 0, len(files))

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".feature") {
			paths = append(paths, filepath.Join(path, f.Name()))
		}
	}

	for _, path := range paths {
		path := path

		t.Run(path, func(t *testing.T) {
			opt.Paths = []string{path}
			suite := godog.TestSuite{
				Name:                 "Integration",
				TestSuiteInitializer: nil,
				ScenarioInitializer: func(s *godog.ScenarioContext) {
					featureContext(t, s)
				},
				Options: &opt,
			}
			status := suite.Run()

			if status != 0 {
				fmt.Println(out.String())
				assert.Fail(t, "one or more scenarios failed in feature: "+path)
			}
		})
	}
}

func registerClock(c *clockdog.Clock, ctx *godog.ScenarioContext) {
	ctx.Step(`the time is now`, func() error {
		return isNow(c)
	})

	ctx.Step(`the time is not now`, func() error {
		return isNotNow(c)
	})

	ctx.Step(`the time is "([^"]*)"`, func(s string) error {
		return expectTime(c, s)
	})

	ctx.Step(`wait for ([^\s]*)`, waitFor)
}

func isNow(c *clockdog.Clock) error {
	now := time.Now()
	ts := c.Now()
	min := now.Add(-10 * time.Millisecond)
	max := now.Add(10 * time.Millisecond)

	if min.After(ts) || max.Before(ts) {
		return fmt.Errorf("expected: %q < %q < %q", min.String(), ts.String(), max.String())
	}

	return nil
}

func isNotNow(c *clockdog.Clock) error {
	now := time.Now()
	ts := c.Now()
	min := now.Add(-10 * time.Millisecond)
	max := now.Add(10 * time.Millisecond)

	if ts.After(min) && ts.Before(max) {
		return fmt.Errorf("the time is now")
	}

	return nil
}

func expectTime(c *clockdog.Clock, s string) error {
	expected, err := timeparser.Parse(s)
	if err != nil {
		return err
	}

	now := c.Now()

	if now != expected {
		return fmt.Errorf("expected: %q, got %q", expected.String(), now.String())
	}

	return nil
}

func waitFor(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	<-time.After(d)

	return nil
}
