package main

import (
	"fmt"
	"quick"
)

func main() {
 	data := []int{13,19,9,5,12,8,7,4,21,2,6,11}
	quick.Sort(data,0,11)
	fmt.Println(data)
}