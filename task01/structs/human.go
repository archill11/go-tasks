package structs

import (
	"sync"
)

// Human представляет человека с его целями.
type Human struct {
	Name    string
	Age     int
	Mu *sync.RWMutex // для потокобезопасного использования map.
	Goals map[string]struct{} // goals - цели реализованы в виде map для бысроты поиска и простоты удаления.
}

// NewHuman конструктор
func NewHuman(name string, age int) Human {
	h:= Human{
		Name: name,
		Age: age,
		Mu:      &sync.RWMutex{},
		Goals:   make(map[string]struct{}),
	}
	return h
}

// AddGoal добавляет цель к списку целей человека.
func (h *Human) AddGoal(goal string) {
	h.Mu.Lock()
	defer h.Mu.Unlock()
	h.Goals[goal] = struct{}{}
}

// GetGoals - получить все цели в виде слайса строк.
func (h *Human) GetGoals() []string {
	h.Mu.RLock()
	defer h.Mu.RUnlock()
	result := make([]string, 0, len(h.Goals))
	for goal := range h.Goals {
		result = append(result, goal)
	}
	return result
}
