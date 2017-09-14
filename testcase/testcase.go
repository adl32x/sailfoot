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
	"fmt"
	"path/filepath"
)

type Testcase struct {
	Driver   driver.TestDriver
	Command  Command
	KnownCommands map[string]Command
}

type Command struct {
	Commands [][]string
	LabelLocation map[string]int
	Variables map[string]string
}

func (c *Command) Init() {
	c.LabelLocation = map[string]int{}
	c.Variables = map[string]string{}
}

func NewTestCase(d driver.TestDriver) *Testcase {
	t := &Testcase{}
	t.Driver = d
	t.KnownCommands = map[string]Command{}
	return t
}

func (c *Command) Run(driver driver.TestDriver, knownCommands map[string]Command) {
	for rowNumber := 0; rowNumber < len(c.Commands); rowNumber++ {
		command := c.Commands[rowNumber]

		skip_sleep := false
		result := true

		if command[0] == "click" {
			result = driver.Click(command[1])
		} else if command[0] == "navigate" {
			result = driver.Navigate(command[1])
		} else if command[0] == "has_text" {
			result = driver.HasText(command[1], command[2])
		} else if command[0] == "input" {
			result = driver.Input(command[1], command[2])
		} else if command[0] == "sleep" {
			sleep_time, _ := strconv.Atoi(command[1])
			time.Sleep(time.Duration(sleep_time) * time.Millisecond)
		} else if command[0] == "log" {
			result = driver.Log(command[1])
		} else if command[0] == "label" {
			log.Infof("label, ´%s´", command[1])
			c.LabelLocation[command[1]] = rowNumber
			skip_sleep = true
		} else if command[0] == "jump" {
			log.Infof("jump, ´%s´", command[1])
			rowNumber = c.LabelLocation[command[1]] -1
			skip_sleep = true
		} else if command[0] == "read" {
			var value string
			value, result = driver.Read(command[1])
			c.Variables[command[2]] = value
		} else {
			keyword := knownCommands[command[0]]
			// TODO check if the command exists
			keyword.Run(driver, knownCommands)
		}

		if result == false {
			driver.Stop()
			os.Exit(1)
			return
		}

		if skip_sleep {
			time.Sleep(150 * time.Millisecond)
		}

	}
}


func (t *Testcase) Run() {
	t.Driver.Start()
	t.Command.Run(t.Driver, t.KnownCommands)
	t.Driver.Stop()
}

func (t *Testcase) Load(filename string) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Error("Err: ", err)
		os.Exit(1)
	}

	var location string
	if strings.Contains(filename, "/") {
		fileNameArray := strings.Split(filename, "/")
		fileNameArray = fileNameArray[0:len(fileNameArray)-1]
		location = strings.Join(fileNameArray, "/")
		location = location + "/keywords/"
	} else {
		location = "keywords/"
	}

	filepath.Walk(location, func(path string, _ os.FileInfo, _ error) error {
		if strings.Contains(path, ".txt") {
			fmt.Println(path)
			keyword := strings.Replace(filepath.Base(path), ".txt", "", -1)
			file, _ := ioutil.ReadFile(path)
			t.KnownCommands[keyword] = fileToCommands(file)
		}

		return nil
	})

	// os.Exit(0)

	t.Command = fileToCommands(file)
}

func fileToCommands(file []byte) Command {
	str := string(file)
	log.Debug("File content: ", str)
	rows := strings.Split(str, "\n")
	c := Command{}
	c.Init()

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
			if strings.HasPrefix(*c, "'") && strings.HasSuffix(*c, "'") {
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
		}

		c.Commands = append(c.Commands, command)

	}

	return c
}
