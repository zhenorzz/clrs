package C02_3_Designing_Algorithms

import (
	"testing"
	"fmt"
)

func TestMergeSortWithoutSentinel(t *testing.T) {
	numbers := []int{5, 2, 4, 6, 1,3}
	MergeSortWithoutSentinel(numbers, 0, 5)
	fmt.Println(numbers)
}
