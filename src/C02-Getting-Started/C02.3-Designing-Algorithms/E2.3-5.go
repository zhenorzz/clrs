package C02_3_Designing_Algorithms

import "errors"

func BinarySearch(A []int, v int) (int, error) {
	left, right := 0, len(A)
	for left < right {
		mid := (left + right) / 2
		if A[mid] == v {
			return mid, nil
		} else if A[mid] < v {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return 0, errors.New("it can not find in the given array")
}
