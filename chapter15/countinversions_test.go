package chapter15

import (
	"math/rand"
	"testing"
	"time"
)

func TestCountInversions(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	arr := rand.Perm(100)

	var expected int

	// brute-force computation
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				expected++
			}
		}
	}

	actual := CountInversions(arr)

	if expected != actual {
		t.Errorf("CountInversions: expected %d, actual %d", expected, actual)
	}
}
