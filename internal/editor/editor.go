package editor

import (
	"os"
	"os/exec"
)

func OpenEditor(FILENAME string) error {
	editor := "vim"

	c := exec.Command(editor, FILENAME)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()

	return err
}
