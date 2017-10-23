package main

import (
	"fmt"
	"heaps"
	"imath"
)

func main() {
	max := imath.MaxInt
	var young = heaps.Young{[][]int{{2,3,4,5},{8,9,12,14},{16,max,max,max},{max,max,max,max}}, 4, 4}
	fmt.Println(young.YoungFind(21))
}