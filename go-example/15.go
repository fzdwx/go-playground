package page

import (
	"bytes"
	"fmt"
)

func main() {

	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
		fmt.Println(item)
	}
	fmt.Println(items)

	s := []int{10, 20, 30, 40, 50}

	qwe := s[1:1]
	fmt.Println(qwe)
}

func ttttt() {
	i := *new([10]int)
	for i2 := range i {
		i[i2] = i2
	}
	s := i[2:4]
	fmt.Println(s)

	fmt.Println(i)

	var buffer bytes.Buffer

	buffer.WriteString("Hello")
	fmt.Println(buffer.String())
	b2 := bytes.Buffer{}
	b2.ReadFrom(&buffer)

	b2.WriteString(" world")

	fmt.Println(b2.String())
}

func fibArray(n int, arr []int) []int {
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

func ttt() {
	var arr [5]int

	for i := range arr {
		arr[i] = i * 2
	}

	for i := range arr {
		fmt.Println(arr[i])
	}
}
