// problem 11.9

package chapter11

import "math"

func SquareRoot(x float64) float64 {
	tol := 0.0000001
	var lo, hi float64
	if x < 1.0 {
		lo, hi = x, 1.0
	} else {
		lo, hi = 1.0, x
	}
	max_i := 100
	var cand float64
	for i := 0; i < max_i; i++ {
		cand = (lo + hi) / 2
		sq := cand * cand
		if math.Abs(sq-x) < tol {
			return cand
		} else if sq > x {
			hi = cand
		} else {
			lo = cand
		}
	}
	return cand
}
