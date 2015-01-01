// Problem 15.9

package chapter15

import "math"

type MaxPair struct {
	w, h int
}

func min(a, b int) int {
	if b < a {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if b > a {
		return b
	} else {
		return a
	}
}

func preProcessMatrix(mat [][]bool) [][]MaxPair {
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

	return maxPairMat
}

func MaxSubmatrixRectangle(mat [][]bool) int {
	m, n := len(mat), len(mat[0])

	maxPairMat := preProcessMatrix(mat)

	largest := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] && maxPairMat[i][j].w*maxPairMat[i][j].h > largest {
				minw := math.MaxInt32
				for k := 0; k < maxPairMat[i][j].h; k++ {
					minw = min(minw, maxPairMat[i+k][j].w)
					largest = max(largest, (k+1)*minw)
				}
			}
		}
	}

	return largest
}

func MaxSubmatrixSquare(mat [][]bool) int {
	m, n := len(mat), len(mat[0])

	maxPairMat := preProcessMatrix(mat)

	largest := 0

	longestSide := make([][]int, m)
	for i := range longestSide {
		longestSide[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] {
				s := min(maxPairMat[i][j].h, maxPairMat[i][j].w)
				if i+1 < m && j+1 < n {
					s = min(longestSide[i+1][j+1]+1, s)
				}
				longestSide[i][j] = s
				largest = max(largest, s*s)
			}
		}
	}

	return largest
}
