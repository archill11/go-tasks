package main

// Задание 2
// Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых
// из массива (2,4,6,8,10) и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // используем sync.WaitGroup, чтобы дождаться окончания работы воркеров.
	wg.Add(len(nums))
	for _, num := range nums {
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n)
		}(num)
	}
	wg.Wait() // ждём окончания
}
