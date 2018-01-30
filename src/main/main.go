package main

import (
	"mysort"
	"fmt"
)

func main() {
	data := []int{8, 19, 9, 5, 8, 8, 7, 4, 21, 2, 6, 11}
	mysort.MergeSort(data, 0, 10)
	fmt.Println(data)
}
