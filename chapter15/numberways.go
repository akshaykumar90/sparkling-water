// Problem 15.16

package chapter15

func NumberWays(n, m int) int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
	}

	mat[0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			up, left := 0, 0
			if i > 0 {
				up = mat[i-1][j]
			}
			if j > 0 {
				left = mat[i][j-1]
			}
			mat[i][j] += up + left
		}
	}

	return mat[n-1][m-1]
}

func NumberWaysObstacles(n, m int, obstacles [][]bool) int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, m)
	}

	if !obstacles[0][0] {
		mat[0][0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !obstacles[i][j] {
				up, left := 0, 0
				if i > 0 {
					up = mat[i-1][j]
				}
				if j > 0 {
					left = mat[i][j-1]
				}
				mat[i][j] += up + left
			}
		}
	}

	return mat[n-1][m-1]
}
