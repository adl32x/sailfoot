package sailfoot

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"

	"github.com/adl32x/sailfoot/driver"
	"github.com/adl32x/sailfoot/log"
	"github.com/adl32x/sailfoot/utils"
)

type Case struct {
	Driver        driver.TestDriver
	RootKeyword   *Keyword
	KnownKeywords map[string]*Keyword
}

type Keyword struct {
	Commands      [][]string
	LabelLocation map[string]int
	Variables     map[string]string
	IsATest       bool
	Passed        bool
	TestCaseName  string
	LastResult    bool
	Ran           bool
	SkipFail      bool
}

func (k *Keyword) Init() {
	k.LabelLocation = map[string]int{}
	k.Variables = map[string]string{}
	k.IsATest = false
	k.Passed = true
	k.TestCaseName = ""
	k.LastResult = false
	k.Ran = false
	k.SkipFail = false
}

func NewCase(d driver.TestDriver) *Case {
	t := &Case{}
	t.Driver = d
	t.KnownKeywords = map[string]*Keyword{}
	return t
}

func (k *Keyword) Run(driver driver.TestDriver, knownCommands map[string]*Keyword, args []string) bool {
	k.Ran = true
	for rowNumber := 0; rowNumber < len(k.Commands); rowNumber++ {

		command := k.Commands[rowNumber]

		commandTmp := make([]string, len(command))
		copy(commandTmp, command)

		skipSleep := false
		result := true
		skipFail := false

		for i := range command {
			c := &command[i]

			if i == 0 && strings.HasPrefix(*c, "!") {
				skipFail = true
				commandTmp[i] = strings.Trim(*c, "!")
			}

			re := regexp.MustCompile("\\$\\$([0-9]+)\\$\\$")
			// TODO: Make it possible to escape $$.
			var templateArgs = re.FindAllStringSubmatch(*c, -1)

			if templateArgs != nil {
				for _, row := range templateArgs {
					var argn, _ = strconv.Atoi(row[1])

					if argn >= len(args) {
						log.Errorf("%s: Not enough arguments given, replacing with empty string.", commandTmp[0])
						commandTmp[i] = strings.Replace(*c, "$$"+row[1]+"$$", "", -1)
					} else {
						commandTmp[i] = strings.Replace(*c, "$$"+row[1]+"$$", args[argn], -1)
					}
				}
			}

			re = regexp.MustCompile("\\$\\$([A-Za-z][A-Za-z0-9]*)\\$\\$")
			// TODO: Make it possible to escape $$.
			templateArgs = re.FindAllStringSubmatch(*c, -1)

			if templateArgs != nil {
				for _, row := range templateArgs {
					// TODO check that variables exist.
					commandTmp[i] = strings.Replace(*c, "$$"+row[1]+"$$", k.Variables[row[1]], -1)
				}
			}

		}

		if commandTmp[0] == "click" {
			result = driver.Click(false, commandTmp[1])
		} else if commandTmp[0] == "click_x" {
			result = driver.Click(true, commandTmp[1])
		} else if commandTmp[0] == "click_on_text" {
			result = driver.ClickOnText("", commandTmp[1])
		} else if matches, _ := regexp.MatchString("click_[0-9]nth", commandTmp[0]); matches {
			re := regexp.MustCompile("[0-9]")
			nth, _ := strconv.Atoi(re.FindAllString(commandTmp[0], -1)[0])
			result = driver.ClickNth(commandTmp[1], nth)
		} else if commandTmp[0] == "click_closest_to" {
			result = driver.ClickClosestTo(commandTmp[1], commandTmp[2])
		} else if commandTmp[0] == "navigate" {
			result = driver.Navigate(commandTmp[1])
		} else if commandTmp[0] == "new_page" {
			result = driver.NewPage(commandTmp[1])
		} else if commandTmp[0] == "go_to_window" {
			nthWindow, _ := strconv.Atoi(commandTmp[1])
			result = driver.GoToNthWindow(nthWindow)
		} else if commandTmp[0] == "has_text" {
			if len(commandTmp) == 3 {
				result = driver.HasText(commandTmp[1], commandTmp[2])
			} else {
				result = driver.HasText(commandTmp[1], "")
			}
		} else if commandTmp[0] == "input" {
			if len(commandTmp) == 3 {
				result = driver.Input(false, commandTmp[1], commandTmp[2])
			} else {
				result = driver.InputEmpty(commandTmp[1])
			}
		} else if commandTmp[0] == "enter" {
			result = driver.SendKey("EnterKey")
		} else if commandTmp[0] == "backspace" {
			result = driver.SendKey("BackspaceKey")
		} else if commandTmp[0] == "sendkey" {
			result = driver.SendKey(commandTmp[1])
		} else if commandTmp[0] == "input_x" {
			result = driver.Input(true, commandTmp[1], commandTmp[2])
		} else if commandTmp[0] == "sleep" {
			sleepTime, _ := strconv.Atoi(commandTmp[1])
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		} else if commandTmp[0] == "log" {
			result = driver.Log(commandTmp[1])
		} else if commandTmp[0] == "label" {
			log.Logf("label, ´%s´", commandTmp[1])
			k.LabelLocation[commandTmp[1]] = rowNumber
			skipSleep = true
		} else if commandTmp[0] == "jump" {
			log.Logf("jump, ´%s´", commandTmp[1])
			rowNumber = k.LabelLocation[commandTmp[1]] - 1
			skipSleep = true
		} else if commandTmp[0] == "read" {
			var value string
			value, result = driver.Read(commandTmp[1])
			k.Variables[commandTmp[2]] = value
		} else if commandTmp[0] == "testcase" {
			log.Printf("\n🍤 Running testcase: %s \n", commandTmp[1])
			skipSleep = true
		} else if commandTmp[0] == "stop_if_success" {
			if k.LastResult == true {
				return false
			}
		} else if commandTmp[0] == "execute" {
			skipSleep = true
			out, err := utils.Execute(commandTmp[1])
			if err != nil {
				log.Errorf("RunList %s failed, %s", commandTmp[1], err)
			}
			log.Log("Execute", "´"+commandTmp[1]+"´")
			log.Println(aurora.Bold(out))
		} else {
			keyword := knownCommands[commandTmp[0]]
			if keyword == nil {
				log.Errorf("No keyword named ´%s´", commandTmp[0])
				return false
			}
			keyword.SkipFail = skipFail

			exitOk := keyword.Run(driver, knownCommands, commandTmp[1:])
			if !exitOk {
				return false
			}
		}

		if result == false && !k.SkipFail {
			driver.Stop()
			k.Passed = false
			return false
		}

		if result == false && k.SkipFail {
			driver.Stop()
			k.Passed = false
			return true
		}

		if skipSleep {
			time.Sleep(150 * time.Millisecond)
		}

		k.LastResult = result

	}

	if k.IsATest {
		k.Passed = true
	}

	return true
}

func (c *Case) Run() {
	c.Driver.Start()
	exitOk := c.RootKeyword.Run(c.Driver, c.KnownKeywords, nil)
	c.Driver.Stop()

	printResults := false
	hasFailedTest := false
	for i := range c.KnownKeywords {
		command := c.KnownKeywords[i]

		if !command.Ran {
			continue
		}

		if command.IsATest && printResults == false {
			fmt.Print("\n\nResults: \n\n")
			printResults = true
		}

		if command.IsATest && command.Passed {
			fmt.Printf("%s %s\n", aurora.Green("✓"), command.TestCaseName)
		} else if command.IsATest && !command.Passed {
			hasFailedTest = true
			fmt.Printf("%s %s\n", aurora.Red("✗"), command.TestCaseName)
		}
	}

	if !exitOk {
		os.Exit(1)
	} else if exitOk && hasFailedTest {
		os.Exit(2)
	}
}

func (c *Case) StartServer(port int) {
	c.Driver.Start()

	c.Listen(port)
	// t.RunList.Run(t.Driver, t.KnownCommands, nil)

	// t.Driver.Stop()

}
