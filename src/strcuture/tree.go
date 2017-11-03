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

func (T *Tree) Delete(k *Tree) {

}

func (T *Tree) Search(x int) *Tree {
	p := T
	for {
		if p.Key == x {
			return p
		} else if p.Left != nil {
			p = p.Left
		} else if p.Right != nil {
			p = p.Right
		} else {
			return nil
		}
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

func (T Tree) String() string {
	return fmt.Sprintf("%d", T.Key)
}