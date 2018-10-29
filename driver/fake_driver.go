package driver

type FakeDriver struct {
}

func (f *FakeDriver) Start() {}

func (f *FakeDriver) Stop() {}

func (f *FakeDriver) Click(is_xpath bool, arg string) bool {
	return true
}

func (f *FakeDriver) Navigate(arg string) bool {
	return true
}

func (f *FakeDriver) NewPage(arg string) bool {
	return true
}

func (f *FakeDriver) GoToNthWindow(arg int) bool {
	return true
}

func (f *FakeDriver) HasText(arg string, arg2 string) bool {
	return true
}

func (f *FakeDriver) Input(is_xpath bool, arg string, text string) bool {
	return true
}

func (f *FakeDriver) Log(arg string) bool {
	return true
}

func (f *FakeDriver) Read(arg string) (string, bool) {
	return "", true
}

func (w *FakeDriver) ExecuteJavascript(jsString string) bool {
	return true
}

func (w *FakeDriver) ClickOnText(selector string, text string) bool {
	return true
}

func (w *FakeDriver) ClickClosestTo(selector1 string, selector2 string) bool {
	return true
}

func (w *FakeDriver) InputEmpty(text string) bool {
	return true
}

func (w *FakeDriver) SendKey(keycode string) bool {
	return true
}
