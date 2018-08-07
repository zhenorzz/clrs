package C02_Getting_Started

import (
	"testing"
	"fmt"
)

func TestBubbleSort(t *testing.T) {
	numbers := []int{1, 6, 3, 2, 5, 9, 7}
	BubbleSort(numbers)
	fmt.Println(numbers)
}
