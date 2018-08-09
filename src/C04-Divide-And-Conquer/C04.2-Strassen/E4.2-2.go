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
	P1 := make([][]int, size)
	P2 := make([][]int, size)
	P3 := make([][]int, size)
	P4 := make([][]int, size)
	P5 := make([][]int, size)
	P6 := make([][]int, size)
	P7 := make([][]int, size)
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

		P1[i] = make([]int, size)
		P2[i] = make([]int, size)
		P3[i] = make([]int, size)
		P4[i] = make([]int, size)
		P5[i] = make([]int, size)
		P6[i] = make([]int, size)
		P7[i] = make([]int, size)
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
	//P1 = (A11 + A22) x (B11 + B22)
	AA = matrixAdd(A11, A22)
	BB = matrixAdd(B11, B22)
	P1 = Strassen(AA, BB)

	//P2 = (A21 + A22) x B11
	AA = matrixAdd(A21, A22)
	P2 = Strassen(AA, B11)

	//P3=A11(B12-B22)
	BB = matrixSub(B12, B22)
	P3 = Strassen(A11, BB)

	//P4=A22(B21-B11)
	BB = matrixSub(B21, B11)
	P4 = Strassen(A22, BB)

	//P5=(A11+A12)B22
	AA = matrixAdd(A11, A12)
	P5 = Strassen(AA, B22)

	//P6=(A21-A11)(B11+B12)
	AA = matrixSub(A21, A11)
	BB = matrixAdd(B11, B12)
	P6 = Strassen(AA, BB)

	//P7=(A12-A22)(B21+B22)
	AA = matrixSub(A12, A22)
	BB = matrixAdd(B21, B22)
	P7 = Strassen(AA, BB)

	//C11 = P1 + P4 - P5 + P7
	C11 = matrixAdd(matrixAdd(P1, P4), matrixSub(P7, P5))

	//C12 = P3 + P5
	C12 = matrixAdd(P3, P5)

	//C21 = P2 + P4
	C21 = matrixAdd(P2, P4)

	//C22 = P1 + P3 - P2 + P6
	C22 = matrixAdd(matrixSub(P1, P2), matrixAdd(P3, P6))

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
