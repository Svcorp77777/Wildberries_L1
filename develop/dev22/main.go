package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := big.NewInt(5000000)
	b := big.NewInt(3000000)

	multiplication := new(big.Int)
	multiplication.Mul(a, b)
	fmt.Println("Результат умножения = ", multiplication)

	division := new(big.Int)
	division.Div(a, b)
	fmt.Println("Результат деления = ", division)

	summation := new(big.Int)
	summation.Add(a, b)
	fmt.Println("Результат сложение = ", summation)

	subtraction := new(big.Int)
	subtraction.Sub(a, b)
	fmt.Println("Результат вычитания = ", subtraction)
}
