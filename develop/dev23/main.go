package main

import (
	"fmt"
)

// Удалить i-ый элемент из слайса.

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	i := 3

	array[i] = array[len(array) - 1]

	array =	array[:len(array) - 1]



	fmt.Println("Слайс после удаления ", array)

}

