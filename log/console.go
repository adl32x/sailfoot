package log

import (
	"fmt"
	"os"

	aurora "github.com/logrusorgru/aurora"
)

var colors = true

func Log(args ...interface{}) {
	fmt.Printf("%s ", aurora.Green("-"))
	fmt.Println(args...)
}

func Logf(format string, args ...interface{}) {
	fmt.Printf("%s ", aurora.Green("-"))
	fmt.Printf(format, args...)
	fmt.Printf("\n")
}

func Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func LogOk(str string) {
	fmt.Printf("%s %s\n", aurora.Green("✓"), str)
}

func LogFail(str string) {
	fmt.Printf("%s %s\n", aurora.Red("✗"), aurora.Red(str))
}

func Error(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		fmt.Print(aurora.Red(args[i]))
	}
	fmt.Print("\n")
}

func Errorf(format string, args ...interface{}) {
	fmt.Print(aurora.Red("✗ "))
	fmt.Printf(format, args...)
	fmt.Print("\n")
}

func Println(args ...interface{}) {
	fmt.Println(args...)
}

func Debug(args ...interface{}) {
	if os.Getenv("DEBUG") != "true" {
		return
	}
	for i := 0; i < len(args); i++ {
		fmt.Print(aurora.Cyan(args[i]))
	}
	fmt.Print("\n")
}
