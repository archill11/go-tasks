package main

import (
	"fmt"
)

// Задание 12
// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

// Идиоматический способ реализации set в Go — использование карты
type Set struct {
	s map[string]struct{}
}

// NewStringSet создаёт новое множество из переданного слайса.
func NewStringSet(v []string) Set {
	set := Set{
		s : make(map[string]struct{}),
	}
	for _, s := range v {
		set.s[s] = struct{}{}
	}
	return set
}

// Has возвращает true, если elem является членом множетсва.
func (s *Set) Has(elem string) bool {
	_, ok := s.s[elem]
	return ok
}

// Add добавляет во множество новый элемент.
func (s *Set) Add(elem string) {
	s.s[elem] = struct{}{}
}

// Delete удаляет элемент из множетсва.
func (s *Set) Delete(elem string) {
	delete(s.s, elem)
}

func main() {
	v := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewStringSet(v)
	fmt.Println(set)

	fmt.Println(set.Has("tree"))
	set.Delete("tree")
	set.Delete("tree")
	fmt.Println(set.Has("tree"))

	fmt.Println(set)
}
