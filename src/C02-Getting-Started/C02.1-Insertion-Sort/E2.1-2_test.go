package C02_1_Insertion_Sort

import (
	"testing"
	"fmt"
)

func TestDesc(t *testing.T) {
	numbers := []int{5, 2, 4, 6, 1, 3}
	Desc(numbers)
	fmt.Println(numbers)
}
