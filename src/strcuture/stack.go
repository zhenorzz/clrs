package strcuture

type Stack struct {
	Top int
	Length int
	Key []int
}
//var stack = strcuture.Stack{
//	Top: 0,
//	Key: make([]int,9),
//	Length:9}
//fmt.Println(stack.StackEmpty())
//stack.Push(4)
//stack.Push(1)
//stack.Push(3)
//x1 := stack.Pop()
//stack.Push(8)
//x2 := stack.Pop()
//fmt.Println(x1, x2, stack)
func (S *Stack) StackEmpty() bool {
	if S.Top == 0 {
		return true
	} else {
		return false
	}
}

func (S *Stack) Push(x int) {
	S.Key[S.Top] = x
	S.Top++
}

func (S *Stack) Pop() int {
	S.Top--
	x := S.Key[S.Top]
	S.Key[S.Top] = 0
	return x
}