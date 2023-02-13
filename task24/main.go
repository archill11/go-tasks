package main

// Задание 24
// Разработать программу нахождения расстояния между двумя точками, которые представлены
// в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

import (
	"fmt"
	"math"
)

// Point - точка на плоскости.
type Point struct {
	x, y float64
}

// NewPoint конструктор
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance возвращает расстояние от данной точки до родительской.
func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt( // по тереме пифагора
		math.Pow((p2.x - p.x), 2) + math.Pow((p2.y - p.y), 2),
	)
}

func main() {
	p1 := NewPoint(1, 0)
	p2 := NewPoint(1, 3)
	fmt.Println(p1.Distance(p2))
}
