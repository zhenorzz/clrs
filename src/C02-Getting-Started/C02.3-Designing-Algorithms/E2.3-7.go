package C02_3_Designing_Algorithms

func FindX(A []int, x int) bool {
	MergeSort(A, 0, len(A)-1)
	i, j := 0, len(A)-1
	for i < j {
		if A[i]+A[j] == x {
			return true
		} else if A[i]+A[j] < x {
			i++
		} else {
			j--
		}
	}
	return false;
}
