package main

import "fmt"

func main() {
	ints := []int{10, 20, 30, 40, 50}
	ints[1] = 100

	fmt.Println(ints)

	ints = append(ints, 10, 40, 50)

	fmt.Println(ints)

}
