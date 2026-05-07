package main

import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// var (
// 	v1 = Vertex{1, 2}
// 	v2 = Vertex{X: 1}
// 	v3 = Vertex{}
// 	p  = &Vertex{1, 2}
// )

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}
