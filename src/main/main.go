package main

import (
	"fmt"
	"lru"
)

func main() {
	obj := lru.Constructor(3)
	param := obj.Get(2)
	obj.Put(1,1)
	fmt.Println(param)
}
