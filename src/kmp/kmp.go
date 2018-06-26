package kmp

import (
	"fmt"
	"errors"
)

type KMP struct {
	pattern string
	prefix  []int
	size    int
}

func New(pattern string) (*KMP, error) {
	prefix, err := computePrefix(pattern)
	if err != nil {
		return nil, err
	}
	return &KMP{
		pattern: pattern,
		prefix:  prefix,
		size:    len(pattern),
	}, nil
}
func computePrefix(pattern string) ([]int, error) {
	patternLength := len(pattern)
	if patternLength < 2 {
		if patternLength == 0 {
			return nil, errors.New("pattern must contain as least one character")
		}
		return []int{-1}, nil
	}
	next := make([]int, patternLength)
	next[0], next[1] = -1, 0

	pos, count := 2, 0
	for pos < patternLength {
		if pattern[pos-1] == pattern[count] {
			count++
			next[pos] = count
			pos++
		} else {
			if count > 0 {
				count = next[count]
			} else {
				pos++
			}
		}
	}
	return next, nil
}

func (kmp *KMP) FindStringIndex(s string) int {
	if len(s) < kmp.size {
		return -1
	}
	m, i := 0, 0
	for m+i < len(s) {
		if kmp.pattern[i] == s[m+i] {
			if i== kmp.size-1 {
				return m
			}
			i++
		} else {
			m = m + i - kmp.prefix[i]
			if kmp.prefix[i] > -1 {
				i = kmp.prefix[i]
			} else {
				i = 0
			}
		}
	}
	return -1
}

func (kmp *KMP) String() string {
	return fmt.Sprintf("pattern: %v\nprefix: %v", kmp.pattern, kmp.prefix)
}
