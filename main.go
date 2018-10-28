package main

import (
	"flag"
	"fmt"

	"github.com/adl32x/sailfoot/driver"
	"github.com/adl32x/sailfoot/sailfoot"
)

func main() {
	startFile := flag.String("file", "start.txt", "Start file")
	driverType := flag.String("driver", "chrome", "Possible values: chrome, firefox, phantomjs, selenium")
	// browser := flag.String("browser", "chrome", "chrome / firefox / phantomjs")
	runner := flag.String("runner", "cli", "Runner type (cli / server)")
	port := flag.Int("port", 3000, "Runner port (cli / server)")
	flag.Parse()

	fmt.Printf("🍤 Sailfoot v0.1. Startfile: %s\n\n", *startFile)

	var sf *sailfoot.Case
	if *driverType == "fake" {
		sf = sailfoot.NewCase(&driver.FakeDriver{})
	} else {
		webdriver := &driver.WebDriver{}
		webdriver.Init(driverType)
		sf = sailfoot.NewCase(webdriver)
	}

	sf.Load(*startFile)

	if *runner == "server" {
		sf.StartServer(*port)
	} else {
		sf.Run()
	}

}
