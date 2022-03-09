package page

import "fmt"

//
// for init;condition;post{
//		// do something
//	}
//
func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}

}

func ex2() {
	str := "G"
	for i := 0; i < 25; i++ {
		fmt.Println(str)
		str += "G"
	}
}

func ex1() {
	i := 0

START:
	fmt.Printf("the counter is at %d\n ", i)
	i++
	if i < 15 {
		goto START

	}
}

func func2() {
	str := "GO is a beautiful language!"
	strLen := len(str)
	fmt.Printf("The length of str is: %d\n", strLen)
	for i := 0; i < strLen; i++ {
		fmt.Printf("Character on position %d is: %c\n", i, str[i])
	}

	fmt.Println("FFFFFFFFFFFFFFFFFFFFFFFFFFFF")

	str2 := "中国话"
	str2Len := len(str2)
	fmt.Printf("The length of str2 is: %d\n", str2Len)
	for i := 0; i < str2Len; i++ {
		fmt.Printf("Character on position %d is: %c\n", i, str2[i])
	}
}

func func1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}
