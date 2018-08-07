package C02_3_Designing_Algorithms

import (
	"testing"
	"fmt"
)

func TestMergeSort(t *testing.T) {
	numbers := []int{5, 2, 4, 6, 1}
	MergeSort(numbers, 0, 4)
	fmt.Println(numbers)
}
