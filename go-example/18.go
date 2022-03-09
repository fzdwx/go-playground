package page

import (
	"fmt"
	"regexp"
)

func main() {
	qwe, _ := regexp.Compile("qwe")
	fmt.Println("123")

	match := qwe.Match([]byte("qqqqqqqqqqqqwe"))
	fmt.Println(match)
}
