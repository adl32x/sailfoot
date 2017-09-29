package main

import (
	"testing"
	"fmt"
	"github.com/adl32x/sailfoot/testcase"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestLoad(t *testing.T) {

	var test testcase.Testcase

	test.Load("tests/unit.txt")

	assertEqual(t, test.Commands[0][0], "click", "")
	assertEqual(t, test.Commands[0][1], "div", "")

	assertEqual(t, test.Commands[1][0], "navigate", "")

	assertEqual(t, test.Commands[2][0], "has_text", "")
	assertEqual(t, test.Commands[2][1], "#test .class", "")
	assertEqual(t, test.Commands[2][2], "It's, not its!", "")
}


