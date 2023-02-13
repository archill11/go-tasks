package main

// Задание 14
// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

import "fmt"

// printType определяет тип аргумента
func printType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("==== int ====")
		return
	case string:
		fmt.Println("==== string ====")
		return
	case bool:
		fmt.Println("==== bool ====")
		return
	case chan string:
		fmt.Println("==== channel ====")
		return
	}
}

func main() {
	xs := []interface{}{
		123456789,
		"this is string",
		true,
		make(chan string),
	}
	for _, x := range xs {
		printType(x)
	}
}
