package util_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/usagikeri/gbrn/internal/util"
)

func TestPwdFiles(t *testing.T) {
	FILENAME := "testfile"
	tests := []struct {
		files []string
		want  bool
	}{
		{[]string{"testdata", "util.go", "util_test.go"}, true},
		{[]string{"testdata", "util.go", "editor.go"}, false},
	}

	for _, tt := range tests {
		if got := reflect.DeepEqual(util.PwdFiles(FILENAME), tt.files); got != tt.want {
			t.Fatalf("want = %t, but %t", tt.want, got)
		}
	}

	defer func() {
		if err := os.Remove(FILENAME); err != nil {
			fmt.Println(err)
		}
	}()
}

func TestFile2Slice(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Errorf("%t", err)
	}
	FILENAME := filepath.Join(path, "/testdata/test.txt")

	tests := []struct {
		files []string
		want  bool
	}{
		{[]string{"util.go", "util_test.go"}, true},
		{[]string{"util.go", "editor.go"}, false},
	}
	for _, tt := range tests {
		if got := reflect.DeepEqual(util.File2slice(FILENAME), tt.files); got != tt.want {
			t.Fatalf("want = %t, but %t", tt.want, got)
		}
	}
}

func TestCheckDuplicate(t *testing.T) {
	tests := []struct {
		chars []string
		want  bool
	}{
		{[]string{"aaa", "bbb", "ccc"}, true},
		{[]string{"aaa", "bbb", "aaa"}, false},
	}
	for _, tt := range tests {
		if got := util.CheckDuplicate(tt.chars); got != tt.want {
			t.Fatalf("want = %t", tt.want)
		}
	}
}
