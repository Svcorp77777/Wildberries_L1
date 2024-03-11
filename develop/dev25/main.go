package main

import "time"

// Реализовать собственную функцию sleep.

func main() {

	sleep(5)

}

func sleep(t int) {
	<-time.After(time.Second * time.Duration(t))
}
