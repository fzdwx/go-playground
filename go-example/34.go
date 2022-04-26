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
	// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md#1427-%E4%BF%A1%E5%8F%B7%E9%87%8F%E6%A8%A1%E5%BC%8F
	// 习惯用法：通道工厂模式
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
