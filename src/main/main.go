package main

import (
	"fmt"
	"heaps"
)

func main() {
	var stack = heaps.Stack{[]int{5,3,17,10,84,19,6,22,9}, 9, 9}
	stack.HeapSort()
	fmt.Println(stack)
}