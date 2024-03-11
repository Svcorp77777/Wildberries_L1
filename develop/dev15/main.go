package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// Негативные: Присваивая переменной justString переменную v[:100] происходит обрезание строки но фактически в памяти также будет
// хранится вся созданная строка

// Исправление: Выполняя данное действие justString = string([]byte(v)) создается новая строка которя хранит только нужный фрагмент строки

var justString string

func createHugeString(num int) string {
	return strings.Repeat("go", num)
}

func someFunc() {
	v := createHugeString(1 << 10)

	v = v[:100]

	justString = string([]byte(v))

	v = ""

	fmt.Println(justString)
}

func main() {
	someFunc()
}
