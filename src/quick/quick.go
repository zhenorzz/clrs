package quick

import (
	"imath"
)

//data := []int{13,19,9,5,12,8,7,4,21,2,6,11}
//quick.Sort(data,0,11)
//fmt.Println(data)
func Sort(A []int, p, r int) {
	if p < r && len(A) >= (p-r) {
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
		q := i + 1
		A[q], A[r] = A[r], A[q]
		left := 0
		//聚集 key
		for i = q - 1; i >= p; i-- {
			if A[i] == x {
				A[q-1-left], A[i] = A[i], A[q-1-left]
				left++
			}
		}
		Sort(A, p, q-1-left)
		Sort(A, q+1, r)
	}
}

//data := []int{2,2,2,2,2,2,2,2,2,2,2,2}
func Hoare(A []int, p, r int) {
	if p >= r {
		return
	}
	x := A[p]
	i := p
	j := r
	for {
		j--
		for i < r && A[i] < x {
			i++
		}
		for j > p && A[j] >= x {
			j--
		}
		if i < j {
			A[i], A[j] = A[j], A[i]
		} else {
			break
		}
	}
	Hoare(A, p, j)
	Hoare(A, j+1, r)
}
