package main

// Задание 8
// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

import "fmt"

// setBit1 устанавливает i-й бит в числе num в значение 1
func setBit1(num int64, i int) int64 {
	var mask int64 = 00000001 << i // 00001000
	num = num | mask // если нужно установить, то or
	// 00000001
	//    or
	// 00001000
	//   =
	// 00001001
	return num
}
// setBit1 устанавливает i-й бит в числе num в значение 0
func setBit0(num int64, i int) int64 {
	var mask int64 = 00000001 << i // 00001000
	num = num ^ mask // если сбросить, то xor
	// 00000001
	//   xor
	// 00001000
	//    =
	// 00000001
	return num
}

func main() {
	var n int64 = 1
	n = setBit1(n, 10)
	fmt.Println(n)

	n = setBit0(n, 10)
	fmt.Println(n)
}
