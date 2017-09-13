package testcase

import (
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"regexp"
	"github.com/adl32x/sailfoot/driver"
	"time"
	"strconv"
)

type Testcase struct {
	Driver   driver.TestDriver
	Commands [][]string
	LabelLocation map[string]int
	Variables map[string]string
}

func NewTestCase(d driver.TestDriver) *Testcase {
	t := &Testcase{}
	t.Driver = d
	t.LabelLocation = map[string]int{}
	t.Variables = map[string]string{}
	return t
}

func (t *Testcase) Run() {
	t.Driver.Start()
	for rowNumber := 0; rowNumber < len(t.Commands); rowNumber++ {
		command := t.Commands[rowNumber]

		skip_sleep := false
		result := true

		if command[0] == "click" {
			result = t.Driver.Click(command[1])
		} else if command[0] == "navigate" {
			result = t.Driver.Navigate(command[1])
		} else if command[0] == "has_text" {
			result = t.Driver.HasText(command[1], command[2])
		} else if command[0] == "input" {
			result = t.Driver.Input(command[1], command[2])
		} else if command[0] == "sleep" {
			sleep_time, _ := strconv.Atoi(command[1])
			time.Sleep(time.Duration(sleep_time) * time.Millisecond)
		} else if command[0] == "log" {
			result = t.Driver.Log(command[1])
		} else if command[0] == "label" {
			log.Infof("label, ´%s´", command[1])
			t.LabelLocation[command[1]] = rowNumber
			skip_sleep = true
		} else if command[0] == "jump" {
			log.Infof("jump, ´%s´", command[1])
			rowNumber = t.LabelLocation[command[1]] -1
			skip_sleep = true
		} else if command[0] == "read" {
			var value string
			value, result = t.Driver.Read(command[1])
			t.Variables[command[2]] = value
		}

		if result == false {
			t.Driver.Stop()
			os.Exit(1)
			return
		}

		if skip_sleep {
			time.Sleep(150 * time.Millisecond)
		}

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
		} else if command[0] == "sleep" {
			// TODO. Add parsing checks.
		} else if command[0] == "input" {
			// TODO. Add parsing checks.
		} else if command[0] == "log" {
			// TODO. Add parsing checks.
		} else if command[0] == "read" {
			// TODO. Add parsing checks.
		} else {
			log.Fatal("Unknown command ", command)
		}

		t.Commands = append(t.Commands, command)

	}
}
