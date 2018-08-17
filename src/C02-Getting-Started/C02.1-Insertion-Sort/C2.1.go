package C02_1_Insertion_Sort

func Asc(A []int) {
	for i := 1; i < len(A); i++ {
		currentValue := A[i]
		j := i - 1
		for j >= 0 && A[j] > currentValue {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = currentValue
	}
}
