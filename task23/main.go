package main

// Задание 23
// Удалить i-ый элемент из слайса.

import "fmt"


func deleteElement[T comparable](arr []T, i int) []T {
	if i < 0 || i > len(arr) {
		fmt.Println("index out of range")
	}
	return append(arr[:i], arr[i+1:]...) // отрезаем до удаляемого элемента и добавляем то, что после
}

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n:= deleteElement(ints, 5)
	fmt.Println(ints)
	fmt.Println(n)

	strings := []string{"zero", "one", "two", "three", "four", "five"}
	s:= deleteElement(strings, 3)
	fmt.Println(strings)
	fmt.Println(s)
}
