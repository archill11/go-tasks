package structs

// Action наследует методы структуры Human.
type Action struct {
	Human
	ActionType string
}

// NewAction конструктор
func NewAction(hum Human, at string) Action {
	a := Action {
		Human: hum,
		ActionType: at,
	}
	return a
}