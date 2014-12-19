package chapter15

import "testing"

var maximumSubarrayInCircularArrayTests = []struct {
	arr      []int
	expected int
}{
	{[]int{904, 40, 523, 12, -335, -385, -124, 481, -31}, 1929},
}

func TestMaximumSubarrayInCircularArray(t *testing.T) {
	for _, tt := range maximumSubarrayInCircularArrayTests {
		actual := MaximumSubarrayInCircularArray(tt.arr)
		if actual != tt.expected {
			t.Errorf("MaximumSubarrayInCircularArray(%v): expected %d, actual %d",
				tt.arr, tt.expected, actual)
		}
	}
}
