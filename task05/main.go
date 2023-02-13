package main

// Задание 5
// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// writer отправляет последовательные числа в канал. Горутина завершается при истечении контекста.
func writer(ctx context.Context, ch chan<- int) {
	i := 0
	for { // выполнять до лучших времён
		select {
		case <-ctx.Done():
			close(ch) // закрытие канала является сигналом завершения для ридера
			return
		default:
			i+=1
			ch <- i
			time.Sleep(100*time.Millisecond) // задержка для наглядности
		}
	}
}

// reader читает числа из канала до посинения (зачёркнуто) до его закрытия.
func reader(ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		fmt.Println(num)
	}
	wg.Done() // отпускаем главный поток
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // создаём контекст, который завершится через 5 сек.
	defer cancel() // контекст гарантированно закроется
	var wg sync.WaitGroup // последнее звено в цепи остановок - ридер, wg нужна для него
	ch := make(chan int, 1)
	wg.Add(1)
	go writer(ctx, ch)
	go reader(ch, &wg)
	wg.Wait() // ждем остановки ридера
}
