package equations

import "github.com/foreverNP/calmet/pkg/tools"

const (
	Kmax = 1000000 // максимальное количество итераций
)

// RelaxationMethod решает СЛАУ методом релаксации
// (при w == 1 превращается в метод Гаусса – Зейделя)
// A - матрица коэффициентов, B - вектор свободных членов, w - весовой коэффициент, e - точность
// Возвращает вектор решений и количество итераций
func RelaxationMethod(A [][]float64, B []float64, w float64, e float64) ([]float64, int) {
	if w <= 0 || w >= 2 {
		panic("Неверный весовой коэффициент!")
	}

	X1 := make([]float64, len(B))
	X2 := make([]float64, len(B))

	copy(X1, B)
	copy(X2, B)

	K := 0
	for ; K < Kmax; K++ {
		for i := 0; i < len(B); i++ {
			sum := 0.0
			for j := 0; j < len(B); j++ {
				if i != j {
					sum += A[i][j] * X2[j]
				}
			}
			X2[i] = (1.0-w)*X2[i] + (w/A[i][i])*(B[i]-sum)
		}

		if tools.MaxAbsoluteDifference(X2, X1) < e {
			K++
			break
		}

		copy(X1, X2)
	}

	return X2, K
}
