package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func main() {
	s1 := uniqueness("abcd")
	fmt.Println(s1)

	s2 := uniqueness("abCdefAaf")
	fmt.Println(s2)

	s3 := uniqueness("aabcd")
	fmt.Println(s3)
}

func uniqueness(str string) bool {
	var chek map[string]string = make(map[string]string)

	strLow := strings.ToLower(str)

	for _, value := range strLow {
		if _, exist := chek[string(value)]; exist {
			return false
		}

		chek[string(value)] = ""
	}

	return true
}
