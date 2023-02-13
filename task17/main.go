package main

// Задание 17
// Реализовать бинарный поиск встроенными методами языка.

import (
	"fmt"
	"math"
)

// итеративный способ реализации
func binarySearch(arr []int, item int) int {
	start := 0
	end := len(arr)
	var middle int
	position := -1 // позиция искомого элемента
	for start <= end {
		middle = int(math.Floor((float64(start + end)) / 2)) // вычисляем позицию элемента по середине
		if arr[middle] == item {                             // если элемент найден, возвращаем
			return middle
		}
		if item < arr[middle] { // если элемент меньше, того который по середине
			end = middle - 1    // ищем в другой половине массива
		} else {                // если элемент больше, того который по середине
			start = middle + 1  // ищем в другой половине массива
		}
	}
	return position
}

// рекурсивный способ реализации
func recursiveBinarySearch(arr []int, item, start, end int) int {
	middle := int(math.Floor((float64(start + end)) / 2)) // вычисляем позицию элемента по середине
	if item == arr[middle] {                              // если элемент найден, возвращаем
		return middle
	}
	if item < arr[middle] {                                      // если элемент меньше, того который по середине 
		return recursiveBinarySearch(arr, item, start, middle-1) // ищем в другой половине массива
	} else {                                                     // если элемент больше, того который по середине  
		return recursiveBinarySearch(arr, item, middle+1, end)   // ищем в другой половине массива
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 34, 36, 37, 38, 39, 40}
	fmt.Println(binarySearch(arr, 33))
	fmt.Println(recursiveBinarySearch(arr, 33, 0, len(arr)))
}
