package main

import (
	"fmt"
)

const filePath = "text.txt"

func main() {

	t := of2(1, 2)

	t2 := of3(1, 2, "!23")
	fmt.Println(t.t1)
	fmt.Println(t2)
}

type Tuple2[T1 any, T2 any] struct {
	t1 T1
	t2 T2
}

type Tuple3[T1 any, T2 any, T3 any] struct {
	t1 T1
	t2 T2
	t3 T3
}

func of2[T1 any, T2 any](t1 T1, t2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{
		t1,
		t2,
	}
}

func of3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{
		t1,
		t2,
		t3,
	}
}
