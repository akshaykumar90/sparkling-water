package chapter16

import "testing"

func TestWiringCircuitBoard(t *testing.T) {
	// Figure 16.6 - A set of pins and wires between them.
	board := [][]int{
		{1, 10, 9},
		{2, 10, 0},
		{3, 11, 1},
		{4, 12, 2},
		{5, 3},
		{6, 13, 12, 4},
		{7, 5},
		{8, 13, 6},
		{21, 7},
		{0, 10, 14},
		{1, 11, 15, 9},
		{2, 12, 16, 10},
		{3, 5, 17, 11},
		{5, 7, 21, 19},
		{9, 15},
		{10, 16, 14},
		{11, 17, 15},
		{12, 18, 16},
		{19, 17},
		{13, 20, 18},
		{21, 19},
		{13, 8, 20},
	}

	actual := WiringCircuitBoard(board)
	if actual != true {
		t.Errorf("WiringCircuitBoard: expected %t, actual %t", true, actual)
	}
}
