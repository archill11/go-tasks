package main

// Задание 1
// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов
// в структуре Action от родительской структуры Human (аналог наследования).

import (
	"fmt"
	t "myapp/task01/structs"
)

func main() {
	// определяем человека
	human := t.NewHuman("Тест Тестов", 35)

	// добавляем человеку цель
	human.AddGoal("learn chinese")

	// определяем действие
	action := t.NewAction(human, "сходить на фитнес")

	// метод AddGoal "наследуется" структурой Action,
	// поэтому цель можно задать и через реализацию этой структуры:
	action.AddGoal("learn how to edit movies")
	// либо так:
	action.Human.AddGoal("bank a million yen")


	// определим Тесту новое действие:
	action1 := t.NewAction(human, "писать чистый код")
	// цели Теста доступны и через это действие:
	goals := action1.GetGoals()

	fmt.Println(goals)
}
