package chapter5

import (
	"math/rand"
	"testing"
	"time"
)

func countBitsSetTo1(x uint64) int {
	cnt := 0
	for ; x > 0; x &= (x - 1) {
		cnt++
	}
	return cnt
}

func TestClosestIntSameBits(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	x := uint64(rand.Int63())

	res := ClosestIntSameBits(x)

	t.Logf("ClosestIntSameBits(%d): %d", x, res)

	if countBitsSetTo1(res) != countBitsSetTo1(x) {
		t.Errorf("ClosestIntSameBits(%d): got %d, bit count: expected %d, actual %d",
			x, res, countBitsSetTo1(x), countBitsSetTo1(res))
	}
}
