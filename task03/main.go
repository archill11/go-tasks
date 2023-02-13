package main

// Задание 3
// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+4^2+6^2….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// useAtomic использует возможности пакета sync.Atomic.
func useAtomic(numbers []int, w *sync.WaitGroup) {
	defer w.Done()
	var result int64
	var wg sync.WaitGroup // используем sync.WaitGroup, чтобы дождаться окончания работы воркеров.
	wg.Add(len(numbers))
	for _, num := range numbers {
		go func(n int) {
			defer wg.Done()
			atomic.AddInt64(&result, int64(n*n)) // atomic гарантирует, что не будет гонки данных
		}(num)
	}
	wg.Wait() // ждём окончания
	fmt.Printf("useAtomic: result: %d\n", int(result))
}

// useChannels использует передачу данных и результата по каналам.
func useChannels(numbers []int, w *sync.WaitGroup) {
	defer w.Done()
	var wg sync.WaitGroup
	chSqr := make(chan int, len(numbers)) // канал для передачи квадратов
	chResult := make(chan int)            // канал для результата
	wg.Add(len(numbers))
	for _, num := range numbers {
		go func(n int) {
			defer wg.Done()
			chSqr <- (n * n) // передаем квадрат в канал
		}(num)
	}
	go func() { // эта горутина собирает квадраты из канала и складывает их
		var result int
		for num := range chSqr {
			result += num
		}
		// когда закроем канал
		chResult <- result
	}()
	wg.Wait() // ждём окончания
	close(chSqr) // закрытие канала даёт сигнал горутине послать результат
	fmt.Printf("useChannels: result: %d\n", <-chResult)
	close(chResult)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg1 sync.WaitGroup
	wg1.Add(2)
	go useAtomic(nums, &wg1)
	go useChannels(nums, &wg1)
	wg1.Wait()

	
	const sliceSize = 79999999 // этот пример деменстрирует что атомик выполняется быстрее на больших вычислениях
	numbers := make([]int, 0, sliceSize)
	for i := 1; i <= sliceSize; i++ {
		numbers = append(numbers, i*2)
	}
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go useAtomic(numbers, &wg2)
	go useChannels(numbers, &wg2)
	wg2.Wait()
}
