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
	"github.com/adl32x/sailfoot/utils"
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
	IsATest bool
	Passed bool
	TestCaseName string
	LastResult bool
}

func (c *Command) Init() {
	c.LabelLocation = map[string]int{}
	c.Variables = map[string]string{}
	c.IsATest = false
	c.Passed = true
	c.TestCaseName = ""
	c.LastResult = false
}

func NewTestCase(d driver.TestDriver) *Testcase {
	t := &Testcase{}
	t.Driver = d
	t.KnownCommands = map[string]Command{}
	return t
}

func (c *Command) Run(driver driver.TestDriver, knownCommands map[string]Command, args []string) {
	for rowNumber := 0; rowNumber < len(c.Commands); rowNumber++ {

		command := c.Commands[rowNumber]

		skip_sleep := false
		result := true
		skip_fail := false

		for i := range command {
			c := &command[i]

			if i == 0 && strings.HasPrefix(*c,"!") {
				skip_fail = true
				command[i] = strings.Trim(*c, "!")
			}

			//
			re := regexp.MustCompile("\\$\\$([0-9+])\\$\\$")
			var templateArgs = re.FindAllStringSubmatch(*c, -1)

			if templateArgs != nil {
				for _, row := range templateArgs {
					var argn, _ = strconv.Atoi(row[1])

					command[i] = strings.Replace(*c, "$$"+row[1]+"$$", args[argn], -1)
					// TODO: Maybe check if args[argn] exists. Also make it possible to escape $$.
				}
			}

		}

		if command[0] == "click" {
			result = driver.Click(false, command[1])
		} else if command[0] == "click_x" {
			result = driver.Click(true, command[1])
		} else if command[0] == "navigate" {
			result = driver.Navigate(command[1])
		} else if command[0] == "has_text" {
			result = driver.HasText(command[1], command[2])
		} else if command[0] == "input" {
			result = driver.Input(false, command[1], command[2])
		} else if command[0] == "input_x" {
			result = driver.Input(true, command[1], command[2])
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
		} else if command[0] == "testcase" {
			skip_sleep = true
		} else if command[0] == "stop_if_success" {
			if c.LastResult == true {
				return
			}
		} else if command[0] == "execute" {
			skip_sleep = true
			out, err := utils.Execute(command[1])
			if err != nil {
				log.Fatalf("Command %s failed, %s", command[1], err)
			}
			log.Printf("execute, output: %s", out)
		} else {
			keyword := knownCommands[command[0]]
			// TODO check if the command exists
			keyword.Run(driver, knownCommands, command[1:])
		}

		if result == false && skip_fail == false {
			driver.Stop()
			os.Exit(1)
			return
		}

		if skip_sleep {
			time.Sleep(150 * time.Millisecond)
		}

		c.LastResult = result

	}

	if c.IsATest {
		c.Passed = true
	}
}


func (t *Testcase) Run() {
	t.Driver.Start()
	t.Command.Run(t.Driver, t.KnownCommands, nil)
	t.Driver.Stop()

	for i, _ := range t.KnownCommands {
		command := t.KnownCommands[i]
		if command.IsATest && command.Passed{
			fmt.Printf("%s - Passed\n", command.TestCaseName)
		} else if command.IsATest && !command.Passed {
			fmt.Printf("%s - Failed\n", command.TestCaseName)
		}
	}
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
		} else if command[0] == "execute" {
			// TODO. Add parsing checks.
		} else if command[0] == "testcase" {
			c.IsATest = true
			c.TestCaseName = command[1]
		}

		c.Commands = append(c.Commands, command)

	}

	return c
}
