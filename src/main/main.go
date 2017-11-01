package main

import (
	"fmt"
	"strcuture"
)

func main() {
	var node = strcuture.Node{}
	var L = strcuture.List{Nil: &node}
	L.Nil.Prev = L.Nil
	L.Nil.Next = L.Nil
	L.Insert(&strcuture.Node{Key:5})
	L.Insert(&strcuture.Node{Key:3})
	wait := L.Search(3)
	L.Delete(wait)
	fmt.Println(L.Search(3))
}