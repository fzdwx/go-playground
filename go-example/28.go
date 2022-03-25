package main

import (
	"errors"
)

var myerr = errors.New("this is my err")

func myadd(a, b int) error {
	return myerr
}

func main() {

	err := myadd(1, 2)
	if err != nil {
		panic(err)
	}
}
