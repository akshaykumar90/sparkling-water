package chapter6

func SudokuChecker(arr [][]int) bool {
	rows := make([][]bool, 9)
	cols := make([][]bool, 9)
	subarray := make([][]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 9)
		cols[i] = make([]bool, 9)
		subarray[i] = make([]bool, 9)
	}

	for j := range arr {
		for k, n := range arr[j] {
			if n == 0 {
				continue
			}

			switch {
			case rows[j][n-1] == true:
				return false
			case cols[k][n-1] == true:
				return false
			case subarray[j/3*3+k/3][n-1] == true:
				return false
			default:
				rows[j][n-1] = true
				cols[k][n-1] = true
				subarray[j/3*3+k/3][n-1] = true
			}
		}
	}
	return true
}

