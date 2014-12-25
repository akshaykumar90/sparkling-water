package chapter5

import (
	"fmt"
	"testing"
)

func TestPowerSet(t *testing.T) {
	var powerSetTests = [][]int{
		{},
		{1},
		{1, 2, 3},
	}

	for _, tt := range powerSetTests {
		fmt.Printf("PowerSet(%v):\n", tt)
		PowerSet(tt)
	}
}
