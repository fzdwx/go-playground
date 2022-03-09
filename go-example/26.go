package main

import (
	"flag"
	"fmt"
)

var countFlag = flag.Int("c", 0, "print the count")

func main() {
	flag.PrintDefaults()
	flag.Parse()

	fmt.Println(*countFlag)
}
