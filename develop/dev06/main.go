package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	exit := make(chan os.Signal, 1)
	stop := make(chan bool)

	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go goStopChan(stop)
	go goStopTime()
	go goStopOs(exit)

	time.Sleep(3 * time.Second)

	stop <- true

	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	
	time.Sleep(20 * time.Second)
	fmt.Println("Программа завершила работу")

}

func goStopOs(exit <-chan os.Signal) {
	for {
		select {
		case <-exit:
			fmt.Println("Горутина по сигналу остановлена")
			return
		case <- time.After(1 * time.Second):
			fmt.Println("Горутина по сигналу работает")
			

		}
	}
}

func goStopTime() {

	fmt.Println("Горутина по времени запущена")

	time.Sleep(2 * time.Second)

	fmt.Println("Горутина по времени завершена")

	
}

func goStopChan(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Горутина кнала остановлена")
			return
		default:
			fmt.Println("Горутина канала работает")
			time.Sleep(1 * time.Second)

		}
	}
}
