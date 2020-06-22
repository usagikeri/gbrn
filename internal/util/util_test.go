package util_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/usagikeri/gbrn/internal/util"
)

func TestPwdFilesSuccess(t *testing.T) {
	FILENAME := "testfile"
	expect := []string{"util.go", "util_test.go"}
	actual := util.PwdFiles(FILENAME)

	for i, _ := range expect {
		if !reflect.DeepEqual(expect[i], actual[i]) {
			t.Fatal("error")
		}
	}

	defer func() {
		if err := os.Remove(FILENAME); err != nil {
			fmt.Println(err)
		}
	}()
}

func TestPwdFilesFaild(t *testing.T) {
	FILENAME := "testfile"
	expect := []string{"editor.go", "util.go"}
	actual := util.PwdFiles(FILENAME)

	for i, _ := range expect {
		if reflect.DeepEqual(expect[i], actual[i]) {
			t.Fatal("error")
		}
	}

	defer func() {
		if err := os.Remove(FILENAME); err != nil {
			fmt.Println(err)
		}
	}()
}
