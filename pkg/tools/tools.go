package tools

import "math"

// DotProduct Скалярное произведение двух векторов
func DotProduct(vector1, vector2 []float64) float64 {
	if len(vector1) != len(vector2) {
		panic("Длины векторов должны совпадать")
	}

	result := 0.0
	for i := 0; i < len(vector1); i++ {
		result += vector1[i] * vector2[i]
	}

	return result
}

// EuclideanNorm Евклидова норма вектора
func EuclideanNorm(vector []float64) float64 {
	result := 0.0
	for i := 0; i < len(vector); i++ {
		result += vector[i] * vector[i]
	}

	return math.Sqrt(result)
}

// SubtractVectors Вычитание двух векторов
func SubtractVectors(vector1, vector2 []float64) []float64 {
	if len(vector1) != len(vector2) {
		panic("Длины векторов должны совпадать")
	}

	result := make([]float64, len(vector1))

	for i := 0; i < len(vector1); i++ {
		result[i] = vector1[i] - vector2[i]
	}

	return result
}

// MultiplyMatrices умножение матриц
func MultiplyMatrices(matrix1, matrix2 [][]float64) [][]float64 {
	rows1, cols1 := len(matrix1), len(matrix1[0])
	rows2, cols2 := len(matrix2), len(matrix2[0])

	// Проверяем, можно ли умножить матрицы
	if cols1 != rows2 {
		panic("Невозможно умножить матрицы. Количество столбцов первой матрицы не равно количеству строк второй матрицы.")
	}

	// Создаем новую матрицу результатов
	result := make([][]float64, rows1)
	for i := range result {
		result[i] = make([]float64, cols2)
	}

	// Вычисляем произведение
	for i := 0; i < rows1; i++ {
		for j := 0; j < cols2; j++ {
			for k := 0; k < cols1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result
}

// MatrixNorm
// кубическая/строковая норма матрицы
func MatrixNorm(matrix [][]float64) float64 {
	var maxSum float64

	for i := range matrix {
		var sum float64
		for j := range matrix[0] {
			sum += math.Abs(matrix[i][j])
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

// IsDiagonallyDominant проверка на диагональное преобладание
func IsDiagonallyDominant(matrix [][]float64) bool {
	for i := 0; i < len(matrix); i++ {
		sum := 0.0

		for j := 0; j < len(matrix[0]); j++ {
			if i != j {
				sum += math.Abs(matrix[i][j])
			}
		}

		if matrix[i][i] <= sum {
			return false
		}
	}

	return true
}

// MaxAbsoluteDifference маскимальная по модулю разность элементво вектора
func MaxAbsoluteDifference(vec1, vec2 []float64) float64 {
	return UniformNorm(SubtractVectors(vec1, vec2))
}

// UniformNorm Равномерная норма вектора
func UniformNorm(vector []float64) float64 {
	maxElem := math.Abs(vector[0])
	for i := 1; i < len(vector); i++ {
		absValue := math.Abs(vector[i])
		if absValue > maxElem {
			maxElem = absValue
		}
	}
	return maxElem
}

// MultiplyMatrixByScalar умножение матриц на скаляр
func MultiplyMatrixByScalar(matrix [][]float64, scalar float64) [][]float64 {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = matrix[i][j] * scalar
		}
	}

	return result
}

// TransposeMatrix транспонирует заданную матрицу.
func TransposeMatrix(matrix [][]float64) [][]float64 {
	// Определяем количество строк и столбцов в исходной матрице.
	rows := len(matrix)
	cols := len(matrix[0])

	// Создаем новую матрицу с перевернутыми размерами (количество строк становится количеством столбцов и наоборот).
	result := make([][]float64, cols)

	// Инициализируем новую матрицу.
	for i := 0; i < cols; i++ {
		result[i] = make([]float64, rows)
	}

	// Заполняем новую матрицу элементами, транспонированными из исходной матрицы.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

// FindMaxOffDiagonalElement находит максимальный по модулю элемент вне главной диагонали матрицы.
func FindMaxOffDiagonalElement(A [][]float64) (int, int) {
	maxVal := 0.0
	p, q := 0, 0

	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			if math.Abs(A[i][j]) > maxVal {
				maxVal = math.Abs(A[i][j])
				p = i
				q = j
			}
		}
	}

	return p, q
}

// Off возвращает сумму квадратов элементов вне главной диагонали матрицы.
func Off(A [][]float64) float64 {
	maxVal := 0.0

	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			maxVal += 2 * A[i][j] * A[i][j]
		}
	}

	return maxVal
}
