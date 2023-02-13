package main

// Задание 16
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivotIndex := len(arr) - 1 // индекс опорного элемента 
	pivot := arr[pivotIndex]   // опорный элемент 
	less := make([]int, 0)     // числа меньше опорного
	greater := make([]int, 0)  // числа больше опорного
	for i, el := range arr {
		if i == pivotIndex {
			continue
		}
		if el < pivot {
			less = append(less, el)
		}else{
			greater = append(greater, el)
		}
	}
	res := make([]int, 0, len(arr))
	res = append(res, quickSort(less)...)
	res = append(res, pivot)
	res = append(res, quickSort(greater)...)
	return res
}

func main() {
	arr := []int{8, 19, 120, 55, 78, 4, 8, 9, 44, 444, 6, 1, -5, -100, 5858, 5, 22, 46, 44, 120, 55, 78, 4, 8, 9, 44, 444, 6, 1, -5, -100, 5858, 5, 22, 46, 44}
	fmt.Println( quickSort(arr) )
}