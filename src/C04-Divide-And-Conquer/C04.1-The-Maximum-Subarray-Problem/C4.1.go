package C04_1_The_Maximum_Subarray_Problem

import "imath"

func findMaxCrossingSubarray(A []int, low, mid, high int) (maxLeft, maxRight, total int) {
	leftSum := imath.MinInt
	sum := 0
	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := imath.MinInt
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += A[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	total = leftSum + rightSum
	return
}

func FindMaximumSubarray(A []int, low, high int) (int, int, int) {	
	if high == low {
		return low, high, A[low]
	}

	mid := (low + high) / 2
	leftLow, leftHigh, leftSum := FindMaximumSubarray(A, low, mid)
	rightLow, rightHigh, rightSum := FindMaximumSubarray(A, mid+1, high)
	crossLow, crossHigh, crossSum := findMaxCrossingSubarray(A, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLow, rightHigh, rightSum
	} else {
		return crossLow, crossHigh, crossSum
	}
}
