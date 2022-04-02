package main

import "fmt"

func main() {

	ch := make(chan int)

	go getSum(1, 10, ch)

	fmt.Println(<-ch)
}

func getSum(i2 int, i22 int, ch chan int) {
	ch <- i2 + i22
}
