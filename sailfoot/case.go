package sailfoot

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/adl32x/sailfoot/driver"
	"github.com/adl32x/sailfoot/utils"
	log "github.com/sirupsen/logrus"
)

type Case struct {
	Driver        driver.TestDriver
	RunList       RunList
	KnownCommands map[string]RunList
}

type RunList struct {
	Commands      [][]string
	LabelLocation map[string]int
	Variables     map[string]string
	IsATest       bool
	Passed        bool
	TestCaseName  string
	LastResult    bool
}

func (c *RunList) Init() {
	c.LabelLocation = map[string]int{}
	c.Variables = map[string]string{}
	c.IsATest = false
	c.Passed = true
	c.TestCaseName = ""
	c.LastResult = false
}

func NewCase(d driver.TestDriver) *Case {
	t := &Case{}
	t.Driver = d
	t.KnownCommands = map[string]RunList{}
	return t
}

func (c *RunList) Run(driver driver.TestDriver, knownCommands map[string]RunList, args []string) {
	for rowNumber := 0; rowNumber < len(c.Commands); rowNumber++ {

		command := c.Commands[rowNumber]

		commandTmp := make([]string, len(command))
		copy(commandTmp, command)

		skip_sleep := false
		result := true
		skip_fail := false

		for i := range command {
			c := &command[i]

			if i == 0 && strings.HasPrefix(*c, "!") {
				skip_fail = true
				commandTmp[i] = strings.Trim(*c, "!")
			}

			//
			re := regexp.MustCompile("\\$\\$([0-9+])\\$\\$")
			var templateArgs = re.FindAllStringSubmatch(*c, -1)

			if templateArgs != nil {
				for _, row := range templateArgs {
					var argn, _ = strconv.Atoi(row[1])

					commandTmp[i] = strings.Replace(*c, "$$"+row[1]+"$$", args[argn], -1)
					// TODO: Maybe check if args[argn] exists. Also make it possible to escape $$.
				}
			}

		}

		if commandTmp[0] == "click" {
			result = driver.Click(false, commandTmp[1])
		} else if commandTmp[0] == "click_x" {
			result = driver.Click(true, commandTmp[1])
		} else if commandTmp[0] == "click_on_text" {
			result = driver.ClickOnText("", commandTmp[1])
		} else if commandTmp[0] == "navigate" {
			result = driver.Navigate(commandTmp[1])
		} else if commandTmp[0] == "has_text" {
			result = driver.HasText(commandTmp[1], commandTmp[2])
		} else if commandTmp[0] == "input" {
			result = driver.Input(false, commandTmp[1], commandTmp[2])
		} else if commandTmp[0] == "input_x" {
			result = driver.Input(true, commandTmp[1], commandTmp[2])
		} else if commandTmp[0] == "sleep" {
			sleep_time, _ := strconv.Atoi(commandTmp[1])
			time.Sleep(time.Duration(sleep_time) * time.Millisecond)
		} else if commandTmp[0] == "log" {
			result = driver.Log(commandTmp[1])
		} else if commandTmp[0] == "label" {
			log.Infof("label, ´%s´", commandTmp[1])
			c.LabelLocation[commandTmp[1]] = rowNumber
			skip_sleep = true
		} else if commandTmp[0] == "jump" {
			log.Infof("jump, ´%s´", commandTmp[1])
			rowNumber = c.LabelLocation[commandTmp[1]] - 1
			skip_sleep = true
		} else if commandTmp[0] == "read" {
			var value string
			value, result = driver.Read(commandTmp[1])
			c.Variables[commandTmp[2]] = value
		} else if commandTmp[0] == "testcase" {
			skip_sleep = true
		} else if commandTmp[0] == "stop_if_success" {
			if c.LastResult == true {
				return
			}
		} else if commandTmp[0] == "execute" {
			skip_sleep = true
			out, err := utils.Execute(commandTmp[1])
			if err != nil {
				log.Fatalf("RunList %s failed, %s", commandTmp[1], err)
			}
			log.Printf("execute, output: %s", out)
		} else {
			keyword := knownCommands[commandTmp[0]]
			// TODO check if the command exists
			keyword.Run(driver, knownCommands, commandTmp[1:])
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

func (t *Case) Run() {
	t.Driver.Start()
	t.RunList.Run(t.Driver, t.KnownCommands, nil)
	t.Driver.Stop()

	for i := range t.KnownCommands {
		command := t.KnownCommands[i]
		if command.IsATest && command.Passed {
			fmt.Printf("%s - Passed\n", command.TestCaseName)
		} else if command.IsATest && !command.Passed {
			fmt.Printf("%s - Failed\n", command.TestCaseName)
		}
	}
}

func (t *Case) StartServer(port int) {
	t.Driver.Start()

	t.Listen(port)
	// t.RunList.Run(t.Driver, t.KnownCommands, nil)

	// t.Driver.Stop()

}

func (t *Case) Load(filename string) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Error("Err: ", err)
		os.Exit(1)
	}

	var location string
	if strings.Contains(filename, "/") {
		fileNameArray := strings.Split(filename, "/")
		fileNameArray = fileNameArray[0 : len(fileNameArray)-1]
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

	t.RunList = fileToCommands(file)
}

func fileToCommands(file []byte) RunList {
	str := string(file)
	log.Debug("File content: ", str)
	rows := strings.Split(str, "\n")
	c := RunList{}
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
