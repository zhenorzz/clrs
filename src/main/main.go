package main

import (
	"fmt"
	"strcuture"
)

func main() {
	var stack = strcuture.Queue{
		Head: 0,
		Tail: 0,
		Key: make([]int,9),
		Length:9}
	stack.Enqueue(4)
	stack.Enqueue(1)
	stack.Enqueue(3)
	x1 := stack.Dequeue()
	stack.Enqueue(8)
	x2 := stack.Dequeue()
	fmt.Println(x1, x2, stack)
}