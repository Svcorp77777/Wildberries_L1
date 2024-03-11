package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	var runes []rune
	line := "главрыба"
	runes = []rune(line)

	left := 0
	right := len(runes) -1
	
	for left <= right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	fmt.Println(string(runes))
	
}
