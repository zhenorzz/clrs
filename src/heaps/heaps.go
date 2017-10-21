package heaps

type Stack struct {
	Heaps []int
	Size int
	Length int
}

func Left(i int) int {
	return (i<<1) + 1
}

func Right(i int) int {
	return (i<<1) + 2
}
//var stack = heaps.Stack{[]int{27,17,14,6,13,10,1,5,7,12}, 10, 10}
//stack.MaxHeapify(3)
func (stack *Stack)MaxHeapify(i int) {
	left := Left(i)
	largest := i
	if stack.Size > left && stack.Heaps[left] > stack.Heaps[i] {
		largest = left
	}
	right := Right(i)
	if stack.Size > right && stack.Heaps[right] > stack.Heaps[largest] {
		largest = right
	}
	if largest != i {
		temp := stack.Heaps[i]
		stack.Heaps[i] = stack.Heaps[largest]
		stack.Heaps[largest] = temp
		stack.MaxHeapify(largest)
	}
}

//var stack = heaps.Stack{[]int{27,17,14,6,13,10,1,5,7,12}, 10, 10}
//stack.MaxHeapifyLoop(3)
func (stack *Stack)MaxHeapifyLoop(i int) {
	largest := i
	for  {
		left := Left(largest)
		if stack.Size > left && stack.Heaps[left] > stack.Heaps[largest] {
			largest = left
		}
		right := Right(largest)
		if stack.Size > right && stack.Heaps[right] > stack.Heaps[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		temp := stack.Heaps[i]
		stack.Heaps[i] = stack.Heaps[largest]
		stack.Heaps[largest] = temp
		i = largest
	}
}

func (stack *Stack)MinHeapify(i int) {
	left := Left(i)
	largest := i
	if stack.Size > left && stack.Heaps[left] < stack.Heaps[i] {
		largest = left
	}
	right := Right(i)
	if stack.Size > right && stack.Heaps[right] < stack.Heaps[largest] {
		largest = right
	}
	if largest != i {
		temp := stack.Heaps[i]
		stack.Heaps[i] = stack.Heaps[largest]
		stack.Heaps[largest] = temp
		stack.MinHeapify(largest)
	}
}

//var stack = heaps.Stack{[]int{5,3,17,10,84,19,6,22,9}, 9, 9}
//stack.BuildMaxHeap()
func (stack *Stack)BuildMaxHeap() {
	for i := stack.Length>>1 - 1; i >= 0 ; i-- {
		stack.MaxHeapify(i)
	}
}
//var stack = heaps.Stack{[]int{5,3,17,10,84,19,6,22,9}, 9, 9}
//stack.HeapSort()
func (stack *Stack)HeapSort() {
	stack.BuildMaxHeap()
	for i := stack.Size-1; i >= 1; i-- {
		stack.Heaps[0], stack.Heaps[i] = stack.Heaps[i], stack.Heaps[0]
		stack.Size--
		stack.MaxHeapify(0)
	}
}