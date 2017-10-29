package liner

import "fmt"

//data := []int{6,0,2,0,1,3,4,6,1,3,2}
//fmt.Println(liner.CountingSort(data,6))
func CountingSort(A []int, k int) []int {
	//下标从0开始 所以区间范围要加1
	k++
	C := make([]int, k)
	B := make([]int,len(A))
	for i := 0; i < len(A); i++ {
		C[A[i]]++
	}
	//确定元素位置
	for i := 1; i < k; i++ {
		C[i] += C[i-1]
	}
	for i := len(A)-1; i >=0 ; i-- {
		C[A[i]]--
		B[C[A[i]]] = A[i]
	}
	return B
}

//data := []int{329,457,657,839,436,720,355}
//fmt.Println(liner.RadixSort(data,10,3))
func RadixSort(A []int, d, k int) []int {
	x := 1
	for k > 0 {
		t := make([]int,len(A))
		B := make([]int,len(A))
		for i := 0; i < len(A); i++ {
			t[i] = A[i]%(x*10)/x
		}
		x *= 10
		C := make([]int, d)
		for i := 0; i < len(A); i++ {
			C[t[i]]++
		}
		//确定元素位置
		for i := 1; i < d; i++ {
			C[i] += C[i-1]
		}
		for i := len(A)-1; i >=0 ; i-- {
			C[t[i]]--
			B[C[t[i]]] = A[i]
		}
		k--
		A = B
	}
	return A
}

func CountingSortInPlace(A []int, k int) {
	fmt.Println(A)
	counts := make([]int, k+1)
	n := len(A)
	for i := 0; i < n; i++ {
		counts[A[i]]++
	}
	fmt.Println(counts)
	for val, sum := 0, 0; val < k+1; val++ {
		temp := counts[val]
		counts[val] = sum
		sum += temp
	}
	fmt.Println(counts)
	for i := n - 1; i >= 0; i-- {
		val := A[i]
		j := counts[val]
		if j < i {
			// Process a fresh cycle. Since the index 'i' moves
			// downward and the counts move upward, it is
			// guaranteed that a value is never moved twice.
			for {
				counts[val]++
				val, A[j] = A[j], val
				j = counts[val]
				if j >= i {
					break
				}
			}
			A[i] = val
		}

	}
}