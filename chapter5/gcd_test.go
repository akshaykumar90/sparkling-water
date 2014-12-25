package chapter5

import (
	"math/rand"
	"testing"
	"time"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func TestGCD(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 100; i++ {
		a, b := rand.Intn(1000)+1, rand.Intn(1000)+1
		expected := gcd(a, b)
		actual := GCD(a, b)
		if actual != expected {
			t.Errorf("GCD(%d, %d): expected %d, actual %d", a, b, expected, actual)
		}
	}
}
