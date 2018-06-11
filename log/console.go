package log

import (
	"fmt"

	aurora "github.com/logrusorgru/aurora"
)

var colors = true

func Log(str string) {
	fmt.Println("%s "+str, aurora.Green("."))
}

func Error(args ...interface{}) {
	for i := 0; i < len(args); i++ {
		fmt.Print(aurora.Red(args[i]))
	}
	fmt.Print("\n")

}
