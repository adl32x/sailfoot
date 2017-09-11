package main

import (
	"flag"
	//log "github.com/sirupsen/logrus"
	"aloha-r/testcase"
	"aloha-r/driver"
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
