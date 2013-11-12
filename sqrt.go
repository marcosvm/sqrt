package sqrt

import (
	"math"
)

// Calcutates square root of x using
// Newton's approximation up to p precison
// returning the square root and the number of
// iterations.
func Sqrt(x float64, p float64) (float64, int) {
	i := 0
	z := float64(1)
	d := float64(0)
	for math.Abs(z-d) > p {
		i++
		d = z
		z = z - ((z*z - x) / 2 * z)
	}
	return z, i
}
