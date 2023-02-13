package main

// Задание 13
// Поменять местами два числа без создания временной переменной.

import "fmt"

func main() {
	a := 15
	b := 7

	a = a + b // с помощью сложения и вычитания:
	b = a - b
	a = a - b
	fmt.Println(a, b)

	if a != 0 && b != 0 { // с помощью умножения и деления:
		a = a * b
		b = a / b
		a = a / b
		fmt.Println(a, b)
	}

	a = a ^ b // с помощью XOR:
	b = b ^ a
	a = a ^ b
	fmt.Println(a, b)
}
