package main

import (
	"fmt"
	"os"
	"tictactoe/src/field"
	"tictactoe/src/input-manager"
)

func main() {
	// intro
	fmt.Printf("Welcome to TicTacToe!\n")
	fmt.Printf("How to play the game:\n")
	fmt.Printf("    You can use WASD to move your cursor, and ENTER to select\n")
	fmt.Printf("    You can press \"Q\" to quit the game\n")
	fmt.Printf("Press any key to start the game!\n\n")
	input_manager.TakeInput()

	// start
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

		if key == 113 { // q
			fmt.Printf("Exiting...\n")
			os.Exit(0)
		} else if key == 119 { // w
			f.ChangeSelected(0)
		} else if key == 100 { // d
			f.ChangeSelected(1)
		} else if key == 115 { // s
			f.ChangeSelected(2)
		} else if key == 97 { // a
			f.ChangeSelected(3)
		} else if key == 13 { // enter
			f.Select()
		}

		f.DisplayField()
	}
}
