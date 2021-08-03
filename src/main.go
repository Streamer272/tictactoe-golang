package main

import "tictactoe/src/field"

func main() {
	f := field.NewField()
	f.DisplayField()
	f.ChangeBox(5, "X")
	f.DisplayField()
}
