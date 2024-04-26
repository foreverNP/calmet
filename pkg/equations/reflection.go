package equations

import "github.com/foreverNP/calmet/pkg/tools"

// ReflectionMethod Метод отражений для решения СЛАУ
// A - матрица коэффициентов
// B - вектор свободных членов
func ReflectionMethod(A [][]float64, B []float64) ([]float64, [][]float64, [][]float64) {
	N := len(A)
	X := make([]float64, N)
	Q := make([][]float64, N)
	for i := range Q {
		Q[i] = make([]float64, N)
	}
	for i := range Q { // Q = I
		Q[i][i] = 1
	}

	for i := 0; i < N-1; i++ {
		a := make([]float64, N-i)  // i-ый вектор-столбец элементов от i до N
		ai := make([]float64, N-i) // вектор-столбец, который хотим получить на месте i-ого, ai[0] = euclideanNorm(a), остальные - 0

		for j := 0; j < N-i; j++ {
			a[j] = A[i+j][i]
		}

		ai[0] = tools.EuclideanNorm(a)

		w := tools.SubtractVectors(a, ai) // искомый вектор нормали преобразования w
		norm := tools.EuclideanNorm(w)
		if norm != 0 {
			for i, value := range w {
				w[i] = value / norm
			}
		}

		// Замена i-ого столбца
		for j := i; j < N; j++ {
			A[j][i] = ai[j-i]
		}

		// Вычисление столбоцов от i+1 до N, в каждом меняются элементы с от i+1 до N
		for j := i + 1; j < N; j++ {
			for k := 0; k < N-i; k++ {
				ai[k] = A[k+i][j]
			}

			pr := tools.DotProduct(w, ai)
			for k := i; k < N; k++ {
				A[k][j] = A[k][j] - 2*w[k-i]*pr
			}
		}

		/////////////////////////////////////////////////////////////////////

		// Вычисление изменненого вектора свободных членов
		pr := tools.DotProduct(w, B[i:])
		for j := i; j < N; j++ {
			B[j] = B[j] - 2*w[j-i]*pr
		}

		/////////////////////////////////////////////////////////////////////

		// Вычислние матрицы преобразования
		Qi := make([][]float64, N-i)
		for j := range Qi {
			Qi[j] = make([]float64, N-i)
		}
		for j := range Qi { // Qi = I
			Qi[j][j] = 1
		}
		for j := 0; j < N-i; j++ {
			for k := 0; k < N-i; k++ {
				Qi[j][k] = Qi[j][k] - 2*w[j]*w[k] // I - 2wwT
			}
		}

		QPart := make([][]float64, N) // Подматрица матрицы Q Nx(N-i) последних N-i столбцов
		for j := range QPart {
			for k := 0; k < N-i; k++ {
				QPart[j] = append(QPart[j], Q[j][k+i])
			}
		}

		Qi = tools.MultiplyMatrices(QPart, Qi)

		for j := range Q { // Замена последних N-i столбцов на  QPart * Qi
			for k := 0; k < N-i; k++ {
				Q[j][k+i] = Qi[j][k]
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

	// R = A
	R := make([][]float64, N)
	for i := range R {
		for j := range R {
			R[i] = append(R[i], A[i][j])
		}
	}

	return X, R, Q
}
