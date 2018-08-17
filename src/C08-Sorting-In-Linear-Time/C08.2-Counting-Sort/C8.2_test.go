package C08_2_Counting_Sort

import (
	"testing"
	"fmt"
)

func TestCountingSort(t *testing.T) {
	A := []int{6, 8, 2, 3, 6, 7, 2, 0, 1}
	fmt.Println(CountingSort(A, 8))
}
