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
	assertEqual(t, sf.RootKeyword.Commands[3][1], "mongo TESTDBE2E --eval 'db.dropDatabase()'", "")
	assertEqual(t, sf.RootKeyword.Commands[4][1], "foo eval 'bar' ", "")
	assertEqual(t, sf.RootKeyword.Commands[5][1], "foo eval 'bar' baz", "")
}

func TestRegex(t *testing.T) {
	var str = "click .element-$$0$$"
	re := regexp.MustCompile("\\$\\$([0-9+])\\$\\$")
	var result = re.FindAllStringSubmatch(str, -1)

	for _, row := range result {
		fmt.Println(row[1])
	}
}

func TestLineSplit(t *testing.T) {
	r1 := sailfoot.SplitLine("hello world")
	r2 := sailfoot.SplitLine("hello 'world' 'foo bar'")
	r3 := sailfoot.SplitLine("hello 'it\\'s' 'a new world'")
	r4 := sailfoot.SplitLine("execute 'mongo TESTDBE2E --eval \\'db.dropDatabase()\\''")
	r5 := sailfoot.SplitLine("execute 'mongo TESTDBE2E --eval \\'db.dropDatabase()\\' '")

	// fmt.Println(strings.Join(r1, ","))
	// fmt.Println(strings.Join(r2, ","))
	// fmt.Println(strings.Join(r3, ","))
	// fmt.Println(strings.Join(r4, ","))
	// fmt.Println(strings.Join(r5, ","))

	assertEqual(t, r1[0], "hello", "")
	assertEqual(t, r1[1], "world", "")

	assertEqual(t, r2[0], "hello", "")
	assertEqual(t, r2[1], "world", "")
	assertEqual(t, r2[2], "foo bar", "")

	assertEqual(t, r3[1], "it's", "")
	assertEqual(t, r4[1], "mongo TESTDBE2E --eval 'db.dropDatabase()'", "")
	assertEqual(t, r5[1], "mongo TESTDBE2E --eval 'db.dropDatabase()' ", "")
}
