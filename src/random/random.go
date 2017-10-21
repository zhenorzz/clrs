package random

import (
	"time"
	"math/rand"
)

func Randomize(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func swap(a int, b int) (int, int) {
	return b, a
}

func RandomizeInPlace(A []int) []int {
	var randomIndex int
	n := len(A)
	for i := 0; i < n; i++ {
		randomIndex = Randomize(i, n)
		A[i], A[randomIndex] = swap(A[i], A[randomIndex])
	}
	return A
}

//fmt.Println(random.PermuteBySorting([]int{2,4,1,3}))
func PermuteBySorting(A []int) []int {
	n := len(A)
	P := make([]int, n)
	for i := 0; i < n; i++ {
		P[i] = Randomize(i, n * n * n)
	}

	for i := 1; i < n; i++ {
		key := P[i]
		key2 := A[i]
		j := i - 1
		for j >= 0 && P[j] > key {
			P[j+1] = P[j]
			A[j+1] = A[j]
			j--
		}
		P[j+1] = key
		A[j+1] = key2
	}
	return A
}

