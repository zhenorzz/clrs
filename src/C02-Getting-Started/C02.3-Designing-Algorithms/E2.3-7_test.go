package C02_3_Designing_Algorithms

import (
	"testing"
	"fmt"
)

func TestFindX(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(FindX(numbers, 0))
}
