package C02_Getting_Started

func BubbleSort(A []int) {
	for i := 0; i < len(A); i++ {
		for j := len(A) - 1; j > i; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
		}
	}
}
