package main

import (
	"fmt"
	"mysort"
)

func main() {
	fmt.Println(int(^uint(0)>>1))
	fmt.Println(mysort.MergeSort([]int{3,41,52,26,38,57,9,49},0,7))
}