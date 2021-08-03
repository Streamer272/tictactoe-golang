package field

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
		f.Boxes[i] = Box{i, ""}
	}

	return f
}
