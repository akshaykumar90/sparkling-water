package chapter11

import "testing"

func TestBinarySearchFirstK(t *testing.T) {
	arr := []int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}

	var binarySearchFirstKTests = []struct {
		k        int
		expected int
	}{
		{108, 3},
		{285, 6},
		{300, -1},
	}

	for _, tt := range binarySearchFirstKTests {
		actual := BinarySearchFirstK(arr, tt.k)
		if actual != tt.expected {
			t.Errorf("BinarySearchFirstK(%v, %v): expected %d, actual %d",
				arr, tt.k, tt.expected, actual)
		}
	}
}
