package C08_2_Counting_Sort

func CountingSort(A []int, k int) []int {
	B := make([]int, len(A))
	C := make([]int, k+1)
	for i := 0; i < len(A); i++ {
		C[A[i]]++
	}
	for i := 1; i <= k; i++ {
		C[i] = C[i] + C[i-1]
	}
	for i := len(A)-1; i >= 0; i-- {
		C[A[i]]--
		B[C[A[i]]] = A[i]
	}
	return B
}
