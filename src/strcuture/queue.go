package strcuture

type Queue struct {
	Head int
	Tail int
	Key []int
	Length int
}

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