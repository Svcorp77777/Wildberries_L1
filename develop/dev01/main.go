package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

type Human struct {
	name string
	age  uint
}

type Action struct {
	Human
}

func (h *Human) ageUp() {
	h.age++
}

func NewAction() *Action {
	return &Action{
		Human: Human{},
	}
}

func main() {

}
