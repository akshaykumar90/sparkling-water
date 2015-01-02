package chapter10

import (
	"fmt"
	"testing"
)

func TestApproximateSort(t *testing.T) {
	in := []int{1, 4, 3, 6, 2, 7, 8, 5, 9, 10}

	out := ApproximateSort(in, 3)

	fmt.Printf("ApproximateSort (%v , 3) = %v\n", in, out)
	fmt.Println()

}
