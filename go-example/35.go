package main

import "fmt"

func main() {
	ch := make(chan int)
	defer close(ch)

	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			}
		}
	}()
	for {
		ch <- 1
	}
}
