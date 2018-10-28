package main

import (
	"flag"
	"fmt"

	"github.com/adl32x/sailfoot/driver"
	"github.com/adl32x/sailfoot/sailfoot"
)

func main() {
	startFile := flag.String("file", "start.txt", "Start file")
	driverType := flag.String("driver", "default", "(experimental) Driver type")
	runner := flag.String("runner", "cli", "Runner type (cli / server)")
	port := flag.Int("port", 3000, "Runner port (cli / server)")
	// browser := flag.String("browser", "chrome", "chrome / firefox / phantomjs")
	flag.Parse()

	fmt.Printf("🍤 Sailfoot. Startfile: %s\n\n", *startFile)

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
