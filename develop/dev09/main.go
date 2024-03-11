package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала:
// в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	array := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg.Add(3)

	go func ()  {
		for _, value := range array {
			ch1 <- value
		}
		wg.Done()

		close(ch1)
	}()
	
	go func ()  {
		for i := range ch1 {
			i *= 2
			ch2 <- i

		}

		wg.Done()

		close(ch2)
	}()

	go func ()  {
		for j := range ch2 {
			fmt.Println(j)
		}

		wg.Done()
	}()


		wg.Wait()
}