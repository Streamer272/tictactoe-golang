package input_manager

import (
	"github.com/mattn/go-tty"
)

func TakeInput() rune {
	reader, _ := tty.Open()
	defer reader.Close()

	r, _ := reader.ReadRune()
	return r
}
