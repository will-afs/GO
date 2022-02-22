package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func cat(filename string) {
	f, err := os.Open(filename)
	defer f.Close()
	check(err)
	buf := make([]byte, 15)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		check(err)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
}

func main() {
	cat("../ex4/main.go")
}
