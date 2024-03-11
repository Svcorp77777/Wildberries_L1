package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	wg sync.WaitGroup
	counter int64
}

func (c *Counter) plus() {
	atomic.AddInt64(&c.counter, 1)
	c.wg.Done()
}

func main() {
	c := Counter{}

	

	for i := 0; i < 1000; i++ {
		c.wg.Add(1)
		
		go c.plus()
	}

	

	c.wg.Wait()
	fmt.Println(c.counter)

}
