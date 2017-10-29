package main

import (
	"fmt"
	"liner"
)

func main() {
	data := []int{6,0,2,0,1,3,4,6,1,3,2}
	liner.CountingSortInPlace(data,10)
	fmt.Println(data)
}