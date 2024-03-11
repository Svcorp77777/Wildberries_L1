package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	ch := make(chan int)
	var second int

	fmt.Print("Введите время до завершения программы в секундах: ")
	fmt.Scanf("%d", &second)

	go func() {
		for i := 0; ; i++ {
			ch <- i

			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	time.Sleep(time.Duration(second) * time.Second)
}
