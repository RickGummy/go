package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 0, 1

	return func() int {
		result := first
		first, second = second, first+second
		return result
	}
}

func main() {
	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}

// package main
//
// import (
// 	"fmt"
// )
//
// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }
//
// func main() {
// 	pos, neg := adder(), adder()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }
