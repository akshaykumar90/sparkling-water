package chapter11

import "testing"

var matrixSearchTests = []struct {
	arr      [][]int
	x        int
	expected bool
}{
	{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 8, true},
	{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 10, false},
	{[][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}, 3, true},
	{[][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}, 5, false},
}

func TestMatrixSearch(t *testing.T) {
	for _, tt := range matrixSearchTests {
		actual := MatrixSearch(tt.arr, tt.x)
		if actual != tt.expected {
			t.Errorf("MatrixSearch(%v, %d): expected %t, actual %t",
				tt.arr, tt.x, tt.expected, actual)
		}
	}
}
