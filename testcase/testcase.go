package testcase

import (
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"regexp"
	"aloha-r/driver"
	"time"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
}

type Testcase struct {
	Driver   driver.TestDriver
	Commands [][]string
}

func NewTestCase(d driver.TestDriver) *Testcase {
	t := &Testcase{}
	t.Driver = d
	return t
}

func (t *Testcase) Run() {
	t.Driver.Start()
	for _, command := range t.Commands {
		result := true
		if command[0] == "click" {
			result = t.Driver.Click(command[1])
		} else if command[0] == "navigate" {
			result = t.Driver.Navigate(command[1])
		} else if command[0] == "has_text" {
			result = t.Driver.HasText(command[1], command[2])
		} else if command[0] == "input" {
			result = t.Driver.Input(command[1], command[2])
		}
		if result == false {
			t.Driver.Stop()
			os.Exit(1)
			return
		}

		time.Sleep(150 * time.Millisecond)
	}
	t.Driver.Stop()
}

func (t *Testcase) Load(filename string) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Error("Err: ", err)
		os.Exit(1)
	}

	str := string(file)
	log.Debug("File content: ", str)
	rows := strings.Split(str, "\n")

	for _, row := range rows {
		row = strings.Trim(row, " \t")
		if len(row) == 0 {
			continue
		}

		if string(row[0]) == "#" {
			continue
		}

		// ('([^']|\\')*'|[\S]+)+
		//re, _ := regexp.Compile(`'(?:[^\\']|\\.)*'`)
		re, _ := regexp.Compile(`('(\\'|[^'])*'|[\S]+)+`)

		command := re.FindAllString(row, -1)

		for i := range command {
			c := &command[i]
			if strings.HasPrefix(*c,"'") && strings.HasSuffix(*c,"'") {
				command[i] = strings.Trim(*c, "'")
				command[i] = strings.Replace(*c, "\\'", "'", -1)
			}
		}

		if command[0] == "click" {
			// TODO. Add parsing checks.
		} else if command[0] == "navigate" {
			// TODO. Add parsing checks.
		} else if command[0] == "has_text" {
			// TODO. Add parsing checks.
		} else if command[0] == "jump" {
			// TODO. Add parsing checks.
		} else if command[0] == "label" {
			// TODO. Add parsing checks.
		} else if command[0] == "input" {
			// TODO. Add parsing checks.
		} else {
			log.Fatal("Unknown command ", command)
		}

		t.Commands = append(t.Commands, command)

	}
}