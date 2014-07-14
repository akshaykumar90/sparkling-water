package chapter6

import "testing"

func TestSudokuChecker(t *testing.T) {
	testSudoku := [][]int{
		[]int{5,3,0, 0,7,0, 0,0,0},
		[]int{6,0,0, 1,9,5, 0,0,0},
		[]int{0,9,8, 0,0,0, 0,6,0},

		[]int{8,0,0, 0,6,0, 0,0,3},
		[]int{4,0,0, 8,0,3, 0,0,1},
		[]int{7,0,0, 0,2,0, 0,0,6},

		[]int{0,6,0, 0,0,0, 2,8,0},
		[]int{0,0,0, 4,1,9, 0,0,5},
		[]int{0,0,0, 0,8,0, 0,7,9},
	}

	if !SudokuChecker(testSudoku) {
		t.Errorf("SudokuChecker: expected true, actual false")
	}
}
