package field

import "fmt"

type Box struct {
	Position int
	Value    string
}

type Field struct {
	Boxes [9]Box
}

func NewField() Field {
	f := Field{}

	for i := 0; i < 9; i++ {
		f.Boxes[i] = Box{i, "0"}
	}

	return f
}

func (f Field) DisplayField() {
	fmt.Printf("\n")

	for i := 0; i < 9; i++ {
		box := f.Boxes[i]

		if (i-2)%3 == 0 {
			fmt.Printf("%v\n", box.Value)

			if i != 8 {
				fmt.Printf("---------\n")
			}
		} else if i%3 == 0 {
			fmt.Printf("%v", box.Value)
		} else {
			fmt.Printf(" | %v | ", box.Value)
		}
	}

	fmt.Printf("\n")
}

func (f *Field) ChangeBox(index int, value string) {
	f.Boxes[index].Value = value
}
