package C07_Quick_Sort

import (
	"math/rand"
	"fmt"
)

func QuickSort(A []int, p, r int) {
	if p < r {
		q := Partition(A, p, r)
		QuickSort(A, p, q-1)
		QuickSort(A, q+1, r)
	}
}
func Partition(A []int, p int, r int) int {
	randomIndex := Randomize(p, r)
	A[r], A[randomIndex] = A[randomIndex], A[r]
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	i++
	A[i], A[r] = A[r], A[i]
	return i
}

func Randomize(min, max int) int {
	return rand.Intn(max-min) + min
}

func EqualPartition(A []int, p int, r int) (int, int) {
	x := A[r]
	A[r], A[p] = A[p], A[r]
	i := p - 1
	k := p
	for j := p + 1; j <= r-1; j++ {
		if A[j] < x {
			i = i + 1
			k = k + 2
			A[i], A[j] = A[j], A[i]
			A[k], A[j] = A[j], A[k]
		}
		if A[j] == x {
			k = k + 1
			A[k], A[j] = A[j], A[k]
		}
		fmt.Println(A)
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1, k + 1
}
