package integral

import "github.com/foreverNP/calmet/pkg/node"

// IntegrateGaussLegendre вычисляет приближенное значение интеграла функции f от a до b
// с помощью квадратурной формулы Гаусса-Лежандра с n узлами
func IntegrateGaussLegendre(f integrand, a, b float64, n int) float64 {
	nodes := node.BuildGaussLegendreNodes(a, b, n)
	sum := 0.0
	for i := range nodes {
		sum += nodes[i].Y * f(nodes[i].X)
	}

	return sum
}
