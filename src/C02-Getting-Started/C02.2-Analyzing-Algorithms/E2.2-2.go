package C02_2_Analyzing_Algorithms

func SelectionSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		min := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		numbers[min], numbers[i] = numbers[i], numbers[min]
	}
}
