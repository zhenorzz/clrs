package mysort

//插入排序 非降序
func InsertionAsc(s []int) []int {
	for j := 1; j<len(s); j++ {
		key := s[j]
		i := j-1
		for i >= 0 && s[i] > key {
			s[i+1] = s[i]
			i--
		}
		s[i+1] = key
	}
	return s
}

//插入排序 非升序
func InsertionDesc(s []int) []int {
	for j := 1; j<len(s); j++ {
		key := s[j]
		i := j-1
		for i >= 0 && s[i] < key {
			s[i+1] = s[i]
			i--
		}
		s[i+1] = key
	}
	return s
}

//find
func Find(s []int, v int) int{
	for i := 0; i<len(s); i++{
		if s[i] == v {
			return i
		}
	}
	return 0
}

//相加
func AddBytes(a []byte, b []byte) []byte{
	c := make([]byte,len(a)+1)
	carry := []byte{0}
	i := 1
	for i < len(a) {
		c[i] = a[i-1]+b[i-1]+carry[0]
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
		for j := i+1; j<len(s); j++ {
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

func Merge(s []int, p, q, r int) []int {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+1)
	R := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		L[i] = s[p + i]
	}
	for i := 0; i < n2; i++ {
		R[i] = s[q + i + 1]
	}

	L[n1] = 0
	R[n2] = 0
	i := 0
	j := 0
	for k := p; k < r; k++ {
		if L[i] <= R[j] {
			s[k] = L[i]
			i++
		} else {
			s[k] = R[j]
			j++
		}
	}
	return s
}