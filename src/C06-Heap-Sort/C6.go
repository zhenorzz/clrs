package C06_Heap_Sort

import (
	"imath"
)

type Heap struct {
	Node   []int
	Length int
	Size   int
}

func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return 2*i + 1
}

func Right(i int) int {
	return (2 * i) + 2
}

func (heap Heap) MaxHeapify(i int) {
	l := Left(i)
	r := Right(i)
	largest := 0
	if l < heap.Size && heap.Node[l] > heap.Node[i] {
		largest = l
	} else {
		largest = i
	}
	if r < heap.Size && heap.Node[r] > heap.Node[largest] {
		largest = r
	}
	if largest != i {
		heap.Node[i], heap.Node[largest] = heap.Node[largest], heap.Node[i]
		heap.MaxHeapify(largest)
	}
}

func (heap *Heap) Build() {
	heap.Size = heap.Length
	for i := heap.Length/2 - 1; i >= 0; i-- {
		heap.MaxHeapify(i)
	}
}

func (heap Heap) Sort() {
	heap.Build()
	for i := heap.Length - 1; i >= 1; i-- {
		heap.Node[0], heap.Node[i] = heap.Node[i], heap.Node[0]
		heap.Size--
		heap.MaxHeapify(0)
	}
}

func (heap Heap) Maximum() int {
	return heap.Node[0]
}

func (heap Heap) ExtractMax() int {
	max := heap.Node[0]
	heap.Node[0] = heap.Node[heap.Size-1]
	heap.Size--
	heap.MaxHeapify(0)
	return max
}

func (heap Heap) Update(i, key int) {
	heap.Node[i] = key
	for i > 0 && heap.Node[Parent(i)] < heap.Node[i] {
		heap.Node[Parent(i)], heap.Node[i] = heap.Node[i], heap.Node[Parent(i)]
		i = Parent(i)
	}
}

func (heap Heap) Insert(key int) {
	heap.Size++
	heap.Node[heap.Size-1] = imath.MinInt
	heap.Update(heap.Size-1, key)
}

func (heap Heap) Delete(i int) {
	heap.Node[i] = heap.Node[heap.Size-1]
	heap.Size--
	heap.MaxHeapify(i)
}
