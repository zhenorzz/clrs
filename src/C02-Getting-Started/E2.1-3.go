package C02_Getting_Started

import "errors"

func LinearSearch(numbers []int, v int) (int, error) {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == v {
			return i, nil
		}
	}
	return 0, errors.New("it can not find in the given array")
}
