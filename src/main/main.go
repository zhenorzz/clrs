package main

import (
	"quick"
	"fmt"
)

func main() {
	data := []int{8,19,9,5,8,8,7,4,21,2,6,11}
	quick.Sort(data,0,11)
	fmt.Println(data)
}