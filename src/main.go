package main

import (
	"fmt"
	"tictactoe/src/field"
)

func main() {
	f := field.NewField()
	fmt.Printf("%v\n", f.Boxes[0])
}
