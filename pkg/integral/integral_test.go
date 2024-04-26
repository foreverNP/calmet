package integral

import (
	"math"
	"testing"
)

const (
	e = 1e-7
	a = 0
	b = math.Pi / 4.0
)

var (
	f = func(x float64) float64 {
		return math.Pow((math.Sin(x)-math.Cos(x))/(math.Sin(x)+math.Cos(x)), 3.0)
	}

	I = math.Log(math.Sqrt2) - 0.5
)

func TestIntegrateGaussLegendre(t *testing.T) {
	resultGaussLegendre := IntegrateGaussLegendre(f, a, b, 4)

	if math.Abs(I-resultGaussLegendre) > e*1000 {
		t.Errorf("expected: %v, got: %v", I, resultGaussLegendre)
	}
}

func TestIntegrateSimpson(t *testing.T) {
	resultSimpson := IntegrateSimpson(f, a, b, e, nil)

	if math.Abs(I-resultSimpson) > e {
		t.Errorf("expected: %v, got: %v", I, resultSimpson)
	}
}

func TestIntegrateTrapezoidal(t *testing.T) {
	resultTrapezoid := IntegrateTrapezoidal(f, a, b, e, nil)

	if math.Abs(I-resultTrapezoid) > e {
		t.Errorf("expected: %v, got: %v", I, resultTrapezoid)
	}
}
