package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/usagikeri/gbrn/internal/editor"
	"github.com/usagikeri/gbrn/internal/util"
)

var FILENAME string

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error)
	}
	FILENAME = filepath.Join(path, ".tempFiles")
}

func main() {
	pwdFiles := util.PwdFiles(FILENAME)
	err := editor.OpenEditor(FILENAME)
	if err != nil {
		log.Fatal(err.Error)
	}
	rename := util.File2slice(FILENAME)

	flag := util.CheckDuplicate(rename)
	if !flag {
		log.Fatal("Error: File name duplicate\n")
	}

	if len(pwdFiles) != len(rename) {
		log.Fatal("Error: Invalid format\n")
	}

	for i, _ := range pwdFiles {
		if pwdFiles[i] != rename[i] {
			err := os.Rename(pwdFiles[i], rename[i])
			if err != nil {
				log.Fatal(err.Error)
			}
		}
	}

	defer func() {
		if err := os.Remove(FILENAME); err != nil {
			fmt.Println(err)
		}
	}()
}
