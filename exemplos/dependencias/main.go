package main

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/jpgsaraceni/gopher-alpha/exemplos/dependencias/sum"
)

func main() {
	id := uuid.NewString()
	fmt.Printf("Execution ID: %s", id)

	const num1 = 1
	const num2 = 2

	num3 := sum.Sum(num1, num2)
	fmt.Println(num3)
}
