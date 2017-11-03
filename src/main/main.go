package main

import (
	"strcuture"
)

func main() {
	var root = strcuture.Tree{Key:18,Root:nil,Left:nil,Right:nil}
	root.Insert(12)
	root.Insert(11)
	root.Insert(21)
	root.Traverse()
}