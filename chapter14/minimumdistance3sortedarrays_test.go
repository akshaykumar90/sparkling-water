package chapter14

import (
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMinimumDistance3SortedArrays(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	noise := rand.Perm(1000)

	a := noise[:33]
	b := noise[33:66]
	c := noise[66:99]

	sort.Ints(a)
	sort.Ints(b)
	sort.Ints(c)

	min := math.MaxInt32

	// brute-force computation
	for i := 0; i < 33; i++ {
		for j := 0; j < 33; j++ {
			for k := 0; k < 33; k++ {
				var s, l int
				if a[i] < b[j] {
					s, l = a[i], b[j]
				} else {
					s, l = b[j], a[i]
				}
				if c[k] < s {
					s = c[k]
				} else if c[k] > l {
					l = c[k]
				}
				if min > l-s {
					min = l - s
				}
			}
		}
	}

	actual := MinimumDistance3SortedArrays([][]int{a, b, c})

	if min != actual {
		t.Errorf("MinimumDistance3SortedArrays: expected %d, actual %d", min, actual)
	}
}
