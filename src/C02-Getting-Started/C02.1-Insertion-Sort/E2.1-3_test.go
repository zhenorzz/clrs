package C02_1_Insertion_Sort

import (
	"testing"
	"fmt"
)

func TestLinearSearch(t *testing.T) {
	numbers := []int{5, 2, 4, 6, 1, 3}
	pos, err := LinearSearch(numbers, 4)
	fmt.Println(pos, err)
	pos, err = LinearSearch(numbers, 7)
	fmt.Println(pos, err)
}
