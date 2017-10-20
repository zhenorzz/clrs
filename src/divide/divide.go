package divide

import (
	"imath"
)

//最大子数组跨越
func FindMaxCrossingSubarry(A []int, low, mid, high int) (int, int, int) {
	leftSum := imath.MinInt
	sum := 0
	maxLeft := 0
	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := imath.MinInt
	sum = 0
	maxRight := 0
	for i := mid + 1; i <= high; i++ {
		sum += A[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft,maxRight,leftSum + rightSum
}

//最大子数组分治
//divide.FindMaximumSubarry([]int{13,-3,-25,20,-3,-16,-23,18,20,-7,12,-5,-22,15,-4,7},0,13)
func FindMaximumSubarry(A []int, low, high int) (int, int, int)  {
	if high == low {
		return low, high, A[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := FindMaximumSubarry(A, low, mid)
		rightLow, rightHigh, rightSum := FindMaximumSubarry(A, mid+1, high)
		crossLow, crossHigh, crossSum := FindMaxCrossingSubarry(A, low, mid, high)
		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if  rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

//暴力求解日期
//	fmt.Println(divide.FindIODate([]int{100,113,110,85,105,102,86,63,81,101,94,106,101,79,94,90,97}))

func FindIODate(A []int) (int, int, int) {
	length := len(A)
	sum := 0
	low := 0
	high := 0
	for i := 0; i < length; i++ {
		temp := 0
		j := i+1

		for j < length - 1  {
			temp += A[j] - A[j-1]
			if temp > sum {
				sum = temp
				low = i
				high =j
			}
			j++
		}
	}
	return low, high, sum
}
//暴力求解子数组问题
func ForceMaximumSubarry(A []int) (int, int, int) {
	length := len(A)
	sum := 0
	low := 0
	high := 0
	for i := 0; i < length; i++ {
		temp := 0
		j := i

		for j < length  {
			temp += A[j]
			if temp > sum {
				sum = temp
				low = i
				high =j
			}
			j++
		}
	}
	return low, high, sum
}

//暴力求解子数组问题
func LineMaximumSubarry(A []int) (int, int, int) {
	length := len(A)
	sum := A[0]
	low := 0
	high := 0
	temp := 0
	for i := 0; i < length; i++ {
		if temp == 0 {
			low = i
		}
		temp += A[i]
		if temp > sum {
			sum = temp
			high = i
		}
		if temp < 0 {
			temp = 0
		}
	}
	return low, high, sum
}
