package main

import (
	"flag"
	"fmt"
	"os"
)

// go run 27.go qwe.txt ../hello.txt
func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		cat(os.Stdin)
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.OpenFile(flag.Arg(i), os.O_RDONLY, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(f)
		f.Close()
	}
}

var i = 0

func cat(f *os.File) {
	const maxBufSize = 512
	var buf [maxBufSize]byte
	for {
		i++
		fmt.Println("i:", i)
		switch read, err := f.Read(buf[:]); true {
		case read < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case read == 0:
			return
		case read > 0:
			if nw, ew := os.Stdout.Write(buf[0:read]); nw != read {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}
		}
	}
}
