package interpoly

import (
	"github.com/foreverNP/calmet/pkg/node"
	"math"
	"testing"
)

const (
	a = -3.0
	b = 3.0
	n = 30
	e = 1e-7
)

var (
	f1 = func(x float64) float64 {
		return x * x * math.Cos(2*x)
	}
)

func TestNewtonPoly_Solve(t *testing.T) {
	var (
		// P1 Построение полинома Ньютона для f1 с равноотстоящими узлами
		P1 = NewtonPolyBuilder(node.BuildEquidistantNodes(f1, a, b, n))
		// C1 Построение полинома Ньютона для f1 с Чебышевскими узлами
		C1 = NewtonPolyBuilder(node.BuildChebyshevNodes(f1, a, b, n))
	)

	// Проверка значений полиномов в узлах интерполяции
	for i := 0; i < n; i++ {
		x := P1.nodes[i].X
		if math.Abs(P1.Solve(x)-f1(x)) > e {
			t.Errorf("P1: %v != %v", P1.Solve(x), f1(x))
		}
		if math.Abs(C1.Solve(x)-f1(x)) > e {
			t.Errorf("C1: %v != %v", C1.Solve(x), f1(x))
		}
	}
}
