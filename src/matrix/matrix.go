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

func Strassen(A [][]int, B [][]int, size int) [][]int {
	if size == 1 {
		matrixMul := make([][]int, 1)
		matrixMul[0] = make([]int, 1)
		matrixMul[0][0] = A[0][0] * B[0][0]
		return matrixMul
	} else {
		C := make([][]int, size)
		for i := 0; i < size; i++ {
			C[i] = make([]int, size)
		}
		size >>= 1
		A11 := make([][]int, size)
		A12 := make([][]int, size)
		A21 := make([][]int, size)
		A22 := make([][]int, size)
		B11 := make([][]int, size)
		B12 := make([][]int, size)
		B21 := make([][]int, size)
		B22 := make([][]int, size)
		C11 := make([][]int, size)
		C12 := make([][]int, size)
		C21 := make([][]int, size)
		C22 := make([][]int, size)
		M1 := make([][]int, size)
		M2 := make([][]int, size)
		M3 := make([][]int, size)
		M4 := make([][]int, size)
		M5 := make([][]int, size)
		M6 := make([][]int, size)
		M7 := make([][]int, size)
		AA := make([][]int, size)
		BB := make([][]int, size)
		for i := 0; i < size; i++ {
			A11[i] = make([]int, size)
			A12[i] = make([]int, size)
			A21[i] = make([]int, size)
			A22[i] = make([]int, size)

			B11[i] = make([]int, size)
			B12[i] = make([]int, size)
			B21[i] = make([]int, size)
			B22[i] = make([]int, size)

			C11[i] = make([]int, size)
			C12[i] = make([]int, size)
			C21[i] = make([]int, size)
			C22[i] = make([]int, size)

			M1[i] = make([]int, size)
			M2[i] = make([]int, size)
			M3[i] = make([]int, size)
			M4[i] = make([]int, size)
			M5[i] = make([]int, size)
			M6[i] = make([]int, size)
			M7[i] = make([]int, size)
			AA[i] = make([]int, size)
			BB[i] = make([]int, size)
		}
		//splitting input Matrixes, into 4 submatrices each.
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				A11[i][j] = A[i][j]
				A12[i][j] = A[i][j + size]
				A21[i][j] = A[i + size][j]
				A22[i][j] = A[i + size][j + size]

				B11[i][j] = B[i][j]
				B12[i][j] = B[i][j + size]
				B21[i][j] = B[i + size][j]
				B22[i][j] = B[i + size][j + size]
			}
		}

		//M1 = (A11 + A22) x (B11 + B22)
		matrixAdd(size, A11, A22, &AA)
		matrixAdd(size, B11, B22, &BB)
		M1 = Strassen(AA, BB, size)

		//M2 = (A21 + A22) x B11
		matrixAdd(size, A21, A22, &AA)
		M2 = Strassen(AA, B11, size)

		//M3=A11(B12-B22)
		matrixSub(size, B12, B22, &BB)
		M3 = Strassen(A11, BB, size)

		//M4=A22(B21-B11)
		matrixSub(size, B21, B11, &BB)
		M4 = Strassen(A22, BB, size)

		//M5=(A11+A12)B22
		matrixAdd(size, A11, A12, &AA)
		M5 = Strassen(AA, B22,size)

		//M6=(A21-A11)(B11+B12)
		matrixSub(size, A21, A11, &AA)
		matrixAdd(size, B11, B12, &BB)
		M6 = Strassen(AA, BB,size)

		//M7=(A12-A22)(B21+B22)
		matrixSub(size, A12, A22, &AA)
		matrixAdd(size, B21, B22, &BB)
		M7 = Strassen(AA, BB, size)

		//C11 = M1 + M4 - M5 + M7
		matrixAdd(size, M1, M4, &AA)
		matrixSub(size, M7, M5, &BB)
		matrixAdd(size, AA, BB, &C11)

		//C12 = M3 + M5
		matrixAdd(size, M3, M5, &C12)

		//C21 = M2 + M4
		matrixAdd(size, M2, M4, &C21)

		//C22 = M1 + M3 - M2 + M6
		matrixSub(size, M1, M2, &AA)
		matrixAdd(size, M3, M6, &BB)
		matrixAdd(size, AA, BB, &C22)

		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				C[i][j] = C11[i][j]
				C[i][j + size] = C12[i][j]
				C[i + size][j] = C21[i][j]
				C[i + size][j + size] = C22[i][j]
			}
		}
		return C
	}
}

//矩阵加法
func matrixAdd(size int, A [][]int, B [][]int, C *[][]int) {
	for i := 0; i < size; i++ {
		for j:=0; j < size; j++ {
			(*C)[i][j] = A[i][j] + B[i][j]
		}
	}
}

//矩阵减法
func matrixSub(size int, A [][]int, B [][]int, C *[][]int) {
	for i := 0; i < size; i++ {
		for j:=0; j < size; j++ {
			(*C)[i][j] = A[i][j] - B[i][j]
		}
	}
}