package dynamic

import "fmt"

func Length(X, Y []int) {
	m := len(X) + 1
	n := len(Y) + 1
	c := make([][]int, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if X[i-1] == Y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
			} else {
				c[i][j] = c[i][j-1]
			}
		}
	}
	Print(c, X, Y)
}

func Print(c [][]int, X, Y []int) {
	i := len(X)
	j := len(Y)
	n := c[i][j]
	s := make([]int, n)
	for i > 0 && j > 0 {
		if X[i-1] == Y[j-1] {
			n--
			s[n] = X[i-1]
			j--
			i--
		} else if c[i-1][j] >= c[i][j-1] {
			i--
		} else {
			j--
		}
	}
	fmt.Println(s)
}
