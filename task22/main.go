package main

import (
	"fmt"
	"math/big"
)

// Задание 22
// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := big.NewInt(22)
	a.Exp(a, big.NewInt(40), nil)
	fmt.Printf("a::: %d\n", a)

	b := big.NewInt(23)
	b.Exp(b, big.NewInt(30), nil)
	fmt.Printf("b::: %d\n", b)
	
	aPlusB := a.Add(a, b)
	fmt.Printf("plus::: %d\n", aPlusB)

	aMinusB := a.Sub(a, b)
	fmt.Printf("munis::: %d\n", aMinusB)
	
	aMultB := a.Mul(a, b)
	fmt.Printf("mul::: %d\n", aMultB)

	abDiv := a.Div(a, b)
	fmt.Printf("div::: %d\n", abDiv)
}
