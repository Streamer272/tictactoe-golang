package main

import (
	"fmt"
	"os"
	"tictactoe/src/field"
	"tictactoe/src/input-manager"
)

func GetHumanSymbol() string {
	fmt.Printf("Please enter symbol you want to use: \n")
	humanSymbol := string(input_manager.TakeInput())

	if string(humanSymbol) == "0" || string(humanSymbol) == "\n" {
		fmt.Printf("Sorry, that is not a valid symbol. Please enter another symbol...\n")
		humanSymbol = GetHumanSymbol()
	}

	return humanSymbol
}

func GetPcSymbol() string {
	fmt.Printf("Please enter symbol you want pc to use: \n")
	pcSymbol := string(input_manager.TakeInput())

	if string(pcSymbol) == "0" || string(pcSymbol) == "\n" {
		fmt.Printf("Sorry, that is not a valid symbol. Please enter another symbol...\n")
		pcSymbol = GetPcSymbol()
	}

	return pcSymbol
}

func StartGame(humanSymbol string, pcSymbol string) {
	f := field.NewField(humanSymbol, pcSymbol)
	f.DisplayField()

	running := true

	for running {
		key := input_manager.TakeInput()
		selected := false

		if key == 113 { // q
			fmt.Printf("Are you sure you want to quit? [Y\\n] ")
			key := input_manager.TakeInput()
			fmt.Printf("\n")
			if key == 121 || key == 89 {
				fmt.Printf("Exiting...\n")
				os.Exit(0)
			}
		} else if key == 119 { // w
			f.ChangeSelected(0)
		} else if key == 100 { // d
			f.ChangeSelected(1)
		} else if key == 115 { // s
			f.ChangeSelected(2)
		} else if key == 97 { // a
			f.ChangeSelected(3)
		} else if key == 13 { // enter
			if f.Select() {
				selected = true
			}
		}

		f.DisplayField()

		winner := f.CheckForWinner()
		if winner != "" {
			fmt.Printf("And the winner is... \"%v\"!!\n", winner)
			running = false
		}

		isFieldEmpty := f.IsFieldEmpty()
		if isFieldEmpty {
			fmt.Printf("Field is full, no one is winner...\n")
			running = false
		}

		if running && selected {
			f.ChangeRandomBox()
			f.DisplayField()
		}
	}

	fmt.Printf("Do you want to play again? [Y\\n] ")
	key := input_manager.TakeInput()
	fmt.Printf("\n")
	if key == 121 || key == 89 {
		StartGame(humanSymbol, pcSymbol)
		return
	}
}

func main() {
	fmt.Printf("Welcome to TicTacToe!\n")
	fmt.Printf("How to play the game:\n")
	fmt.Printf("    You can use WASD to move your cursor, and ENTER to select\n")
	fmt.Printf("    You can press \"Q\" to quit the game\n")
	fmt.Printf("Press any key to start the game!\n\n")
	input_manager.TakeInput()

	StartGame(GetHumanSymbol(), GetPcSymbol())
}
