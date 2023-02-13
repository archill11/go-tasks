package main

// Задание 11
// Реализовать пересечение двух неупорядоченных множеств.

import "fmt"

// intersection возвращает массив, включающий в себя все элементы, которые содержатся в обоих переданных массивах.
func intersection[T comparable](setA, setB []T) []T {
	result := make([]T, 0, len(setA))
	mapB := make(map[T]bool) // преобразуем один из массивов в map
	for _, item := range setB {
		mapB[item] = true
	}
	for _, item := range setA { // перебираем все элементы оставшегося массива и ищем совпадения в мапе
		if _, ok := mapB[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	res:= intersection(
		[]int{2, 6, 8, 4, 12, 14, 10, 18, 16, 33},
		[]int{2, 6, 8, 4, 12, 14, 10, 19},
	)
	fmt.Println(res)
}
