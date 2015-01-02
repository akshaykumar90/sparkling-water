package chapter10

import (
	"fmt"
	"github.com/akshaykumar90/sparkling-water/common"
	"testing"
)

func TestApproximateSort(t *testing.T) {
	in := []int{1, 4, 3, 6, 2, 7, 8, 5, 9, 10}

	actual := ApproximateSort(in, 3)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	handle := fmt.Sprintf("ApproximateSort(%v, 3) -> %v", in, actual)

	common.AssertIntsAreEqual(t, handle, expected, actual)
}
