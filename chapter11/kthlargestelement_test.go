package chapter11

import (
	"math/rand"
	"testing"
)

func TestKthLargestElement(t *testing.T) {
	var kthLargestElementTests = []struct {
		size     int
		k        int
		expected int
	}{
		{100, 1, 99},
		{100, 100, 0},
		{100, 50, 50},
	}

	for _, tt := range kthLargestElementTests {
		actual, _ := KthLargestElement(rand.Perm(tt.size), tt.k)
		if actual != tt.expected {
			t.Errorf("KthLargestElement(%d, %d): expected %d, actual %d",
				tt.size, tt.k, tt.expected, actual)
		}
	}
}
