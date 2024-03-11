package main

import (
	"fmt"
	"sync"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	array := []int{2, 4, 6, 8, 10}
	var res, square int

	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for _, value := range array {
			ch <- value
		}

		wg.Done()

		close(ch)
	}()

	go func() {
		for value := range ch {
			square = value * value
			res += square
		}
		fmt.Println(res)

		wg.Done()
	}()

	wg.Wait()
}
