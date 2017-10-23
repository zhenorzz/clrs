package heaps

import (
	"errors"
	"imath"
)

type Stack struct {
	Heaps []int
	Size int
	Length int
}

func Parent(i int) int {
	return (i-1) >> 1
}
func Left(i int) int {
	return (i<<1) + 1
}

func Right(i int) int {
	return (i<<1) + 2
}

//var stack = heaps.Stack{[]int{27,17,14,6,13,10,1,5,7,12}, 10, 10}
//stack.MaxHeapify(3)
func (stack *Stack) MaxHeapify(i int) {
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
func (stack *Stack) MaxHeapifyLoop(i int) {
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

func (stack *Stack) MinHeapify(i int) {
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
func (stack *Stack) BuildMaxHeap() {
	for i := stack.Length>>1 - 1; i >= 0 ; i-- {
		stack.MaxHeapify(i)
	}
}
//var stack = heaps.Stack{[]int{5,3,17,10,84,19,6,22,9}, 9, 9}
//stack.HeapSort()
func (stack *Stack) HeapSort() {
	stack.BuildMaxHeap()
	for i := stack.Size-1; i >= 1; i-- {
		stack.Heaps[0], stack.Heaps[i] = stack.Heaps[i], stack.Heaps[0]
		stack.Size--
		stack.MaxHeapify(0)
	}
}

//返回最大值
func (stack *Stack) HeapMaximun() int {
	return stack.Heaps[0]
}

//去掉并返回stack中的最大元素
func (stack *Stack) HeapExtractMax() (int, error) {
	if stack.Size < 1 {
		return 0, errors.New("heap underflow")
	}
	max := stack.HeapMaximun()
	stack.Heaps[0] = stack.Heaps[stack.Size-1]
	stack.Size--
	stack.MaxHeapify(0)
	return max, nil
}

//替换元素
//var stack = heaps.Stack{[]int{5,13,2,25,7,17,20,8,4}, 9, 9}
//stack.BuildMaxHeap()
//fmt.Println(stack.HeapIncreaseKey(5,16))
func (stack *Stack) HeapIncreaseKey(i, key int) error {
	if key < stack.Heaps[i] {
		return errors.New("new key is smaller than current key")
	}
	stack.Heaps[i] = key
	for i > 0 &&stack.Heaps[Parent(i)] < stack.Heaps[i]  {
		stack.Heaps[Parent(i)], stack.Heaps[i] = stack.Heaps[i], stack.Heaps[Parent(i)]
		i = Parent(i)
	}
	return nil
}

//var stack = heaps.Stack{[]int{5,13,2,25,7,17,20,8,4}, 9, 9}
//stack.BuildMaxHeap()
//stack.MaxHeapInsert(1)
func (stack *Stack) MaxHeapInsert(key int) {
	stack.Size++
	stack.Length++
	stack.Heaps = append(stack.Heaps, imath.MinInt)
	stack.HeapIncreaseKey(stack.Size-1, key)
}

//var stack = heaps.Stack{[]int{5,13,2,25,7,17,20,8,4}, 9, 9}
//stack.BuildMaxHeap()
//stack.HeapDelete(2)
func (stack *Stack) HeapDelete(i int) bool {
	stack.HeapIncreaseKey(i, imath.MaxInt)
	stack.Size--
	stack.Heaps[0] = stack.Heaps[stack.Size]
	stack.MaxHeapify(0)
	return true
}

//6.5-9
//四个最小堆链表，每次HeapExtractMin 组合成新有有序链表

//D叉树
type Dstack struct {
	Heaps []int
	Size int
	Length int
	Dary int
}

func (stack *Dstack) Dparent(i int) int {
	return (i - 1) / stack.Dary
}

func (stack *Dstack) Dchild(k, i int) int {

	return (i + 1) * stack.Dary - k % stack.Dary
}

//返回叶子结点的开始位置
func (stack *Dstack) Dleaf(k, i int) int {
	return (i + 1) * stack.Dary - k % stack.Dary
}

func (stack *Dstack) DMaxHeapify(i int) {
	largest := i
	for k := 0; k < stack.Dary; k++ {
		child := stack.Dchild(k, i)
		if child < stack.Size && stack.Heaps[child] > stack.Heaps[largest] {
			largest = child
		}
	}
	if largest != i {
		stack.Heaps[i], stack.Heaps[largest] = stack.Heaps[largest], stack.Heaps[i]
		stack.DMaxHeapify(largest)
	}
}

func (stack *Dstack) BuildDMaxHeap() {
	for i := 3 ; i >= 0 ; i-- {
		stack.DMaxHeapify(i)
	}
}

//去掉并返回stack中的最大元素
func (stack *Dstack) DheapExtractMax() (int, error) {
	if stack.Size < 1 {
		return 0, errors.New("heap underflow")
	}
	max := stack.Heaps[0]
	stack.Heaps[0] = stack.Heaps[stack.Size-1]
	stack.Size--
	stack.DMaxHeapify(0)
	return max, nil
}

func (stack *Dstack) Dinsert(key int) {
	stack.Heaps = append(stack.Heaps, key)
	i := stack.Size
	stack.Size++
	stack.Length++
	for i > 0 && stack.Heaps[stack.Dparent(i)] < stack.Heaps[i] {
		stack.Heaps[stack.Dparent(i)], stack.Heaps[i] = stack.Heaps[i] ,stack.Heaps[stack.Dparent(i)]
		i = stack.Dparent(i)
	}
}

func (stack *Dstack) DinsertKey(i, key int) error {
	if key < stack.Heaps[i] {
		return errors.New("new key is smaller than current key")
	}
	stack.Heaps[i] = key
	for i > 0 &&stack.Heaps[Parent(i)] < stack.Heaps[i]  {
		stack.Heaps[Parent(i)], stack.Heaps[i] = stack.Heaps[i], stack.Heaps[Parent(i)]
		i = Parent(i)
	}
	return nil
}

//young矩阵
type Young struct {
	Data [][]int
	Row int
	Col int
}

func (young *Young) IsYoung() bool {
	for i := 0; i < young.Row; i++ {
		for j := 0; j < young.Col - 1; j++ {
			if young.Data[i][j] > young.Data[i][j+1] {
				return false
			}
		}
	}
	for i := 0; i < young.Col; i++ {
		for j := 0; j < young.Row - 1; j++ {
			if young.Data[j][i] > young.Data[j+1][i] {
				return false
			}
		}
	}
	return true
}