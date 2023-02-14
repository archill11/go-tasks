package main

// Какой самый эффективный способ конкатенации строк?

import (
	"strings"
)

var params = []string{"string-1", "string-2", "string-3", "string-4"}

func main() {
	println(cCopy(params))
	println(cString(params))
}

func cCopy(args []string) string { //<- на тесте видно что это самый эффективный способ
	lenB := make([]byte, len(args)*len(args[0]))
	offset := 0
	for _, arg := range args {
		// fmt.Println(i, arg, offset, lenB[offset:], lenB)
		offset += copy(lenB[offset:], arg)
	}
	return string(lenB)
}

func cString(args []string) string {
	builder := strings.Builder{}
	for _, arg := range args {
		builder.WriteString(arg)
	}
	return builder.String()
}

