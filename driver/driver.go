package driver

// TestDriver is and interface that exposes all browser functionality.
// Current implementing classes: FakeDriver, WebDriver (written with Agouti)
type TestDriver interface {
	Start()
	Stop()
	Click(bool, string) bool
	ClickNth(string, int) bool
	ClickClosestTo(string, string) bool
	Navigate(string) bool
	NewPage(string) bool
	GoToNthWindow(int) bool
	HasText(string, string) bool
	Input(bool, string, string) bool
	InputEmpty(string) bool
	Log(string) bool
	Read(string) (string, bool)
	ExecuteJavascript(string) bool
	ClickOnText(string, string) bool
	SendKey(string) bool
	WindowSize(int, int) bool
}

var WEBDRIVER_KEYCODES = map[string]string{
	"NullKey":       string('\ue000'),
	"CancelKey":     string('\ue001'),
	"HelpKey":       string('\ue002'),
	"BackspaceKey":  string('\ue003'),
	"TabKey":        string('\ue004'),
	"ClearKey":      string('\ue005'),
	"ReturnKey":     string('\ue006'),
	"EnterKey":      string('\ue007'),
	"ShiftKey":      string('\ue008'),
	"ControlKey":    string('\ue009'),
	"AltKey":        string('\ue00a'),
	"PauseKey":      string('\ue00b'),
	"EscapeKey":     string('\ue00c'),
	"SpaceKey":      string('\ue00d'),
	"PageUpKey":     string('\ue00e'),
	"PageDownKey":   string('\ue00f'),
	"EndKey":        string('\ue010'),
	"HomeKey":       string('\ue011'),
	"LeftArrowKey":  string('\ue012'),
	"UpArrowKey":    string('\ue013'),
	"RightArrowKey": string('\ue014'),
	"DownArrowKey":  string('\ue015'),
	"InsertKey":     string('\ue016'),
	"DeleteKey":     string('\ue017'),
	"SemicolonKey":  string('\ue018'),
	"EqualsKey":     string('\ue019'),
	"Numpad0Key":    string('\ue01a'),
	"Numpad1Key":    string('\ue01b'),
	"Numpad2Key":    string('\ue01c'),
	"Numpad3Key":    string('\ue01d'),
	"Numpad4Key":    string('\ue01e'),
	"Numpad5Key":    string('\ue01f'),
	"Numpad6Key":    string('\ue020'),
	"Numpad7Key":    string('\ue021'),
	"Numpad8Key":    string('\ue022'),
	"Numpad9Key":    string('\ue023'),
	"MultiplyKey":   string('\ue024'),
	"AddKey":        string('\ue025'),
	"SeparatorKey":  string('\ue026'),
	"SubstractKey":  string('\ue027'),
	"DecimalKey":    string('\ue028'),
	"DivideKey":     string('\ue029'),
	"F1Key":         string('\ue031'),
	"F2Key":         string('\ue032'),
	"F3Key":         string('\ue033'),
	"F4Key":         string('\ue034'),
	"F5Key":         string('\ue035'),
	"F6Key":         string('\ue036'),
	"F7Key":         string('\ue037'),
	"F8Key":         string('\ue038'),
	"F9Key":         string('\ue039'),
	"F10Key":        string('\ue03a'),
	"F11Key":        string('\ue03b'),
	"F12Key":        string('\ue03c'),
	"MetaKey":       string('\ue03d'),
}
