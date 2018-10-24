package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Log("Test")
	Logf("Test %s %s", "two", "huh")
	LogOk("Test")
	LogFail("Test")
	Error("Err.", "Reason.")
}
