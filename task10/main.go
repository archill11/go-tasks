package main

// Задание 10
// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

import "fmt"

// groupTemp группирует числа из переданного слайса в группы с шагом в 10 градусов.
func groupTemp(temperatures []float32) map[int][]float32 {
	groupsMap := make(map[int][]float32)
	for _, t := range temperatures { // вычисляем идентификатор группы
		gId := int(t/10) * 10
		v, ok := groupsMap[gId]
		if ok { // если в мапе есть такой ключ аппендим туда 
			groupsMap[gId] = append(v, t)
		}else{ // иначе создаем такой ключ и кладем значение
			groupsMap[gId] = []float32{t}
		}
	}
	return groupsMap
}

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 115.5, 24.5, -21.0, 32.5}
	groups := groupTemp(temperatures)
	for id, g := range groups {
		fmt.Printf("Группа %d: ", id)
		for _, t := range g {
			fmt.Printf(" %.1f,", t)
		}
		fmt.Println()
	}
}
