package main

import (
	"fmt"
	"os"
	"tictactoe/src/field"
	"tictactoe/src/input-manager"
)

func main() {
	fmt.Printf("Please enter symbol you want to use: \n")
	humanSymbol := string(input_manager.TakeInput())
	fmt.Printf("Entered %v...\n", humanSymbol)
	fmt.Printf("Please enter symbol you want computer to use: \n")
	pcSymbol := string(input_manager.TakeInput())
	fmt.Printf("Entered %v...\n", pcSymbol)

	f := field.NewField(humanSymbol, pcSymbol)
	f.DisplayField()

	for {
		var key rune = input_manager.TakeInput()

		if key == 113 {
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		} else if key == 119 {
			f.ChangeSelected(0)
		} else if key == 100 {
			f.ChangeSelected(1)
		} else if key == 115 {
			f.ChangeSelected(2)
		} else if key == 97 {
			f.ChangeSelected(3)
		} else if key == 13 {
			f.Select()
		}

		f.DisplayField()
	}
}
