package log

import (
	"fmt"

	"github.com/fatih/color"
)

var colors = true
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func Log(str string) {
	fmt.Println("%s "+str, green("."))
}

func Debug(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		fmt.Print((args[i]))
	}
	fmt.Print("\n")
}

func Error(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		fmt.Print(red(args[i]))
	}
	fmt.Print("\n")
}
