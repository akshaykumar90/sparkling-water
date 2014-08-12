package chapter6

import "testing"

var spiralMatrixClockwiseTests = []struct {
	arr      [][]int
	expected []int
}{
	{[][]int{{1}}, []int{1}},
	{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},
	{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
}

func TestSpiralMatrixClockwise(t *testing.T) {
	for _, tt := range spiralMatrixClockwiseTests {
		actual := SpiralMatrixClockwise(tt.arr)

		if len(tt.expected) != len(actual) {
			t.Errorf("SpiralMatrixClockwise(%v): expected length %d, actual length %d",
				tt.arr, len(tt.expected), len(actual))
		} else {
			for i, v := range tt.expected {
				if v != actual[i] {
					t.Errorf("SpiralMatrixClockwise(%v): at %d: expected %d, actual %d",
						tt.arr, i+1, v, actual[i])
				}
			}
		}
	}
}
