package main

// Задание 25
// Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

// sleepTimeAfter использует функцию time.After
func sleepTimeAfter(interval time.Duration) {
	<-time.After(interval)
}

func main() {
	interval := 2*time.Second
	sleepTimeAfter(interval)
	fmt.Println("end")
}
