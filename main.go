package main

import (
	"flag"
	"github.com/adl32x/sailfoot/testcase"
	"github.com/adl32x/sailfoot/driver"
)

func main() {
	startFile := flag.String("file", "start.txt", "Start file")
	driverType := flag.String("driver", "default", "(experimental) Driver type")
	flag.Parse()

	var test *testcase.Testcase
	if *driverType == "fake" {
		test = testcase.NewTestCase(&driver.FakeDriver{})
	} else {
		test = testcase.NewTestCase(&driver.WebDriver{})
	}

	test.Load(*startFile)
	test.Run()
}
