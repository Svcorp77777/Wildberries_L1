package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	line := "snow dog sun"

	res := strings.Split(line, " ")

	left := 0
	right := len(res) -1

	for left <= right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}

	result := strings.Join(res, " ")

	fmt.Println(result)
}