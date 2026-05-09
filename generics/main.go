package main

import (
	"fmt"
)

type List[T any] struct {
	next *List[T]
	val  T
}

// add value to front of list
func (l *List[T]) Push(v T) *List[T] {
	return &List[T]{
		val:  v,
		next: l,
	}
}

// print list
func (l *List[T]) Print() {
	for curr := l; curr != nil; curr = curr.next {
		fmt.Println(curr.val)
	}
}

// count nodes
func (l *List[T]) Length() int {
	count := 0

	for curr := l; curr != nil; curr = curr.next {
		count++
	}

	return count
}

func main() {
	var list *List[int]

	list = list.Push(3)
	list = list.Push(2)
	list = list.Push(1)

	list.Print()

	fmt.Println("Length:", list.Length())
}
