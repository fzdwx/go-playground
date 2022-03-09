package main

import "fmt"

func main() {

	ple := PeoPle{}
	run(ple)
}

type PeoPle struct {
	name string
}

func (this PeoPle) Accept(val interface{}) {
	fmt.Println(val)
}

func run(c Consumer) {
	for i := 0; i < 10; i++ {
		c.Accept(i)
	}
}

type Consumer interface {
	Accept(val interface{})
}
