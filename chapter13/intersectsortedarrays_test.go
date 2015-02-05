package chapter13

import (
	"fmt"
	"github.com/akshaykumar90/sparkling-water/common"
	"testing"
)

func TestIntersectSortedArrays(t *testing.T) {
	a := []int{1, 3, 4, 6, 8, 10, 10}
	b := []int{3, 4, 5, 9, 10, 10}

	expected := []int{3, 4, 10}

	actual := IntersectSortedArrays(a, b)

	handle := fmt.Sprintf("IntersectSortedArrays(%v, %v) -> %v", a, b, actual)

	common.AssertIntsAreEqual(t, handle, expected, actual)
}
