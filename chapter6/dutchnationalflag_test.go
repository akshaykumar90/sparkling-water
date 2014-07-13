package chapter6

import (
	"fmt"
	"testing"
)

func TestDutchNationalFlag(t *testing.T) {
	var dutchNationalFlagTests = []struct {
		a []int
		p int
	}{
		{[]int{1}, 0},
		{[]int{1, 2, 3}, 1},
		{[]int{5, 3, 9, 3, 8, 1}, 1},
	}

	for _, tt := range dutchNationalFlagTests {
		fmt.Println(tt.a)
		DutchNationalFlag(tt.a, tt.p)
		fmt.Println(tt.a)
		fmt.Println()
	}
}
