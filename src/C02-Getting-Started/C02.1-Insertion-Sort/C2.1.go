package C02_1_Insertion_Sort

func Asc(numbers []int) {
	for current := 1; current < len(numbers); current++ {
		currentValue := numbers[current]
		prev := current - 1
		for prev >= 0 && numbers[prev] > currentValue {
			numbers[prev+1] = numbers[prev]
			prev--
		}
		numbers[prev+1] = currentValue
	}
}
