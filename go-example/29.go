package main

import (
	fmt "fmt"
	"testing"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Printf("is the integer %d even? %v\n", i, Even(i))
	}
}

func Even(i int) bool {
	return i%2 == 0
}

func Odd(i int) bool {
	return i%2 != 0
}

func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log(" 10 must be even!")
		t.Fail()
	}
	if Even(7) {
		t.Log(" 7 is not even!")
		t.Fail()
	}

}

func TestOdd(t *testing.T) {
	if !Odd(11) {
		t.Log(" 11 must be odd!")
		t.Fail()
	}
	if Odd(10) {
		t.Log(" 10 is not odd!")
		t.Fail()
	}
}
