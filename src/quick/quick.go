package quick

import "imath"

//data := []int{13,19,9,5,12,8,7,4,21,2,6,11}
//quick.Sort(data,0,11)
//fmt.Println(data)
func Sort(A []int, p, r int) {
	if p < r {
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
		q := i+1
		left, right := 0, 0
		//左边聚集 key
		for i = q - 1; i >= p; i-- {
			if A[i] == x {
				A[q-1-left], A[i] = A[i], A[q-1-left]
				left++
			}
		}
		//右边聚集 key
		for i = q + 1; i <= r; i++ {
			if A[i] == x {
				A[q+1+right], A[i] = A[i], A[q+1+right]
				right++
			}
		}
		Sort(A, p, q-1-left)
		Sort(A, q+1+right, r)
	}
}