package spline

import (
	"github.com/foreverNP/calmet/pkg/node"
	"math"
	"testing"
)

const (
	a      = -3.0
	b      = 3.0
	N      = 15
	points = 100
	e      = 1e-4
)

var (
	f = func(x float64) float64 {
		return math.Sin(x)
	}

	Df = func(x float64) float64 {
		return math.Cos(x)
	}
)

func TestCubicSpline_Solve(t *testing.T) {
	spl := New(Df, node.BuildEquidistantNodes(f, a, b, N))

	step := (b - a) / points
	interErr := 0.0

	for i := 0; i <= points; i++ {
		x := a + float64(i)*step
		y, err := spl.Solve(x)

		if err != nil {
			t.Errorf("expected: %v, got: %v", nil, err)
		}

		yReal := f(x)

		interErr = math.Max(interErr, math.Abs(yReal-y))
	}

	if interErr > e {
		t.Errorf("unexpected error: %v", interErr)
	}
}
