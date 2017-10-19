package main

import (
	"fmt"
	"matrix"
)

func main() {
	monge := [][]int{{37,23,22,32},{21,6,7,10},{53,34,30,31},{32,13,9,6},{43,21,15,8}}
	pos := make(map[int]int)
	fmt.Println(matrix.DivideMonge(&monge,5,4, &pos))
}