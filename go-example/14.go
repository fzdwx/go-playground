package main

import "fmt"

func main() {

	fmt.Println(f()) // 2

	// var x int
	//
	// 这时候的add 就等于如下函数
	// func(delta int) int {
	//		x += delta
	//		return x
	//	}
	add := adder()

	fmt.Println(add(1))
	fmt.Println(add(1))
	fmt.Println(add(1))
	fmt.Println("=============================================================")
	fmt.Println(adder()(1))
	fmt.Println(adder()(1))
	fmt.Println(adder()(1))
	// 6.10

}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
