package main

import "time"

func main() {

	chanA := make(chan int, 2)

	go func() {
		send(chanA)
	}()

	//go func() {
	consume(chanA)
	//}()

	time.Sleep(time.Second * 10)
}

func consume(a chan int) {
	for {
		select {
		case v := <-a:
			println("consume:", v)
		}
	}
}

func send(a chan int) {
	a <- 1
	a <- 2
	a <- 3
	a <- 4
	a <- 5
	a <- 6
}
