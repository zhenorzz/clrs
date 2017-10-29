package choose

func FindMinAndMax(A []int) (min, max int) {
	n := len(A)
	start := 0
	if n % 2 == 0 {
		start = 2
		if A[0] < A[1] {
			min = A[0]
			max = A[1]
		}
	} else {
		start = 1
		min, max = A[0], A[0]
	}
	for i := start; i < n-1; i++ {
		if A[i] < A[i+1] {
			if min > A[i] {
				min = A[i]
			}
			if max < A[i+1] {
				max = A[i+1]
			}
		}
		if A[i] > A[i+1] {
			if min > A[i+1] {
				min = A[i+1]
			}
			if max < A[i] {
				max = A[i]
			}
		}
	}
	return
}