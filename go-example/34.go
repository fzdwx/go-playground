package main

import (
	"fmt"
	"go-playground/semaphore"
)

func main() {

	ch := make(chan int)

	sem := semaphore.Create()

	go sendTo(ch)

	go readTo(ch, sem)

	sem.Wait(1)
	fmt.Println("main: done")
}

func readTo(ch chan int, sem semaphore.Semaphore) {

	for i := range ch {
		fmt.Println("receive:", i)
	}

	sem.Signal()
}

func sendTo(ch chan int) {
	for i := 0; i < 9; i++ {
		ch <- i * 10
	}
	close(ch)
}
