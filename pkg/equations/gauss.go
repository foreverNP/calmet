package equations

// GaussMethod
// Метод Гаусса для решения СЛАУ
// A - матрица коэффициентов
// B - вектор свободных членов
func GaussMethod(A [][]float64, B []float64) []float64 {
	N := len(A)
	X := make([]float64, N)

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			l := A[j][i] / A[i][i]
			B[j] = B[j] - l*B[i]

			A[j][i] = 0
			for k := i + 1; k < N; k++ {
				A[j][k] = A[j][k] - l*A[i][k]
			}
		}
	}

	//Обратный ход
	for i := N - 1; i >= 0; i-- {
		for j := N - 1; j > i; j-- {
			B[i] = B[i] - X[j]*A[i][j]
		}
		X[i] = B[i] / A[i][i]
	}

	return X
}
