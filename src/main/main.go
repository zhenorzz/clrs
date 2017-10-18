package main

import (
	"fmt"
	"matrix"
)

func main() {
	fmt.Println(matrix.SquareMatrixMultiply([][]int{{1,2,3},{2,4,1}}, [][]int{{1,2,3},{2,4,1}}))
}