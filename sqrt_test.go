package sqrt

import (
	"testing"
  "math"
)

func TestSqrt(t *testing.T) {
	const in, out, precision = 2, float64(1.4142135623730951), 1e-4
	if x, _ := Sqrt(in, precision); math.Abs(out - x) > precision {
		t.Errorf("Sqrt(%v, %v) = %v, wants %v precision, got %v", in, precision, out, precision, math.Abs(out - x))
	}
}
