package main

import (
	"fmt"
)

/*
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
*/

// var i, j int = 1, 2

func main() {
	const a int = 0
	var b float64
	var c bool
	var d string

	fmt.Printf("%v %v %v %q\n", a, b, c, d)

	var i int
	j := 2.0
	c = true
	python := false
	java := "no!"

	fmt.Println(i, j, c, python, java)
}
