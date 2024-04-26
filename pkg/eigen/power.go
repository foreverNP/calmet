package eigen

import "github.com/foreverNP/calmet/pkg/tools"

// PowerMethod Метод степеней для нахождения максимального собственного значения и соответствующего собственного вектора
// A - матрица, для которой ищем собственные значения и вектора, e - точность
// Возвращает собственный вектор, собственное значение, невязку и количество итераций
func PowerMethod(A [][]float64, e float64) ([]float64, float64, float64, int) {
	N := len(A)
	y := make([][]float64, N)
	u := make([][]float64, N)
	counter := 0

	for i := 0; i < N; i++ {
		y[i] = make([]float64, 1)
		u[i] = make([]float64, 1)
	}
	y[0][0] = 1
	u[0][0] = 1

	h := tools.DotProduct(tools.TransposeMatrix(u)[0], tools.TransposeMatrix(tools.MultiplyMatrices(A, u))[0])

	for tools.EuclideanNorm(tools.SubtractVectors(tools.TransposeMatrix(tools.MultiplyMatrices(A, u))[0],
		tools.TransposeMatrix(tools.MultiplyMatrixByScalar(u, h))[0])) > e {
		y = tools.MultiplyMatrices(A, u)
		u = tools.MultiplyMatrixByScalar(y, 1/tools.EuclideanNorm(tools.TransposeMatrix(y)[0]))
		h = tools.DotProduct(tools.TransposeMatrix(u)[0], tools.TransposeMatrix(tools.MultiplyMatrices(A, u))[0])
		counter++
	}

	return tools.TransposeMatrix(u)[0], h,
		tools.EuclideanNorm(tools.SubtractVectors(tools.TransposeMatrix(tools.MultiplyMatrices(A, u))[0],
			tools.TransposeMatrix(tools.MultiplyMatrixByScalar(u, h))[0])), counter
}
