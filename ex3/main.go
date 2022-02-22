package main

import (
	"fmt"
)

type MyInt *int

func even(i *int) (rets bool){
	var res bool
	if (*i % 2) == 0 {
		res = true
	} else {
		res = false
	}
	return res
}

// func odd(i int) (rets bool){
// 	var res bool
// 	if even(i) == false{
// 		res = true
// 	} else {
// 		res = false
// 	}
// 	return res
// }

func odd(i MyInt) (rets bool){
	var res bool
	//var i_custom int = int(i)
	if even(i) == false{
		res = true
	} else {
		res = false
	}
	return res
}

func main() (){
	i := 10
	fmt.Println(even(&i))
	var my_i MyInt = &i
	fmt.Println(odd(my_i))
}

