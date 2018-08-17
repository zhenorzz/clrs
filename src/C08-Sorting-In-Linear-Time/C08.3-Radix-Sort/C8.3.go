package C08_3_Radix_Sort

func RadixSort(A []int, d int) []int{
	x := 1
	for d > 0 {
		t := make([]int, len(A))
		B := make([]int, len(A))
		for i := 0; i < len(A); i++ {
			t[i] = A[i] % (x * 10) / x
		}
		x *= 10
		C := make([]int, 10)
		for i := 0; i < len(A); i++ {
			C[t[i]]++
		}
		//确定元素位置
		for i := 1; i < 10; i++ {
			C[i] += C[i-1]
		}
		for i := len(A) - 1; i >= 0; i-- {
			C[t[i]]--
			B[C[t[i]]] = A[i]
		}
		A = B
		d--
	}
	return A
}
