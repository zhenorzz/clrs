package quick

//data := []int{13,19,9,5,12,8,7,4,21,2,6,11}
//quick.Sort(data,0,11)
//fmt.Println(data)
func Sort(A []int, p, r int) {
	if p < r {
		q := Partition(A, p, r)
		Sort(A, p, q-1)
		Sort(A, q, r)
	}
}

func Partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}