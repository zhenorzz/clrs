package main

import (
	"fmt"
	"quick"
)

func main() {
 	data := []int{11,19,9,5,12,8,7,4,11,2,6,21}
	quick.Sort(data,0,11)
	fmt.Println(data)
}