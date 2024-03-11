package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из
// канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

func main() {
	ch := make(chan int)
	var n int

	var wg sync.WaitGroup

	fmt.Print("Введите количество воркеров: ")
	fmt.Scanf("%d", &n)

	wg.Add(1)

	go func ()  {
		for i := 0; ; i++ {
			ch <- i

			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < n; i++ {
		go func ()  {
			for {
				fmt.Printf("Значение: %d Воркер: %d\n",<- ch, i)
			}
		}()
	}

	wg.Wait()
}
