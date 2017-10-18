package main

import (
	"fmt"
	"matrix"
)

func main() {
	fmt.Println(matrix.Strassen([][]int{{1,2},{3,4}}, [][]int{{1,2},{3,4}},2))
}