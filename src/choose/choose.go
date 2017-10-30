package choose

import (
	"imath"
)

func FindMinAndMax(A []int) (min, max int) {
	n := len(A)
	start := 0
	if n%2 == 0 {
		start = 2
		if A[0] < A[1] {
			min = A[0]
			max = A[1]
		}
	} else {
		start = 1
		min, max = A[0], A[0]
	}
	for i := start; i < n-1; i++ {
		if A[i] < A[i+1] {
			if min > A[i] {
				min = A[i]
			}
			if max < A[i+1] {
				max = A[i+1]
			}
		}
		if A[i] > A[i+1] {
			if min > A[i+1] {
				min = A[i+1]
			}
			if max < A[i] {
				max = A[i]
			}
		}
	}
	return
}

//data := []int{6,0,2,0,1,3,4,6,1,3,2}
//fmt.Println(data, choose.RandomizedSelect(data,0,10,10))
func RandomizedSelect(A []int, p, r, i int) int {
	//只存在一个元素i=0
	if p == r {
		return A[p]
	}
	q := RandomizePartition(A, p, r)
	k := q - p
	if i == k {
		return A[q]
	} else if i < k {
		return RandomizedSelect(A, p, q-1, i)
	} else {
		return RandomizedSelect(A, q+1, r, i-k-1)
	}
}

func RandomizedSelectLoop(A []int, p, r, i int) int {
	//只存在一个元素i=0
	if p == r {
		return A[p]
	}
	for p < r {
		q := RandomizePartition(A, p, r)
		k := q - p
		if i == k {
			return A[q]
		} else if i < k {
			r = q - 1
		} else {
			p = q + 1
			i = i - k - 1
		}
	}
	return A[p]
}

func RandomizePartition(A []int, p, r int) int {
	random := imath.Randomize(p, r)
	A[r], A[random] = A[random], A[r]
	x := A[r]
	i := p - 1
	//分两边
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	q := i + 1
	return q
}
