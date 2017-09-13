package driver

type TestDriver interface {
	Start()
	Stop()
	Click(string) bool
	Navigate(string) bool
	HasText(string, string) bool
	Input(string, string) bool
	Log(string) bool
	Read(string) (string, bool)
}


