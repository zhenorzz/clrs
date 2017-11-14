package imath

import (
	"time"
	"math/rand"
)
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt = int(MaxUint >> 1)
	MinInt = -MaxInt - 1
)

func Randomize(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}