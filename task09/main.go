package main

// Задание 9
// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

import (
	"fmt"
	"time"
)

// producer извлекает числа из переданного массива и посылает их в канал prodStream.
func producer(arr []int, prodStream chan<- int) {
	for _, x := range arr {
		prodStream <- x
	}
	close(prodStream)
}

// multiplier читает числа из канала chIn, умножает их на 2 и передаёт в канал multStream.
func multiplier(prodStream <-chan int, multStream chan<- int) {
	for x := range prodStream {
		multStream <- (x * 2)
		time.Sleep(10*time.Millisecond)
	}
	close(multStream)
}

func main() {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}
	prodStream := make(chan int)
	go producer(arr, prodStream)
	multStream := make(chan int)
	go multiplier(prodStream, multStream)
	for x := range multStream {
		fmt.Println(x)
	}
}
