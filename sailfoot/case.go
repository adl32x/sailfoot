package sailfoot

import (
	"fmt"
	"os"
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
	RootKeyword   Keyword
	KnownKeywords map[string]Keyword
}

type Keyword struct {
	Commands      [][]string
	LabelLocation map[string]int
	Variables     map[string]string
	IsATest       bool
	Passed        bool
	TestCaseName  string
	LastResult    bool
}

func (k *Keyword) Init() {
	k.LabelLocation = map[string]int{}
	k.Variables = map[string]string{}
	k.IsATest = false
	k.Passed = true
	k.TestCaseName = ""
	k.LastResult = false
}

func NewCase(d driver.TestDriver) *Case {
	t := &Case{}
	t.Driver = d
	t.KnownKeywords = map[string]Keyword{}
	return t
}

func (k *Keyword) Run(driver driver.TestDriver, knownCommands map[string]Keyword, args []string) {
	for rowNumber := 0; rowNumber < len(k.Commands); rowNumber++ {

		command := k.Commands[rowNumber]

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
			k.LabelLocation[commandTmp[1]] = rowNumber
			skip_sleep = true
		} else if commandTmp[0] == "jump" {
			log.Infof("jump, ´%s´", commandTmp[1])
			rowNumber = k.LabelLocation[commandTmp[1]] - 1
			skip_sleep = true
		} else if commandTmp[0] == "read" {
			var value string
			value, result = driver.Read(commandTmp[1])
			k.Variables[commandTmp[2]] = value
		} else if commandTmp[0] == "testcase" {
			skip_sleep = true
		} else if commandTmp[0] == "stop_if_success" {
			if k.LastResult == true {
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

		k.LastResult = result

	}

	if k.IsATest {
		k.Passed = true
	}
}

func (c *Case) Run() {
	c.Driver.Start()
	c.RootKeyword.Run(c.Driver, c.KnownKeywords, nil)
	c.Driver.Stop()

	for i := range c.KnownKeywords {
		command := c.KnownKeywords[i]
		if command.IsATest && command.Passed {
			fmt.Printf("%s - Passed\n", command.TestCaseName)
		} else if command.IsATest && !command.Passed {
			fmt.Printf("%s - Failed\n", command.TestCaseName)
		}
	}
}

func (c *Case) StartServer(port int) {
	c.Driver.Start()

	c.Listen(port)
	// t.RunList.Run(t.Driver, t.KnownCommands, nil)

	// t.Driver.Stop()

}
