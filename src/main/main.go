package main

import (
	"strcuture"
	"fmt"
)

func main() {
	var root = strcuture.Tree{Key:18,Root:nil,Left:nil,Right:nil}
	root.Insert(12)
	root.Insert(11)
	root.Insert(21)
	root.Traverse()
	root.Traverse()
	fmt.Println(root.Search(18))
}