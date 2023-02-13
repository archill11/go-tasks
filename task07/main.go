package main

// Задание 7
// Реализовать конкурентную запись данных в map.

import (
	"fmt"
	"sync"
)

// Конкурентная запись в map возможна с использованием
//	- sync.Mutex
//	- sync.RWMutex
//	- sync.Map

type Container struct {
    mu       sync.RWMutex
    counters map[string]int
}

func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]+=1
}

func main() {
	// с помощью sync.RWMutex
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }
    var wg sync.WaitGroup
    doIncrement := func(name string, n int) {
        for i := 0; i < n; i++ {
            c.inc(name)
        }
        wg.Done()
    }
    wg.Add(3)
    go doIncrement("a", 10000)
    go doIncrement("a", 10000)
    go doIncrement("b", 10000)
    wg.Wait()
    fmt.Println(c.counters)

	// с помощью sync.Map
	var counters sync.Map
	var wgm sync.WaitGroup
    doIncrementSyncMap := func(name string, n int) {
        for i := 0; i <= n; i++ {
            counters.Store(name, i)
        }
        wgm.Done()
    }
    wgm.Add(2)
    go doIncrementSyncMap("a", 20000)
    go doIncrementSyncMap("b", 10000)
    wgm.Wait()
    fmt.Println(counters.Load("a"))
    fmt.Println(counters.Load("b"))
}
