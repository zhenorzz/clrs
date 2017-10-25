package main

import (
	"fmt"
	"quick"
)

func main() {
 	data := []int{2,2,2,2,2,2,2,2,2,2,2,2}
	quick.Hoare(data,0,11)
	fmt.Println(data)
}