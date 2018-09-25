package main

import (
	"fmt"
)

// Plusable provides plus behaviour
type Plusable struct{}

func (p Plusable) Plus(a, b int) int {
	return a + b
}

// Minusable provides minus behaviour
type Minusable struct{}

func (m Minusable) Minus(a, b int) int {
	return a - b
}

type Calculator struct {
	Plusable  // Embed plus behaviour
	Minusable // EMbed minus behaviour
}

func main() {
	calc := Calculator{}
	a := 10
	b := 3
	fmt.Printf("%d + %d = %d\n", a, b, calc.Plus(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, calc.Minus(a, b))
}
