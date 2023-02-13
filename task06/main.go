package main

// Задание 6
// Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 1. Самостоятельно. Горутина завершается сама, выполнив возложенные на неё задачи. По окончании
// она уведомляет о своём завершении, декрементируя пераданную WaitGroup (так же и в остальных примерах).
func selfStop(wg *sync.WaitGroup) {
	log.Println("=selfStop= started")
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
	log.Println("=selfStop= stopped")
	wg.Done()
}

// 2. Контекст. Сигналом для выхода является истечение переданного контекста.
func contextStop( ctx context.Context, wg *sync.WaitGroup) {
	log.Printf("==contextStop== started")
	for {
		select {
		case <-ctx.Done(): // конекст завершён
			log.Printf("==contextStop== stopped")
			wg.Done()
			return
		default:
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}
}

// 3. Закрытие канала. Сигналом к завершению работы всех горутин блока является закрытие переданного канала.
// Можно также посылать в канал сообщения, закрывая воркеры по одиночке.
func chanCloseStop( done <-chan struct{}, wg *sync.WaitGroup) {
	log.Printf("===chanCloseStop=== started")
	for {
		select {
		case <-done: // сообщение в канал либо закрытие канала
			log.Printf("===chanCloseStop=== stopped")
			wg.Done()
			return
		default:
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)))
		}
	}
}

// 4. WaitGroup. Воркеры в блоке ждут, когда обнулится переданный sync.WaitGroup.
func waitgroupStop(wgStop, wg *sync.WaitGroup) {
	log.Printf("====waitgroupStop==== started")
	wgStop.Wait()
	log.Printf("====waitgroupStop==== stopped")
	wg.Done()
}

// 5. Condition. Горутины ждут выполнения sync.Condition. Получив сигнал, они завершаются. Отправитель сигнала может
// завершить часть воркеров, передав несколько сообщений cond.Signal, либо завершить всех слушателей, вызвав cond.Broadcast.
// condStop - завершение по общему сигналу в sync.Cond.
func condStop(cond *sync.Cond, wg *sync.WaitGroup) {
	log.Printf("======condStop====== started")
	cond.L.Lock()
	defer cond.L.Unlock()
	cond.Wait()
	log.Printf("======condStop====== stopped")
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{} //общая WaitGroup
	cond := &sync.Cond{
		L: &sync.Mutex{},
	}
	stopCh := make(chan struct{})
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	log.Println("main: start all goroutines") // запускаем группы воркеров

	wg.Add(5)
	go selfStop(wg)
	go contextStop(ctxTimeout, wg)
	go chanCloseStop(stopCh, wg)
	go condStop(cond, wg)
	var wgStop sync.WaitGroup
	wgStop.Add(1)
	go waitgroupStop(&wgStop, wg)
	
	time.Sleep(4*time.Second) // имитация работы

	close(stopCh) // завершение воркеров chanCloseStop
	wgStop.Done() // завершение waitgroupStop
	cond.L.Lock()
	cond.Broadcast() // завершение condStop
	cond.L.Unlock()
	wg.Wait() // ждем, когда воркеры закончат выполнение
	log.Printf("main: all goroutines are stopped")
}
