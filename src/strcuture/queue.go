package strcuture

type Queue struct {
	Head int
	Tail int
	Key []int
	Length int
}

//var stack = strcuture.Queue{
//	Head: 0,
//	Tail: 0,
//	Key: make([]int,9),
//	Length:9}
//stack.Enqueue(4)
//stack.Enqueue(1)
//stack.Enqueue(3)
//x1 := stack.Dequeue()
//stack.Enqueue(8)
//x2 := stack.Dequeue()
//fmt.Println(x1, x2, stack)

func (Q *Queue) Enqueue(x int) {
	if Q.Tail + 1 == Q.Head || (Q.Tail == Q.Length && Q.Head == 0) {

	}
	Q.Key[Q.Tail] = x
	if Q.Tail == Q.Length {
		Q.Tail = 1
	} else {
		Q.Tail++
	}
}

func (Q *Queue) Dequeue() int {
	if Q.Head == Q.Tail {

	}
	x := Q.Key[Q.Head]
	Q.Key[Q.Head] = 0
	if Q.Head == Q.Length {
		Q.Head = 1
	} else {
		Q.Head++
	}
	return x
}