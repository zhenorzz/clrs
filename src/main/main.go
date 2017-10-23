package main

import (
	"fmt"
	"heaps"
)

func main() {
	var stack = heaps.Stack{[]int{5,13,2,25,7,17,20,8,4}, 9, 9}
	stack.BuildMaxHeap()
	stack.HeapDelete(2)
	fmt.Println(stack)
}