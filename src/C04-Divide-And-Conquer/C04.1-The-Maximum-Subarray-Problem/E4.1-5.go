package C04_1_The_Maximum_Subarray_Problem

import "imath"

func LinearTime(A []int) (low, high, sum int) {
	sum = imath.MinInt
	tempSum := 0
	tempLow := 0
	for i := 0; i < len(A); i++ {
		tempSum += A[i]
		if tempSum > sum {
			low = tempLow
			high = i
			sum = tempSum
		}
		if tempSum < 0 {
			tempSum = 0
			tempLow = i + 1
		}
	}
	return
}
