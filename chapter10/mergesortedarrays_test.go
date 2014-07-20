package chapter10

import (
	"fmt"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	arrs := [][]int{
		{24, 45, 49},
		{12, 59, 87},
		{44, 99},
	}

	result := MergeSortedArrays(arrs)
	fmt.Println(result)
	fmt.Println()

}
