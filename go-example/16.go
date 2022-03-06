package main

import (
	"container/list"
	"fmt"
)

func main() {

	from := []int{1, 2, 3}
	to := make([]int, 10)

	_ = copy(to, from)

	fmt.Println(to)

	var l list.List

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	fmt.Println(l)
}
