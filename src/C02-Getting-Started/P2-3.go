package C02_Getting_Started

func HornerRule(A []int, x int) int {
	y := 0
	for i := 0; i < len(A); i++ {
		y = A[i]+ x * y
	}
	return y
}
