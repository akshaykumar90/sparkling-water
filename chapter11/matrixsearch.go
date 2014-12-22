// Problem 11.10

package chapter11

func MatrixSearch(arr [][]int, x int) bool {
	n := len(arr)
	i, j := 0, n-1

	for i < n && j >= 0 {
		if e := arr[i][j]; x == e {
			return true
		} else if x < e {
			j--
		} else { // x > e
			i++
		}
	}

	return false
}
