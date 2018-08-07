package C02_1_Insertion_Sort

import (
	"testing"
	"fmt"
)

func TestSelectionSort(t *testing.T) {
	numbers := []int{5, 2, 4, 6, 1, 3}
	SelectionSort(numbers)
	fmt.Println(numbers)
}
