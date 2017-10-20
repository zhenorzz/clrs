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