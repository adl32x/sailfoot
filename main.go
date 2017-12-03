package main

import (
	"flag"
	"github.com/adl32x/sailfoot/testcase"
	"github.com/adl32x/sailfoot/driver"
)

func main() {
	startFile := flag.String("file", "start.txt", "Start file")
	driverType := flag.String("driver", "default", "(experimental) Driver type")
	runner := flag.String("runner", "cli", "Runner type (cli / server)")
	port := flag.Int("runner port", 3000, "Runner port (cli / server)")
	flag.Parse()

	var test *testcase.Testcase
	if *driverType == "fake" {
		test = testcase.NewTestCase(&driver.FakeDriver{})
	} else {
		test = testcase.NewTestCase(&driver.WebDriver{})
	}

	test.Load(*startFile)

	if *runner == "server" {
		test.StartServer(*port)
	} else {
		test.Run()
	}

}
