# Cucumber Clock steps for Golang

[![Build Status](https://github.com/nhatthm/clockdog/actions/workflows/test.yaml/badge.svg)](https://github.com/nhatthm/clockdog/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/nhatthm/clockdog/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/clockdog)
[![Go Report Card](https://goreportcard.com/badge/github.com/nhatthm/httpmock)](https://goreportcard.com/report/github.com/nhatthm/httpmock)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/github.com/nhatthm/clockdog)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

`clockdog` uses [`nhatthm/go-clock`](https://github.com/nhatthm/go-clock) to provide steps for `cucumber/godog` and
makes it easy to run tests with `time`.

## Usage

Initiate the clock and register it to the scenario.

```go
package mypackage

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/nhatthm/clockdog"
)

func TestIntegration(t *testing.T) {
	clock := clockdog.New()
	suite := godog.TestSuite{
		Name:                 "Integration",
		TestSuiteInitializer: nil,
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			clock.RegisterContext(ctx)
		},
		Options: &godog.Options{
			Strict:    true,
			Output:    out,
			Randomize: rand.Int63(),
		},
	}
	
	// Inject the clock to your application then run the suite.
	status := suite.Run()
}
```

Read more about [`nhatthm/go-clock`](https://github.com/nhatthm/go-clock)

## Donation

If this project help you reduce time to develop, you can give me a cup of coffee :)

### Paypal donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;or scan this

<img src="https://user-images.githubusercontent.com/1154587/113494222-ad8cb200-94e6-11eb-9ef3-eb883ada222a.png" width="147px" />
