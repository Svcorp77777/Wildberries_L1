package main

import (
	"fmt"
)

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	var chek map[int]string = make(map[int]string)
	array1 := []int{5, 3, 2, 10, 6}
	array2 := []int{2, 1, 6, 4}
	intersection := []int{}

	for _, value := range array1 {
		chek[value] = ""
	}

	for _, value := range array2 {
		if _, exist := chek[value]; exist {
			intersection = append(intersection, value)
		}
	}

	fmt.Println(intersection)
}
