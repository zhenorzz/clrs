package dynamic

import (
	"fmt"
	"imath"
	"strconv"
)

//dynamic.MatrixChainOrder([]int{5,10,3,12,5,50,6})
func MatrixChainOrder(p []int) {
	n := len(p)
	m := make([][]int, n)
	s := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		s[i] = make([]int, n)
	}
	for l := 2; l < n; l++ {
		for i := 1; i < n-l+1; i++ {
			j := i + l - 1
			m[i][j] = imath.MaxInt
			for k := i; k <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}
	matrix := PrintOptimalParens(s, 1, n-1)
	fmt.Println(matrix)
}

func PrintOptimalParens(s [][]int, i, j int) (matrix string) {
	if i == j {
		matrix += "A" + strconv.Itoa(i)
	} else {
		matrix += "("
		matrix += PrintOptimalParens(s, i, s[i][j])
		matrix += PrintOptimalParens(s, s[i][j]+1, j)
		matrix += ")"
	}
	return matrix
}
