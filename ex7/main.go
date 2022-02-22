package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %v: Started\n", id)
	fmt.Print("\nHello \n")
	time.Sleep(time.Second)
	fmt.Printf("Worker %v: Finished\n", id)
}

func world(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %v: Started\n", id)
	fmt.Print("\nWorld \n")
	time.Sleep(time.Second)
	fmt.Printf("Worker %v: Finished\n", id)
}

func main() {
	var wg sync.WaitGroup
	debut := time.Now()
	id := 1
	fmt.Println("Main: Starting worker", id)
	wg.Add(1)
	go hello(&wg, id)
	id = 2
	fmt.Println("Main: Starting worker", id)
	go world(&wg, id)
	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fin := time.Now()
	fmt.Printf("Main: Completed in %s\n", fin.Sub(debut))
	time.Sleep(1)
}
