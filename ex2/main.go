package main

import (
	"fmt"
	//"os"
	//"strconv"
)

func collatz_conjecture(n uint64) (rets uint64){
	// var nf float32 = float32(n)
	var res uint64
	if (n % 2) == 0 {
		res = n/2
	} else {
		res = 3*n+1
	}
	return res
}

func main() {
	// n, err := strconv.Atoi(os.Args[0])
	var n uint64 = 23
	last_res := n
	for i := n; i>=0 ; i--{
		last_res = collatz_conjecture(last_res)
		fmt.Println(last_res)
	}
}