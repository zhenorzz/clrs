package dynamic

import "fmt"

func Monotonic(S []int) {
	n := len(S)
	c := make([]int, n)
	b := make([]int, n)
	last := make([]int, n)
	maxLen := 0
	c[0] = 1
	last[0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if S[i] < S[j] && c[j]+1 > c[i] {
				c[i] = c[j] + 1
				b[i] = j
				last[i] = i
				maxLen = c[i]
			} else if c[j] > c[i] {
				maxLen = c[j]
				last[i] = last[j]
			}
		}
	}
	fmt.Println(maxLen)
	printMonotonic(b, S, last[n-1])
}

func printMonotonic(b []int, S []int, last int) {
	if b[last] > 0 {
		printMonotonic(b, S, b[last])
	}
	fmt.Println(S[last])
}
