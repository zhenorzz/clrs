package liner

//data := []int{6,0,2,0,1,3,4,6,1,3,2}
//fmt.Println(liner.CountingSort(data,6))
func CountingSort(A []int, k int) []int {
	//下标从0开始 所以区间范围要加1
	k++
	C := make([]int, k)
	B := make([]int,len(A))
	for i := 0; i < len(A); i++ {
		C[A[i]]++
	}
	//确定元素位置
	for i := 1; i < k; i++ {
		C[i] += C[i-1]
	}
	for i := len(A)-1; i >=0 ; i-- {
		C[A[i]]--
		B[C[A[i]]] = A[i]
	}
	return B
}
