package chapter11

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestFindMinMax(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	a := rand.Perm(1000)[:25]

	min := math.MaxInt32
	max := math.MinInt32

	// brute-force computation
	for _, x := range a {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	m, M := FindMinMax(a)

	if min != m || max != M {
		t.Errorf("FindMinMax: expected %d %d, actual %d %d", min, max, m, M)
	}
}
