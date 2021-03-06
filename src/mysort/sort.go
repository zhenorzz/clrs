package mysort

import (
	"imath"
	"fmt"
)

//插入排序 非降序
func InsertionAsc(s []int) {
	for j := 1; j < len(s); j++ {
		value, i := s[j], j - 1
		for i >= 0 && s[i] > value {
			s[i+1] = s[i]
			i--
		}
		s[i+1] = value
	}
	fmt.Println(s)
}

//插入排序 非升序
func InsertionDesc(s []int) {
	for j := 1; j < len(s); j++ {
		value := s[j]
		i := j - 1
		for i >= 0 && s[i] < value {
			s[i+1] = s[i]
			i--
		}
		s[i+1] = value
	}
	fmt.Println(s)
}

//递归插入排序 非降序
func RecursionInsertion(s []int, n int) []int {
	if n < 0 {
	} else {
		RecursionInsertion(s, n-1)
		for i := 0; i < n; i++ {
			if s[i] > s[n] {
				t := s[n]
				s[n] = s[i]
				s[i] = t
			}
		}

	}
	return s
}

//find
func Find(s []int, v int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return i
		}
	}
	return 0
}

//二分查找
func BinaryFind(s []int, v int) (int, bool) {
	length := len(s)
	n := length / 2
	for n > 0 {
		if s[n] > v {
			if length-1 == n {
				return n, false
			}
			length = n
			n /= 2
		} else if s[n] < v {
			if length-1 == n {
				return n, false
			}
			n = (length + n) / 2
		} else {
			return n, true
		}
	}
	return n, false
}

//2.3-7
func IsExistSum(s []int, x int) bool {
	length := len(s)
	for i := 0; i < length; i++ {
		if _, ok := BinaryFind(s, x-s[i]); ok == true {
			return true
		}
	}
	return false
}

//相加
func AddBytes(a []byte, b []byte) []byte {
	c := make([]byte, len(a)+1)
	carry := []byte{0}
	i := 1
	for i < len(a) {
		c[i] = a[i-1] + b[i-1] + carry[0]
		if c[i] == 3 {
			c[i] = 1
			carry[0] = 1
		} else if c[i] == 2 {
			c[i] = 0
			carry[0] = 1
		}
		i++
	}
	c[0] = carry[0]
	return c
}

//选择算法
func SelectSort(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[min] > s[j] {
				min = j
			}
		}
		if min != i {
			t := s[i]
			s[i] = s[min]
			s[min] = t
		}
	}
	return s
}

func Merge(s []int, p, q, r int) {
	leftLength := q - p + 1
	rightLength := r - q
	leftArr := make([]int, leftLength+1)
	rightArr := make([]int, rightLength+1)
	for i := 0; i < leftLength; i++ {
		leftArr[i] = s[p+i]
	}
	for j := 0; j < rightLength; j++ {
		rightArr[j] = s[q+j+1]
	}
	leftArr[leftLength] = imath.MaxInt
	rightArr[rightLength] = imath.MaxInt
	i := 0
	j := 0
	for k := p; k <= r; k++ {
		if leftArr[i] < rightArr[j] {
			s[k] = leftArr[i]
			i++
		} else {
			s[k] = rightArr[j]
			j++
		}
	}

}

func MergeSort(s []int, p, r int) {
	if p < r && len(s) > 0 {
		q := (p + r) / 2
		MergeSort(s, p, q)
		MergeSort(s, q+1, r)
		Merge(s, p, q, r)
	}
}

func Bubble(s []int) []int {
	for i := 0; i < len(s)-2; i++ {
		for j := len(s) - 1; j > i+1; j-- {
			if s[j] < s[j-1] {
				t := s[j-1]
				s[j-1] = s[j]
				s[j] = t
			}
		}
	}
	return s
}
