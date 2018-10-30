package driver

import (
	"os"
	"time"

	log "github.com/adl32x/sailfoot/log"
	"github.com/logrusorgru/aurora"

	"github.com/sclevine/agouti"
)

type WebDriver struct {
	driver *agouti.WebDriver
	page   *agouti.Page
	pages  []*agouti.Page
}

func (w *WebDriver) Start() {
	// w.driver = driver
	w.page, _ = w.driver.NewPage()
	w.pages = append(w.pages, w.page)
	w.page.SetImplicitWait(10000)
}

func (w *WebDriver) Init(driverType *string) {
	if *driverType == "chromeheadless" {
		w.driver = agouti.ChromeDriver(
			agouti.ChromeOptions("args", []string{"--headless", "--disable-gpu"}),
		)
		if err := w.driver.Start(); err != nil {
			log.Error("Failed to start with chromedriver: ", err)
		}
	} else if *driverType == "chrome" {
		w.driver = agouti.ChromeDriver()
		if err := w.driver.Start(); err != nil {
			log.Error("Failed to start chrome: ", err)
		}
	} else if *driverType == "phantomjs" {
		w.driver = agouti.PhantomJS()
		if err := w.driver.Start(); err != nil {
			log.Error("Failed to start phantomjs: ", err)
		}
	} else if *driverType == "firefox" {
		w.driver = agouti.GeckoDriver()
		if err := w.driver.Start(); err != nil {
			log.Error("Failed to start geckodriver: ", err)
		}
	} else {
		log.Error("Unsupported driver %s", *driverType)
		os.Exit(1)
	}

}

func (w *WebDriver) GoToNthWindow(nth int) bool {
	if nth < len(w.pages) {
		w.page = w.pages[nth]
		return true
	}
	return false
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
		log.Logf("click, ´%s´ found multiple elements, clicking first.", arg)
		w.page.First(arg).Click()
	} else {
		el.Click()
	}
	log.Logf("click, ´%s´", arg)
	return true
}

func (w *WebDriver) ClickNth(arg string, nth int) bool {
	w.page.All(arg).At(nth).Click()
	log.Logf("click_%dnth, ´%s´.", nth, arg)
	return true
}

func (w *WebDriver) Navigate(arg string) bool {
	err := w.page.Navigate(arg)
	if err != nil {
		return false
	}
	log.Logf("navigate, ´%s´", arg)
	return true
}

func (w *WebDriver) NewPage(arg string) bool {
	w.page, _ = w.driver.NewPage()
	w.pages = append(w.pages, w.page)
	log.Logf("new_page, ´%s´", arg)
	w.Navigate(arg)
	return true
}

func (w *WebDriver) HasText(arg string, arg2 string) bool {

	if arg2 == "" {
		var number int

		retries := 10
		for {
			w.page.RunScript(JsHasText, map[string]interface{}{"text": arg}, &number)
			if retries == 0 || number == 1 {
				break
			}

			time.Sleep(150 * time.Millisecond)
			retries = retries - 1
		}

		if number == 1 {
			log.Logf("has_text, ´%s´", arg)
		} else {
			log.Errorf("has_text failed. has_text, ´%s´", arg)
		}

		return number == 1
	}

	el := w.page.All(arg)
	count, _ := el.Count()
	if count == 0 {
		log.Errorf("has_text, could not find element ´%s´", arg)
		return false
	}
	if count > 1 {
		log.Error("has_text, too many elements")
		return false
	}

	t, err := el.Text()

	if err != nil {
		log.Error("input, Something went wrong.")
		log.Error(err)
	}

	if t != arg2 {
		log.Errorf("has_text, failed ´%s´ ´%s´", arg, arg2)
		log.Errorf("Text in selection: %s", t)
		return false
	}

	log.Logf("has_text, ´%s´ ´%s´", arg, arg2)
	return true
}

func (w *WebDriver) Input(is_xpath bool, arg string, text string) bool {
	var el *agouti.MultiSelection

	if is_xpath {
		el = w.page.AllByXPath(arg)
	} else {
		el = w.page.All(arg)
	}

	count, _ := el.Count()
	if count == 0 {
		log.Errorf("input, could not find element ´%s´", arg)
		return false
	}
	if count > 1 {
		log.Logf("input, ´%s´ found multiple elements, filling first with ´%s´", arg, text)
		el.First(arg).Fill(text)
	} else {
		el.Fill(text)
	}
	log.Logf("input, ´%s´ ´%s´", arg, text)
	return true
}

func (w *WebDriver) InputEmpty(text string) bool {
	var number int

	retries := 10
	for {
		w.page.RunScript(JsInputEmpty, map[string]interface{}{}, &number)
		if retries == 0 || number == 1 {
			break
		}

		time.Sleep(150 * time.Millisecond)
		retries = retries - 1
	}

	if number == 1 {
		el := w.page.All("input[data-sailfoot-empty]")

		count, _ := el.Count()
		if count != 1 {
			log.Errorf("input. Something went wrong.")
			return false
		}

		el.Fill(text)
		w.page.RunScript(JsInputEmptyReset, map[string]interface{}{}, &number)

		log.Logf("input, ´%s´", text)
	} else {
		log.Errorf("input, ´%s´", text)
	}

	return number == 1
}

func (w *WebDriver) Log(arg string) bool {
	log.Logf("%s %s", aurora.Bold("Log:"), arg)
	return true
}

func (w *WebDriver) Read(arg string) (string, bool) {
	el := w.page.All(arg)

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
		log.Logf("read, ´%s´ found multiple elements, returning first value ´%s´", arg, t)
		return t, true

	} else {
		t, _ := el.Text()
		els, _ := el.Elements()
		tag, _ := els[0].GetName()
		if tag == "input" {
			t, _ = el.Attribute("value")
		}
		log.Logf("read, ´%s´ got value ´%s´", arg, t)
		return t, true
	}
}

func (w *WebDriver) ExecuteJavascript(jsString string) bool {
	w.page.RunScript(jsString, map[string]interface{}{}, struct{}{})
	return true
}

func (w *WebDriver) ClickOnText(selector string, text string) bool {

	var number int

	retries := 10
	for {
		w.page.RunScript(JsClickWithText, map[string]interface{}{"text": text}, &number)
		if retries == 0 || number == 1 {
			break
		}

		time.Sleep(150 * time.Millisecond)
		retries = retries - 1
	}

	if number == 1 {
		log.Logf("click_on_text, ´%s´", text)
	} else {
		log.Errorf("click_on_text, ´%s´", text)
	}

	return number == 1
}

func (w *WebDriver) ClickClosestTo(selector1 string, selector2 string) bool {

	var number int

	retries := 10
	for {
		w.page.RunScript(JsClickClosest, map[string]interface{}{"text": selector1, "text2": selector2}, &number)
		if retries == 0 || number == 1 {
			break
		}

		time.Sleep(150 * time.Millisecond)
		retries = retries - 1
	}

	if number == 1 {
		log.Logf("click_closest_to, ´%s´ ´%s´", selector1, selector2)
	} else {
		log.Errorf("click_closest_to, ´%s´ ´%s´", selector1, selector2)
	}

	return number == 1
}

func (w *WebDriver) SendKey(keycode string) bool {
	var number int
	w.page.RunScript(JsActiveElement, map[string]interface{}{}, &number)

	keyToSend := WEBDRIVER_KEYCODES[keycode]
	if keyToSend == "" {
		keyToSend = keycode
	}

	err := w.page.Find("[data-sailfoot-active-element]").SendKeys(keyToSend)

	w.page.RunScript(JsActiveElementReset, map[string]interface{}{}, &number)

	if err != nil {
		log.Errorf("sendkey, something went wrong")
		log.Error(err)
		return false
	}

	log.Logf("sendkey, ´%s´", keycode)
	return true
}
