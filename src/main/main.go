package main

import (
	"fmt"
	"choose"
)

func main() {
	data := []int{6,0,2,0,1,3,4,6,1,3,2}
	fmt.Println(choose.FindMinAndMax(data))
}