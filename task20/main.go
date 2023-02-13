package main

// Задание 20
// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

// reverseWords меняет порядок слов в строке на противоположный.
func reverseWords(s string) string {
	words := strings.Fields(s)
	fmt.Println(words)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "Лёша на полке клопа нашёл"
	fmt.Println(reverseWords(s))
}
