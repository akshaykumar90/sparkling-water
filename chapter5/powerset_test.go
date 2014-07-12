package chapter5

import "testing"

func TestPowerSet(t *testing.T) {
	t.Skip("Skipping PowerSet tests...")

	var powerSetTests = [][]int{{}, {1}, {1, 2, 3}}

	for _, tt := range powerSetTests {
		PowerSet(tt)
	}
}
