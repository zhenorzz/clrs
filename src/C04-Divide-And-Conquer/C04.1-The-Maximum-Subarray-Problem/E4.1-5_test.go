package C04_1_The_Maximum_Subarray_Problem

import (
	"testing"
	"fmt"
)

func TestLinearTime(t *testing.T) {
	A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Println(LinearTime(A))
}
