package util

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func PwdFiles(FILENAME string) []string {
	contents := ""
	var lines []string

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, file := range files {
		name := file.Name()
		if name != FILENAME {
			contents += name + "\n"
			lines = append(lines, name)
		}
	}

	err = ioutil.WriteFile(FILENAME, []byte(contents), os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}

	return lines
}

func File2slice(FILENAME string) []string {
	f, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatal(err.Error())
	}
	temp := strings.Split(string(f), "\n")
	return temp[:len(temp)-1]
}

func CheckDuplicate(args []string) bool {
	checkmap := map[string]bool{}
	for _, key := range args {
		if !checkmap[key] {
			checkmap[key] = true
		} else {
			return false
		}
	}
	return true
}
