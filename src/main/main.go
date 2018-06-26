package main

import (
	"fmt"
	"trie"
)

func main() {
	t := trie.New()
	t.Put("go","hello world")
	value := t.Get("go")
	fmt.Println(value)
	t.Delete("go")
	fmt.Println(t.Get("go"))
}
