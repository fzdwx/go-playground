package page

import (
	"fmt"
)

func main() {

	//page := Page{Title: "qwe.txt"}
	//err := load(&page)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(string(page.Body))
	//var mu sync.Mutex
	//mu.Lock()
	go func() {
		fmt.Println("helloworld")
		//mu.Unlock()
	}()
}
