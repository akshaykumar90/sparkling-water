// problem 17.8

package chapter17

import "github.com/akshaykumar90/sparkling-water/chapter6"

func sudokuSolveHelper(arr [][]int, i, j int) bool {
	if j == len(arr[0]) {
		i, j = i+1, 0
	}
	if i == len(arr) {
		return true
	}
	if arr[i][j] > 0 {
		return sudokuSolveHelper(arr, i, j+1)
	} else {
		for k := 0; k < 9; k++ {
			arr[i][j] = k + 1
			if chapter6.SudokuChecker(arr) && sudokuSolveHelper(arr, i, j+1) {
				return true
			}
		}
		arr[i][j] = 0
		return false
	}
}

func SudokuSolve(arr [][]int) bool {
	return chapter6.SudokuChecker(arr) && sudokuSolveHelper(arr, 0, 0)
}
