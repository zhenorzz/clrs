package dynamic

import (
	"imath"
)

var p = []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30}

func CutRod(n int) int {
	if n == 0 {
		return 0
	}
	q := imath.MinInt
	for i := 1; i <= n; i++ {
		q = imath.Max(q, p[i-1]+CutRod(n-i))
	}
	return q
}

func MemoizedCutRod(n int) int {
	r := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		r[i] = imath.MinInt
	}
	return MemoizedCutRodAux(n, r)
}
func MemoizedCutRodAux(n int, r []int) int {
	if r[n] >= 0 {
		return r[n]
	}
	q := imath.MinInt
	if n == 0 {
		q = 0
	} else {
		for i := 1; i <= n; i++ {
			q = imath.Max(q, p[i-1]+MemoizedCutRodAux(n-i, r))
		}
	}
	r[n] = q
	return q
}

func BottomUpCutRod(n int) int {
	r := make([]int, n+1)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q := imath.MinInt
		for i := 1; i <= j; i++ {
			q = imath.Max(q, p[i-1]+r[j-i])
		}
		r[j] = q
	}
	return r[n]
}

func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func DynFib(n int) int {
	if n <= 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, a + b
	}
	return b
}
