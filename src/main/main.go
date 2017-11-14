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
	root.Insert(10)
	root.Insert(31)
	root.Insert(25)
	root.Insert(13)
	t := root.Search(18)
	fmt.Println(t.Predecessor())
}