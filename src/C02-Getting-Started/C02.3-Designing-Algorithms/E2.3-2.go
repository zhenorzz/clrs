package C02_3_Designing_Algorithms

func mergeWithoutSentinel(A []int, p, q, r int) {
	n1, n2 := q-p+1, r-q
	L, R := make([]int, n1), make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[q+i+1]
	}
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if i >= n1 {
			A[k] = R[j]
			j++
			continue
		}
		if j >= n2 {
			A[k] = L[i]
			i++
			continue
		}
		if L[i] < R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

func MergeSortWithoutSentinel(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		merge(A, p, q, r)
	}
}
