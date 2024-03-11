package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	var chek map[string]string = make(map[string]string)
	array := []string {"cat", "cat", "dog", "cat", "tree"}
	res := []string {}

	for _, value := range array {
		if _, exist := chek[value]; !exist {
			res = append(res, value)
			chek[value] = ""
		}
	}

	fmt.Println(res)
}