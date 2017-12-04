package main

import (
	"flag"

	"github.com/adl32x/sailfoot/driver"
	"github.com/adl32x/sailfoot/sailfoot"
)

func main() {
	startFile := flag.String("file", "start.txt", "Start file")
	driverType := flag.String("driver", "default", "(experimental) Driver type")
	runner := flag.String("runner", "cli", "Runner type (cli / server)")
	port := flag.Int("runner port", 3000, "Runner port (cli / server)")
	flag.Parse()

	var sf *sailfoot.Case
	if *driverType == "fake" {
		sf = sailfoot.NewCase(&driver.FakeDriver{})
	} else {
		sf = sailfoot.NewCase(&driver.WebDriver{})
	}

	sf.Load(*startFile)

	if *runner == "server" {
		sf.StartServer(*port)
	} else {
		sf.Run()
	}

}
