package chapter10

import (
	"github.com/akshaykumar90/sparkling-water/common"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	arrs := [][]int{
		{24, 45, 49},
		{12, 59, 87},
		{44, 99},
	}

	expected := []int{12, 24, 44, 45, 49, 59, 87, 99}

	actual := MergeSortedArrays(arrs)

	common.AssertIntsAreEqual(t, "MergeSortedArrays", expected, actual)
}
