package page

import "fmt"

func Multiply(x, y int, reply *int) {
	*reply = x * y
}
func main() {
	var reply int
	Multiply(10, 20, &reply)
	fmt.Println(reply)
}

// DEMO1:
func demo1() {
	fmt.Println(test3(func(i int, i2 int) int {
		return i + i2
	}))
}

func test3(add add) int {
	return add(1, 2)
}

type add func(int, int) int

// END
