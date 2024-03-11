package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	number1 := 6
	number2 := 2

	number1 = number1 + number2
	number2 = number1 - number2
	number1 = number1 - number2

	fmt.Printf("Число 1 := %d  Число 2 := %d\n", number1, number2)
}
