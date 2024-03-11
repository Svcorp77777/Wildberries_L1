package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {

	array := []int{2, 4, 6, 8, 10}
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
		for val := range ch {
			fmt.Println(val * val)
		}
		
		wg.Done()
	}()

	wg.Wait()
}
