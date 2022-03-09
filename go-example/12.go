package main

import "fmt"

func main() {

	test333(10)

}

func fib(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	return
}

func fib2(n int) (res int, idx int) {
	idx = n
	if n <= 1 {
		res = 1
	} else {
		r, _ := fib2(n - 1)
		r2, _ := fib2(n - 2)
		res = r + r2
	}
	return
}

func test333(start int) {
	if start > 0 {
		fmt.Println(start)
		test333(start - 1)
	}
}
