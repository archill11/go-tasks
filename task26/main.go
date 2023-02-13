package main

// Задание 26
// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

import (
	"fmt"
	"strings"
)

// uniqueChars проверяет, являются ли все символы в данной строке уникальными
func uniqueChars(s string) bool {
	repMap := make(map[rune]bool)
	for _, r := range strings.ToLower(s) {
		if _, ok := repMap[r]; ok { // если в мапе уже есть такой символ - возвращаем false
			return false
		}
		repMap[r] = true
	}
	return true
}

func main() {
	arr := []string{"abcd", "abCdefAaf", "aAbcd", "abCdefA", "abCdef"}
	for _, s := range arr {
		fmt.Printf("%s: %v\n", s, uniqueChars(s))
	}
}
