package C02_Getting_Started

import (
	"testing"
	"fmt"
)

func TestHornerRule(t *testing.T) {
	numbers := []int{2, -1, -3, 1, -5}
	fmt.Println(HornerRule(numbers, 7))
}
