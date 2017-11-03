package strcuture

import "fmt"

type Tree struct {
	Key int
	Root *Tree
	Left *Tree
	Right *Tree
}

func (T *Tree) Insert(x int) {
	if x < T.Key && T.Left != nil {
		T.Left.Insert(x)
	} else if T.Right != nil {
		T.Right.Insert(x)
	}
	if x < T.Key && T.Left == nil {
		Node := Tree{Key:x,Left:nil,Right:nil,Root:T}
		T.Left = &Node
	} else if x >= T.Key && T.Right == nil {
		Node := Tree{Key:x,Left:nil,Right:nil,Root:T}
		T.Right = &Node
	}
}

func (T *Tree) Traverse() {
	fmt.Println(T)
	if T.Left != nil {
		T.Left.Traverse()
	}
	if T.Right != nil {
		T.Right.Traverse()
	}
}

func (T *Tree) Search(x int) *Tree {
	p := T
	for p != nil {
		if p.Key == x {
			return p
		} else if x < p.Key {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return nil
}

func (T *Tree) Delete(k *Tree) {
	if k.Root != nil {
		k.Root.Left = k.Left
	}
}

func (T Tree) String() string {
	return fmt.Sprintf("%d", T.Key)
}