package chapter11

import "testing"

func TestBinarySearchAiEqI(t *testing.T) {
	var binarySearchAiEqITests = []struct {
		a        []int
		expected int
	}{
		{[]int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 401}, 2},
	}

	for _, tt := range binarySearchAiEqITests {
		actual := BinarySearchAiEqI(tt.a)
		if actual != tt.expected {
			t.Errorf("BinarySearchAiEqI(%v): expected %d, actual %d",
				tt.a, tt.expected, actual)
		}
	}
}
