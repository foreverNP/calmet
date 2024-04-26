package equations

import "github.com/foreverNP/calmet/pkg/tools"

// JacobiMethod решает СЛАУ итерационным методом Якоби
// A - матрица коэффициентов, B - вектор свободных членов, e - точность
func JacobiMethod(A [][]float64, B []float64, e float64) ([]float64, int) {
	X1 := make([]float64, len(B))
	X2 := make([]float64, len(B))

	copy(X1, B)

	K := 0
	for ; K < Kmax; K++ {
		for i := 0; i < len(B); i++ {
			sum := 0.0
			for j := 0; j < len(B); j++ {
				if i != j {
					sum += A[i][j] * X1[j]
				}
			}
			X2[i] = (1.0 / A[i][i]) * (B[i] - sum)
		}

		if tools.MaxAbsoluteDifference(X2, X1) < e {
			K++
			break
		}

		copy(X1, X2)
	}

	return X2, K
}
