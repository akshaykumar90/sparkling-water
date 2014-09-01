package chapter11

import "testing"

func TestBinarySearchCircularArray(t *testing.T) {
	var binarySearchCircularArrayTests = []struct {
		arr      []int
		expected int
	}{
		{[]int{378, 478, 550, 631, 103, 203, 220, 234, 279, 368}, 4},
	}

	for _, tt := range binarySearchCircularArrayTests {
		actual := BinarySearchCircularArray(tt.arr)
		if actual != tt.expected {
			t.Errorf("BinarySearchCircularArray(%v): expected %d, actual %d",
				tt.arr, tt.expected, actual)
		}
	}
}
