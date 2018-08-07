package C02_3_Designing_Algorithms

import (
	"testing"
	"fmt"
)

func TestBinarySearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(BinarySearch(numbers, 8))
}
