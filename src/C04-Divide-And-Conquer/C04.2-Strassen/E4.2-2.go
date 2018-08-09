package C04_2_Strassen

func Strassen(A [][]int, B [][]int) [][]int {
	size := len(A)
	if size == 1 {
		matrixMul := make([][]int, 1)
		matrixMul[0] = make([]int, 1)
		matrixMul[0][0] = A[0][0] * B[0][0]
		return matrixMul
	}
	C := make([][]int, size)
	for i := 0; i < size; i++ {
		C[i] = make([]int, size)
	}
	size = size / 2
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
		for j := 0; j < size; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+size]
			A21[i][j] = A[i+size][j]
			A22[i][j] = A[i+size][j+size]

			B11[i][j] = B[i][j]
			B12[i][j] = B[i][j+size]
			B21[i][j] = B[i+size][j]
			B22[i][j] = B[i+size][j+size]
		}
	}
	//M1 = (A11 + A22) x (B11 + B22)
	AA = matrixAdd(A11, A22)
	BB = matrixAdd(B11, B22)
	M1 = Strassen(AA, BB)

	//M2 = (A21 + A22) x B11
	AA = matrixAdd(A21, A22)
	M2 = Strassen(AA, B11)

	//M3=A11(B12-B22)
	BB = matrixSub(B12, B22)
	M3 = Strassen(A11, BB)

	//M4=A22(B21-B11)
	BB = matrixSub(B21, B11)
	M4 = Strassen(A22, BB)

	//M5=(A11+A12)B22
	AA = matrixAdd(A11, A12)
	M5 = Strassen(AA, B22)

	//M6=(A21-A11)(B11+B12)
	AA = matrixSub(A21, A11)
	BB = matrixAdd(B11, B12)
	M6 = Strassen(AA, BB)

	//M7=(A12-A22)(B21+B22)
	AA = matrixSub(A12, A22)
	BB = matrixAdd(B21, B22)
	M7 = Strassen(AA, BB)

	//C11 = M1 + M4 - M5 + M7
	C11 = matrixAdd(matrixAdd(M1, M4), matrixSub(M7, M5))

	//C12 = M3 + M5
	C12 = matrixAdd(M3, M5)

	//C21 = M2 + M4
	C21 = matrixAdd(M2, M4)

	//C22 = M1 + M3 - M2 + M6
	C22 = matrixAdd(matrixSub(M1, M2), matrixAdd(M3, M6))

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			C[i][j] = C11[i][j]
			C[i][j+size] = C12[i][j]
			C[i+size][j] = C21[i][j]
			C[i+size][j+size] = C22[i][j]
		}
	}
	return C
}

//矩阵加法
func matrixAdd(A [][]int, B [][]int) [][]int {
	size := len(A)
	C := make([][]int, size)
	for i := 0; i < size; i++ {
		C[i] = make([]int, size)
		for j := 0; j < size; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

//矩阵减法
func matrixSub(A [][]int, B [][]int) [][]int {
	size := len(A)
	C := make([][]int, size)
	for i := 0; i < size; i++ {
		C[i] = make([]int, size)
		for j := 0; j < size; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}
