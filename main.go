package main

import (
	"fmt"
	"go-playground/tuple"
)

const filePath = "text.txt"

func main() {

	t := tuple.Of2(1, 2)

	t2 := tuple.Of3(1, 2, "!23")
	fmt.Println(t.V1)
	fmt.Println(t2)

}
