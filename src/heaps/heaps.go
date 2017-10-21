package heaps

func Left(i int) int {
	return (i<<1) + 1
}

func Right(i int) int {
	return (i<<1) + 2
}

func MaxHeapify(A []int, i int) []int {
	length := len(A)
	left := Left(i)
	largest := i
	if length > left && A[left] > A[i] {
		largest = left
	}
	right := Right(i)
	if length > right && A[right] > A[largest] {
		largest = right
	}
	if largest != i {
		temp := A[i]
		A[i] = A[largest]
		A[largest] = temp
		MaxHeapify(A, largest)
	}
	return A
}

func MinHeapify(A []int, i int) []int {
	length := len(A)
	left := Left(i)
	largest := i
	if length > left && A[left] < A[i] {
		largest = left
	}
	right := Right(i)
	if length > right && A[right] < A[largest] {
		largest = right
	}
	if largest != i {
		temp := A[i]
		A[i] = A[largest]
		A[largest] = temp
		MaxHeapify(A, largest)
	}
	return A
}