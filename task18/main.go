package main

// Задание 18
// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// счетчик с испотзованием atomic
type AtomicCounter struct {
	counter uint64
}

// Inc инкрементирует счётчик
func (c *AtomicCounter) Inc() {
	atomic.AddUint64(&c.counter, 1)
}

// счетчик с испотзованием Mutex
type MutexCounter struct {
	count int
	mu sync.Mutex
}

// Inc инкрементирует счётчик
func (c *MutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count+=1
}

func main() {
	// счетчик с испотзованием atomic
	var wg sync.WaitGroup
	c := &AtomicCounter{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 500; i++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.counter)
	fmt.Println("atomic Done")


	// счетчик с испотзованием Mutex
	cm := MutexCounter{}
	var wgm sync.WaitGroup
	for i := 0; i < 5; i++ {
		wgm.Add(1)
		go func() {
			defer wgm.Done()
			cm.Inc()
		}()
	}
	wgm.Wait()
	fmt.Println(cm.count)
	fmt.Println("Mutex Done")
}