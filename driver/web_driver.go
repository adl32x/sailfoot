package driver

import (
	log "github.com/sirupsen/logrus"
	"github.com/sclevine/agouti"
)

type WebDriver struct {
	driver *agouti.WebDriver
	page *agouti.Page
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

func (w *WebDriver) Click(is_xpath bool, arg string) bool {
	var el *agouti.MultiSelection

	if is_xpath {
		el = w.page.AllByXPath(arg)
	} else {
		el = w.page.All(arg)
	}

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
	log.Infof("navigate, ´%s´", arg)
	return true
}

func (w *WebDriver) HasText(arg string, text string) bool {
	el := w.page.All(arg)
	count, err := el.Count()
	if count == 0 {
		log.Errorf("has_text, could not find element ´%s´", arg)
		log.Error(err)
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

func (w *WebDriver) Log(arg string) bool {
	log.Println(arg)
	return true
}

func (w *WebDriver) Read(arg string) (string, bool) {
	el := w.page.Find(arg)

	count, _ := el.Count()

	if count == 0 {
		log.Errorf("read, could not find element ´%s´", arg)
		return "", false
	}
	if count > 1 {
		t, _ := el.First(arg).Text()
		els, _ := el.First(arg).Elements()
		tag, _ := els[0].GetName()
		if tag == "input" {
			t, _ = el.Attribute("value")
		}
		log.Infof("read, ´%s´ found multiple elements, returning first value ´%s´", arg, t)
		return t, true

	} else {
		t, _ := el.Text()
		els, _ := el.Elements()
		tag, _ := els[0].GetName()
		if tag == "input" {
			t, _ = el.Attribute("value")
		}
		log.Infof("read, ´%s´ got value ´%s´", arg, t)
		return t, true
	}
}