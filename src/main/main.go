package main

import (
	"fmt"
	"divide"
)

func main() {
	monge := [][]int{{37,23,22,32},{21,6,5,10},{53,34,30,31},{32,13,9,6},{43,21,15,8}}
	fmt.Println(divide.MongeDivide(monge,5,4))
}