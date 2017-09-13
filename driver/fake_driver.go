package driver

type FakeDriver struct {
}


func (f *FakeDriver) Start() {}

func (f *FakeDriver) Stop() {}

func (f *FakeDriver) Click(arg string) bool {
	return true
}

func (f *FakeDriver) Navigate(arg string) bool {
	return true
}

func (f *FakeDriver) HasText(arg string, text string) bool {
	return true
}

func (f *FakeDriver) Input(arg string, text string) bool {
	return true
}

func (f *FakeDriver) Log(arg string) bool {
	return true
}

func (f *FakeDriver) Read(arg string) (string, bool) {
	return "", true
}