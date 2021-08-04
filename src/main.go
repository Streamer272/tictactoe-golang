package main

import (
	"fmt"
	"os"
	"tictactoe/src/field"
	"tictactoe/src/input-manager"
)

func main() {
	f := field.NewField()
	f.DisplayField()

	for {
		var key rune = input_manager.TakeInput()

		/*
		Key press => q
		113
		Key press => w
		119
		Key press => d
		100
		Key press => a
		97
		Key press => s
		115
		Key press =>
		13
		*/

		if key == 113 {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		}

		f.DisplayField()
	}
}
