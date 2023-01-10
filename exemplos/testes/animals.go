package testes

type Animal struct {
	Name    string
	Species string
}

func (a Animal) IsDog() bool {
	if a.Species == "dog" {
		return true
	}

	return false
}

func (a Animal) HasName() bool {
	if len(a.Name) > 0 {
		return true
	}

	return false
}
