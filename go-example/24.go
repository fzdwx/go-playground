package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//var a, b string
	//fmt.Scanln(&a, &b)
	//
	//fmt.Println("a: " + a)
	//fmt.Println("b: " + b)
	open, err := os.Open("../text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(open)
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
}
