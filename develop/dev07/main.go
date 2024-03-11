package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.

type Record struct {
	mu sync.Mutex
	d  map[int]string
}

func (r *Record) recordMap(key int, data string) {
	r.mu.Lock()
	r.d[key] = data
	r.mu.Unlock()
}

func main() {
	d := Record{d: make(map[int]string)}

	for i := 1; i <= 40; i++ {
		data := fmt.Sprintf("Запись: %d", i)
		go d.recordMap(i, data)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(d.d)
}
