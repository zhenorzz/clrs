package matrix

//矩阵乘积
//fmt.Println(matrix.SquareMatrixMultiply([][]int{{1,2,3},{2,4,1}}, [][]int{{1,2,3},{2,4,1}}))
func SquareMatrixMultiply(A [][]int, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			C[i][j] = 0
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}
