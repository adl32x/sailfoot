package driver

import (
	log "github.com/sirupsen/logrus"
	"github.com/sclevine/agouti"
)


type TestDriver interface {
	Start()
	Stop()
	Click(string) bool
	Navigate(string) bool
	HasText(string, string) bool
	Input(string, string) bool
}

type WebDriver struct {
	driver *agouti.WebDriver
	page *agouti.Page
}

type FakeDriver struct {
}

func (w *WebDriver) Start() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatal("Failed to start Selenium: ", err)
	}
	w.driver = driver
	w.page, _ = w.driver.NewPage(agouti.Browser("chrome"))
}

func (w *WebDriver) Stop() {
	w.driver.Stop()
}

func (w *WebDriver) Click(arg string) bool {
	el := w.page.Find(arg)
	count, _ := el.Count()
	if count == 0 {
		log.Errorf("click, could not find element ´%s´", arg)
		return false
	}
	if count > 1 {
		log.Infof("click, ´%s´ found multiple elements, clicking first.", arg)
		el.First(arg).Click()
	} else {
		el.Click()
	}
	log.Infof("click, ´%s´", arg)
	return true
}

func (w *WebDriver) Navigate(arg string) bool {
	err := w.page.Navigate(arg)
	if err != nil {
		return false
	}
	log.Info("navigate, ´%s´", arg)
	return true
}

func (w *WebDriver) HasText(arg string, text string) bool {
	el := w.page.Find(arg)
	count, _ := el.Count()
	if count == 0 {
		log.Errorf("has_text, could not find element ´%s´", arg)
		return false
	}
	if count > 1 {
		return false
	} else {
		t, _ := el.Text()

		if t != text {
			log.Errorf("has_text, failed ´%s´ ´%s´", arg, text)
			return false
		}
	}
	log.Infof("has_text, ´%s´ ´%s´", arg, text)
	return true
}

func (w *WebDriver) Input(arg string, text string) bool {
	el := w.page.Find(arg)
	count, _ := el.Count()
	if count == 0 {
		log.Errorf("input, could not find element ´%s´", arg)
		return false
	}
	if count > 1 {
		log.Infof("input, ´%s´ found multiple elements, filling first with ´%s´", arg, text)
		el.First(arg).Fill(text)
	} else {
		el.Fill(text)
	}
	log.Infof("input, ´%s´ ´%s´", arg, text)
	return true
}

// FakeDriver:

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