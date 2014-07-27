package chapter11

import "testing"

func TestBinarySearchLargerK(t *testing.T) {
	var binarySearchLargerKTests = []struct {
		a        []int
		k        int
		expected int
	}{
		{[]int{1}, 1, -1},
		{[]int{2}, 1, 0},
		{[]int{1, 1, 1, 1, 1, 1}, 1, -1},
		{[]int{1, 2, 3}, 1, 1},
		{[]int{0, 0, 1, 1, 1, 2, 2, 2}, 1, 5},
		{[]int{1, 1, 1, 2}, 1, 3},
		{[]int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}, 500, -1},
		{[]int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}, 101, 3},
	}

	for _, tt := range binarySearchLargerKTests {
		actual := BinarySearchLargerK(tt.a, tt.k)
		if actual != tt.expected {
			t.Errorf("BinarySearchLargerK(%v, %d): expected %d, actual %d",
				tt.a, tt.k, tt.expected, actual)
		}
	}
}
