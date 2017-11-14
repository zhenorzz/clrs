package strcuture

import "fmt"

type Tree struct {
	Key int
	Root *Tree
	Left *Tree
	Right *Tree
}
//var root = strcuture.Tree{Key:18,Root:nil,Left:nil,Right:nil}
//root.Insert(12)
//root.Insert(11)
//root.Insert(21)
//root.Traverse()
//node := root.Search(18)
//root.Delete(node)
//fmt.Println(root.Search(21))
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
	for p != nil && p.Key != x {
		if x < p.Key {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return p
}

//二叉搜索树的结点删除比插入较为复杂，总体来说，结点的删除可归结为三种情况：
//1、 如果结点z没有孩子节点，那么只需简单地将其删除，并修改父节点，用NIL来替换z；
//2、 如果结点z只有一个孩子，那么将这个孩子节点提升到z的位置，并修改z的父节点，用z的孩子替换z；
//3、 如果结点z有2个孩子，那么查找z的后继y，此外后继一定在z的右子树中，然后让y替换z。
func (T *Tree) Delete(k *Tree) {
	if k.Left == nil && k.Right == nil {
		if k.Root.Left == k {
			k.Root.Left = nil
		} else {
			k.Root.Right = nil
		}
	} else if k.Left != nil && k.Right == nil {
		if k.Root.Left == k {
			k.Root.Left = k.Left
			k.Left.Root = k.Root.Left
		} else {
			k.Root.Right = k.Left
			k.Left.Root = k.Root.Right
		}
	} else if k.Right != nil && k.Left == nil{
		if k.Root.Left == k {
			k.Root.Left = k.Right
			k.Right.Root = k.Root.Left
		} else {
			k.Root.Right = k.Right
			k.Right.Root = k.Root.Right
		}
	} else {
		if k.Right.Left == nil {
			k.Key = k.Right.Key
			k.Right = k.Right.Right
		} else {
			p := k.Right.Left
			for p.Left != nil {
				p = p.Left
			}
			if p.Right != nil {
				p.Right.Root = p.Root
				p.Root.Left = p.Right
			}
			k.Key = p.Key
		}
	}
}

func (T *Tree) Minimum() *Tree {
	p := T
	for p.Left != nil {
		p = p.Left
	}
	return p
}

func (T *Tree) Maximum() *Tree {
	p := T
	for p.Right != nil {
		p = p.Right
	}
	return p
}

func (T *Tree) Successor() *Tree {
	p := T
	if p.Right != nil {
		return p.Right.Minimum()
	}
	y := p.Root
	for y != nil && p == y.Right {
		p = y
		y = y.Root
	}
	return y
}

func (T *Tree) Predecessor() *Tree {
	p := T
	if p.Left != nil {
		return p.Left.Maximum()
	}
	y := p.Root
	for y != nil && p == y.Left {
		p = y
		y = y.Root
	}
	return y
}

//func (T Tree) String() string {
//	return fmt.Sprintf("%d", T.Key)
//}