package C06_Heap_Sort

import (
	"testing"
	"fmt"
)

func TestSort(t *testing.T) {
	heap := Heap{Node: []int{5, 13, 2, 25, 7, 17, 20, 8, 4}, Length: 9}
	heap.Sort()
	fmt.Println(heap)
}
