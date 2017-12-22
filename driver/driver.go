package driver

type TestDriver interface {
	Start()
	Stop()
	Click(bool, string) bool
	Navigate(string) bool
	NewPage(string) bool
	GoToNthWindow(int) bool
	HasText(string, string) bool
	Input(bool, string, string) bool
	Log(string) bool
	Read(string) (string, bool)
	ExecuteJavascript(string) bool
	ClickOnText(string, string) bool
}
