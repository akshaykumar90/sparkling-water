package chapter6

import (
	"math/rand"
	"testing"
	"time"
)

func TestZeroSumSubset(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	for tc := 0; tc < 100; tc++ {
		noise := rand.Perm(1000)

		arr := noise[:100]

		actual := ZeroSumSubset(arr)

		if len(actual) == 0 {
			t.Fatalf("ZeroSumSubset: got empty subset!")
		}

		sum := 0
		for _, v := range actual {
			sum += arr[v]
		}

		if sum%len(arr) != 0 {
			t.Fatalf("ZeroSumSubset: got non-zero sum subset!")
		}
	}
}
