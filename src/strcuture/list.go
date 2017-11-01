package strcuture

type Node struct {
	Prev *Node
	Next *Node
	Key int
}
type List struct {
	Nil *Node
}

//var node = strcuture.Node{}
//var L = strcuture.List{Nil: &node}
//L.Nil.Prev = L.Nil
//L.Nil.Next = L.Nil
//L.Insert(&strcuture.Node{Key:5})
//L.Insert(&strcuture.Node{Key:3})
//wait := L.Search(3)
//L.Delete(wait)
//fmt.Println(L.Search(3))

func (L *List) Search(k int) *Node {
	x := L.Nil.Next
	L.Nil.Key = k
	for x.Key != k{
		x = x.Next
	}
	return x
}

func (L *List) Insert(x *Node) {
	x.Next = L.Nil.Next
	L.Nil.Next.Prev = x
	L.Nil.Next = x
	x.Prev = L.Nil
}

func (L *List) Delete(x *Node) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
}