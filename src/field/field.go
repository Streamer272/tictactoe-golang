package field

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
)

type Box struct {
	Position int
	Value    string
}

type Field struct {
	Boxes    [9]Box
	Selected int
}

func NewField() Field {
	f := Field{Selected: 4}

	for i := 0; i < 9; i++ {
		f.Boxes[i] = Box{i, "0"}
	}

	return f
}

func (f Field) DisplayField() {
	width, height, _ := terminal.GetSize(0)

	for i := 0; i < int((height-5)/2); i++ {
		fmt.Printf("\n")
	}

	for i := 0; i < 9; i++ {
		box := f.Boxes[i]

		if (i-2)%3 == 0 {
			if f.Selected == i {
				fmt.Printf("[ %v ]\n", box.Value)
			} else if f.Selected == i-1 {
				fmt.Printf("] %v\n", box.Value)
			} else {
				fmt.Printf("| %v\n", box.Value)
			}

			if i != 8 {
				for j := 0; j < int((width-9)/2); j++ {
					fmt.Printf(" ")
				}

				fmt.Printf("---------\n")
			}
		} else if i%3 == 0 {
			for j := 0; j < int((width-9)/2)-2; j++ {
				fmt.Printf(" ")
			}

			if f.Selected == i {
				fmt.Printf("[ %v ]", box.Value)
			} else if f.Selected == i+1 {
				fmt.Printf("  %v [", box.Value)
			} else {
				fmt.Printf("  %v |", box.Value)
			}
		} else {
			fmt.Printf(" %v ", box.Value)
		}
	}

	for i := 0; i < int((height-5)/2); i++ {
		fmt.Printf("\n")
	}
}

func (f *Field) ChangeBox(index int, value string) {
	f.Boxes[index].Value = value
}
