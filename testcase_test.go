package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/adl32x/sailfoot/sailfoot"
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

	var sf sailfoot.Case
	sf.KnownKeywords = map[string]sailfoot.Keyword{}

	sf.Load("tests/unit.txt")

	assertEqual(t, sf.RootKeyword.Commands[0][0], "click", "")
	assertEqual(t, sf.RootKeyword.Commands[0][1], "div", "")

	assertEqual(t, sf.RootKeyword.Commands[1][0], "navigate", "")

	assertEqual(t, sf.RootKeyword.Commands[2][0], "has_text", "")
	assertEqual(t, sf.RootKeyword.Commands[2][1], "#test .class", "")
	assertEqual(t, sf.RootKeyword.Commands[2][2], "It's, not its!", "")
}

func TestRegex(t *testing.T) {
	var str = "click .element-$$0$$"
	re := regexp.MustCompile("\\$\\$([0-9+])\\$\\$")
	var result = re.FindAllStringSubmatch(str, -1)

	for _, row := range result {
		fmt.Println(row[1])
	}
}
