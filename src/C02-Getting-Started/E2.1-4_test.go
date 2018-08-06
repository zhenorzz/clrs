package C02_Getting_Started

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	A := []int{1, 0, 1, 1, 1, 0, 1}
	B := []int{0, 1, 1, 0, 1, 0, 0}
	fmt.Println(Add(A, B))
}
