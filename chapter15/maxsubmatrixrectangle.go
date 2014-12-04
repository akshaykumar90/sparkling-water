// Problem 15.9

package chapter15

import "math"

type MaxPair struct {
	w, h int
}

func MaxSubmatrixRectangle(mat [][]bool) int {
	m, n := len(mat), len(mat[0])

	maxPairMat := make([][]MaxPair, m)
	for i := range maxPairMat {
		maxPairMat[i] = make([]MaxPair, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			left := 0
			if j+1 < n {
				left = maxPairMat[i][j+1].w
			}

			down := 0
			if i+1 < m {
				down = maxPairMat[i+1][j].h
			}

			if mat[i][j] {
				maxPairMat[i][j] = MaxPair{left + 1, down + 1}
			}
		}
	}

	largest := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] && maxPairMat[i][j].w*maxPairMat[i][j].h > largest {
				minw := math.MaxInt32
				for k := 0; k < maxPairMat[i][j].h; k++ {
					if maxPairMat[i+k][j].w < minw {
						minw = maxPairMat[i+k][j].w
					}
					area := (k + 1) * minw
					if area > largest {
						largest = area
					}
				}
			}
		}
	}

	return largest
}
