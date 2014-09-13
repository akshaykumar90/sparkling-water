package chapter15

import "testing"

var mat = [][]int{
	{1, 2, 3},
	{3, 4, 5},
	{5, 6, 7},
}

var stringInMatrixTests = []struct {
	xs       []int
	expected bool
}{
	{[]int{1, 3, 4, 6}, true},
	{[]int{1, 2, 3, 4}, false},
}

func TestStringInMatrix(t *testing.T) {
	for _, tt := range stringInMatrixTests {
		actual := StringInMatrix(mat, tt.xs)
		if actual != tt.expected {
			t.Errorf("StringInMatrix(%v): expected %t, actual %t",
				tt.xs, tt.expected, actual)
		}
	}
}
