package C07_Quick_Sort

import (
	"testing"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	A := []int{1,4,4,6,4}
	q, e := EqualPartition(A, 0, 4)
	fmt.Println(q, e, A)
}
