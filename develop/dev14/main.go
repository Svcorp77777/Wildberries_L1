package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func main() {
	ch := make(chan int)
	d := []interface{} {5, "sss", true, 5.4, ch}

	for _, value := range d {
		switch myType := value.(type) {
		case int:
			fmt.Println("Тип переменной int")
		case string:
			fmt.Println("Тип переменной string")
		case bool:
			fmt.Println("Тип переменной bool")
		default:
			fmt.Printf("Тип переменной %T\n", myType)
		}

	}

}
