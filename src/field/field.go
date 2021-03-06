package field

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"math/rand"
)

type Box struct {
	Position int
	Value    string
}

type Field struct {
	Boxes       [9]Box
	Selected    int
	HumanSymbol string
	PcSymbol    string
}

func NewField(humanSymbol string, pcSymbol string) Field {
	f := Field{Selected: 4, HumanSymbol: humanSymbol, PcSymbol: pcSymbol}

	for i := 0; i < 9; i++ {
		f.Boxes[i] = Box{i, "0"}
	}

	return f
}

func (f *Field) DisplayField() {
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

func (f *Field) ChangeSelected(direction int) {
	if direction == 0 {
		nextSelected := f.Selected - 3

		if nextSelected >= 0 {
			f.Selected = nextSelected
		}
	} else if direction == 1 {
		nextSelected := f.Selected + 1

		if nextSelected%3 != 0 {
			f.Selected = nextSelected
		}
	} else if direction == 2 {
		nextSelected := f.Selected + 3

		if nextSelected <= 8 {
			f.Selected = nextSelected
		}
	} else if direction == 3 {
		nextSelected := f.Selected - 1

		if (nextSelected-2)%3 != 0 {
			f.Selected = nextSelected
		}
	}
}

func (f *Field) Select() bool {
	if f.Boxes[f.Selected].Value == "0" {
		f.Boxes[f.Selected].Value = f.HumanSymbol
		return true
	}

	return false
}

func (f *Field) CheckForWinner() string {
	// rows
	if f.Boxes[0].Value == f.Boxes[1].Value && f.Boxes[1].Value == f.Boxes[2].Value && f.Boxes[0].Value != "0" {
		return f.Boxes[0].Value
	} else if f.Boxes[3].Value == f.Boxes[4].Value && f.Boxes[4].Value == f.Boxes[5].Value && f.Boxes[3].Value != "0" {
		return f.Boxes[3].Value
	} else if f.Boxes[6].Value == f.Boxes[7].Value && f.Boxes[7].Value == f.Boxes[8].Value && f.Boxes[6].Value != "0" {
		return f.Boxes[6].Value
	}

	// columns
	if f.Boxes[0].Value == f.Boxes[3].Value && f.Boxes[3].Value == f.Boxes[6].Value && f.Boxes[0].Value != "0" {
		return f.Boxes[0].Value
	} else if f.Boxes[1].Value == f.Boxes[4].Value && f.Boxes[4].Value == f.Boxes[7].Value && f.Boxes[1].Value != "0" {
		return f.Boxes[1].Value
	} else if f.Boxes[2].Value == f.Boxes[5].Value && f.Boxes[5].Value == f.Boxes[8].Value && f.Boxes[2].Value != "0" {
		return f.Boxes[2].Value
	}

	// other
	if f.Boxes[0].Value == f.Boxes[4].Value && f.Boxes[4].Value == f.Boxes[8].Value && f.Boxes[0].Value != "0" {
		return f.Boxes[0].Value
	} else if f.Boxes[2].Value == f.Boxes[4].Value && f.Boxes[4].Value == f.Boxes[6].Value && f.Boxes[2].Value != "0" {
		return f.Boxes[2].Value
	}

	return ""
}

func (f *Field) ChangeRandomBox() {
	for {
		random := rand.Intn(8)

		if f.Boxes[random].Value == "0" {
			f.Boxes[random].Value = f.PcSymbol
			return
		}
	}
}

func (f *Field) IsFieldEmpty() bool {
	for i := 0; i < 9; i++ {
		if f.Boxes[i].Value == "0" {
			return false
		}
	}

	return true
}
