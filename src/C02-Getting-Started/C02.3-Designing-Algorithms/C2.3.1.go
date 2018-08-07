package C02_3_Designing_Algorithms

import "imath"

func merge(A []int, p, q, r int) {
	n1, n2 := q-p+1, r-q
	L, R := make([]int, n1+1), make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[q+i+1]
	}
	L[n1], R[n2] = imath.MaxInt, imath.MaxInt
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if L[i] < R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

func MergeSort(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		merge(A, p, q, r)
	}
}
