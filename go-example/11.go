package page

import (
	"fmt"
	"io"
	"log"
)

func main() {
	test2("1", "2", "3")

	func12("go")
}

func test2(s ...string) {
	for i := range s {
		fmt.Println(i)
	}
}

func func12(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
