// problem 15.10

package chapter15

type Position struct {
	y, x int
}

func StringInMatrix(mat [][]int, xs []int) bool {
	cand := make(map[Position]bool)

	maxY, maxX := len(mat), len(mat[0])

	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if mat[i][j] == xs[0] {
				cand[Position{i, j}] = true
			}
		}
	}

	for _, a := range xs[1:] {
		temp := make(map[Position]bool)
		for pos := range cand {
			y, x := pos.y, pos.x
			if y-1 >= 0 && mat[y-1][x] == a {
				temp[Position{y - 1, x}] = true
			}
			if x+1 < maxX && mat[y][x+1] == a {
				temp[Position{y, x + 1}] = true
			}
			if y+1 < maxY && mat[y+1][x] == a {
				temp[Position{y + 1, x}] = true
			}
			if x-1 >= 0 && mat[y][x-1] == a {
				temp[Position{y, x - 1}] = true
			}
		}
		cand = temp
	}

	return len(cand) > 0
}
