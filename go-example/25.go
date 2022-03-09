package page

import (
	"bufio"
	"os"
)

func main() {

	file, err := os.OpenFile("qwe.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		w.WriteString("Hello World\n")
	}
	w.Flush()
}
