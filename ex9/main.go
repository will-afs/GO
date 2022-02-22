package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_door(wall []bool) {
	n_doors := len(wall)
	rand.Seed(time.Now().UnixNano())
	open_door_id := rand.Intn(n_doors)
	wall[open_door_id] = true
}

func look_for_door(wall []bool, start_index int, end_index int, ch chan int) {
	for i := start_index; i < end_index; i++ {
		if wall[i] == true {
			ch <- i
		}
	}
}

func main() {
	ch := make(chan int)
	wall := make([]bool, 100, 100)
	generate_door(wall)
	step := len(wall) / int(20)
	for i := 0; i < len(wall); i += step {
		go look_for_door(wall, i, i+step, ch)
	}
	fmt.Printf("Open door ID : %d", <-ch)
}
