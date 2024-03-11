package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

type Pet interface {
	Say()
}

type Dog struct{}

func (d *Dog) AdapterSay() {
	fmt.Println("Гав Гав")
}

type PetAdapter struct {
	*Dog
}

func (p *PetAdapter) Say() {
	p.AdapterSay()
}

func NewAdapter(adapter *Dog) Pet {
	return &PetAdapter{adapter}
}

func main() {
	fmt.Println("Мой пёс говорит:")
	adapter := NewAdapter(&Dog{})
	adapter.Say()
}
