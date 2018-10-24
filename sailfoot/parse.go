package sailfoot

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/adl32x/sailfoot/log"
)

func (c *Case) Load(filename string) {
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
			log.Debug("Loaded file: ", path)
			keyword := strings.Replace(filepath.Base(path), ".txt", "", -1)
			walkedFile, _ := ioutil.ReadFile(path)
			c.KnownKeywords[keyword] = fileToCommands(walkedFile)
		}

		return nil
	})

	c.RootKeyword = fileToCommands(file)
}

func fileToCommands(file []byte) Keyword {
	str := string(file)
	log.Debug("File content: ", str)
	rows := strings.Split(str, "\n")
	c := Keyword{}
	c.Init()

	for _, row := range rows {
		row = strings.Trim(row, " \t")
		if len(row) == 0 {
			continue
		}

		if string(row[0]) == "#" {
			continue
		}

		command := SplitLine(row)

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
