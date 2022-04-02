package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("in main")

	go longWait()
	go shortWait()

	fmt.Println("about sleep in main()")
	time.Sleep(20 * time.Second)

	fmt.Println("at the end of main()")

}

func shortWait() {

	fmt.Printf("start shortWait() %v\n", time.Now().UnixNano())
	time.Sleep(2 * time.Second)

	fmt.Println("end shortWait()")
}

func longWait() {

	fmt.Printf("start longWait() %v\n", time.Now().UnixNano())
	time.Sleep(10 * time.Second)
	fmt.Println("end longWait()")
}
