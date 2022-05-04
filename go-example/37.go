package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(time.Second)

	fmt.Println(<-timer.C)
	for {
		t := <-ticker.C

		fmt.Println(t)
	}

}
