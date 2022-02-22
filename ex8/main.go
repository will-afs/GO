package main

import (
	"fmt"
	"time"
)

func hello(ch chan string) {
	fmt.Print("Hello ")
	ch <- "Hello "
}

func world(ch1 chan string, ch2 chan bool) {
	_ = <-ch1
	fmt.Print("World\n")
	ch2 <- true
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan bool)
	go hello(ch1)
	go world(ch1, ch2)
	_ = <-ch2
	close(ch1)
	close(ch2)
	time.Sleep(1)
}
