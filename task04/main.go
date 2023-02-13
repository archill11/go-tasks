package main

// Задание 4
// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// chWriter запускает воркеры и пишет случайные числа в канал, дожидаясь отмены контекста.
// Остановка воркеров реализована через закрытие канала, это наиболее популярный паттерн для такой задачи.
func chWriter(ctx context.Context, n int) {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 1; i <= n; i++ { // запускаем воркеры
		go worker(i, ch, &wg)
	}
	for {
		select {
			case <-ctx.Done(): // контекст завершён!
			close(ch) // закрытие канала служит сигналом остановки одновременно всем воркерам
			wg.Wait() // дожидаемся, пока воркеры завершат работу
			return
		default:
			ch <- rand.Int()
		}
	}
}

// worker читает данные из канала и выводит их в stdout.
func worker(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("Worker %d started", i)
	for number := range ch { // закрытие канала остановит все запущенные воркеры
		time.Sleep(1*time.Second) // задержка для наглядности
		fmt.Println(number)
	}
	log.Printf("Worker %d stopped", i)
}

func main() {
	var arg string
	if len(os.Args) != 2 { // если при запуске не передан аргумент 
		arg = "3"
	}else{
		arg = os.Args[1]
	}
	numOfWorkers, err := strconv.Atoi(arg)
	if err != nil || numOfWorkers < 1 {
		fmt.Printf("wrong number of workers: %s\n", arg)
		os.Exit(1)
	}
	ctx, cancel := context.WithCancel(context.Background()) // создаём контекст с отменой для плавного завершения
	sigint := make(chan os.Signal, 1) // подписываемся на сигнал остановки от ОС
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT)

	go func() { // эта горутина ждёт сигнала от ОС и завершает контекст, давая тем самым сигнал остановить работу
		<-sigint
		log.Println("Shutting down...")
		cancel()
	}()
	chWriter(ctx, numOfWorkers)
	fmt.Println("Service stopped...")
}
