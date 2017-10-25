package main

import (
	"fmt"
	"liner"
)

func main() {
 	data := []int{329,457,657,839,436,720,355}
	fmt.Println(liner.RadixSort(data,10,3))
}