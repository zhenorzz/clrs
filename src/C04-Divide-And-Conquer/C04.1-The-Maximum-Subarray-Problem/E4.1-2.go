package C04_1_The_Maximum_Subarray_Problem

func BruteForce(A []int) (left, right, max int) {
	max = A[0]
	curSum := 0
	for i := 0; i < len(A); i++ {
		curSum = 0
		for j := i; j < len(A); j++ {
			curSum += A[j]
			left = i
			right = j
		}
	}
	return
}
