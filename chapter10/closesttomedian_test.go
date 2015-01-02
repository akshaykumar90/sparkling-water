package chapter10

import (
	"fmt"
	"github.com/akshaykumar90/sparkling-water/common"
	"sort"
	"testing"
)

var closestToMedianTests = []struct {
	a        []int
	k        int
	expected []int
}{
	{
		[]int{7, 14, 10, 12, 2, 11, 29, 3, 4},
		5,
		[]int{7, 10, 11, 12, 14},
	},
	{
		[]int{7, 14, 10, 12, 2},
		5,
		[]int{2, 7, 10, 12, 14},
	},
}

func TestClosestToMedian(t *testing.T) {
	for _, tt := range closestToMedianTests {
		actual := ClosestToMedian(tt.a, tt.k)

		sort.Ints(actual)

		handle := fmt.Sprintf("ClosestToMedian(%v, %d)", tt.a, tt.k)

		common.AssertIntsAreEqual(t, handle, tt.expected, actual)
	}
}
