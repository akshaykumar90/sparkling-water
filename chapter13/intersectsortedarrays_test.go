package chapter13

import (
	"fmt"
	"testing"
)

func TestIntersectSortedArrays(t *testing.T) {
	a := []int{1, 3, 4, 6, 8, 10}
	b := []int{3, 4, 5, 9, 10}

	c := IntersectSortedArrays(a, b)

	fmt.Printf("IntersectSortedArrays (%v , %v) = %v\n", a, b, c)
	fmt.Println()

}
