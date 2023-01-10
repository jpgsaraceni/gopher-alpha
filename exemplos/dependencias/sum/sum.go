package sum

import "fmt"

func Sum(terms ...int) int {
	var sum int
	for _, term := range terms {
		sum += term
	}

	return sum
}

func notExported() {
	fmt.Printf("essa função não pode ser chamada em outro pkg.")
}
