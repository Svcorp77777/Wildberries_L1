package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var number, mask, bit int64

	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &number)

	fmt.Printf("Введите какой i-й бит вы хотите поменять в числе %d на 1 или 0. ", number)
	fmt.Scanf("%d", &bit)

	mask = (1 << bit)

	fmt.Println("При изменении i-й бита на 1 число = ", number|mask) //or

	fmt.Println("При изменении i-й бита на 0 чсло =", number&^mask) // and

}
