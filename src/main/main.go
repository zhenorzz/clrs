package main

import (
	"fmt"
	"heaps"
)

func main() {
	var stack = heaps.Dstack{[]int{5,13,2,25,7,17,20,8,4}, 9, 9, 3}
	stack.BuildDMaxHeap()
	stack.DinsertKey(1, 12)
	fmt.Println(stack)
}