package C08_3_Radix_Sort

import (
	"testing"
	"fmt"
)

func TestRadixSort(t *testing.T) {
	A := []int{329, 457, 657, 839, 436, 720, 355}
	A = RadixSort(A, 3)
	fmt.Println(A)
}
