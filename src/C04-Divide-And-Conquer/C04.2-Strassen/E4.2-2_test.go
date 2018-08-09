package C04_2_Strassen

import (
	"testing"
	"fmt"
)

func TestStrassen(t *testing.T) {
	fmt.Println(Strassen([][]int{{1, 3}, {7, 5}}, [][]int{{6, 8}, {4, 2}}))
}
