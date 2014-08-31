package chapter11

import (
	"math"
	"testing"
)

func TestSquareRoot(t *testing.T) {
	tol := 0.0000001

	ns := []float64{0.3, 0.9, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}

	for _, x := range ns {
		actual := SquareRoot(x)
		if math.Abs(math.Sqrt(x)-actual) > tol {
			t.Errorf("SquareRoot(%.1f): expected %.3f, actual %.3f",
				x, math.Sqrt(x), actual)
		}
	}
}
