package chapter10

import "testing"

var testMaxHeap = []int{10, 7, 5, 5, 4, 5, 5, 4, 5, 4}

var compareKthLargestInHeapTests = []struct {
	k, x     int
	expected int
}{
	{5, 5, 0},
	{5, 4, 1},
	{5, 10, -1},
	{1, 10, 0},
	{1, 5, 1},
	{8, 4, 0},
	{8, 10, -1},
	{8, 0, 1},
	{10, 4, 0},
	{10, 3, 1},
	{10, 7, -1},
}

func TestCompareKthLargestInHeap(t *testing.T) {
	for _, tt := range compareKthLargestInHeapTests {
		actual := CompareKthLargestInHeap(testMaxHeap, tt.k, tt.x)
		if actual != tt.expected {
			t.Errorf("CompareKthLargestInHeap(%d,%d): expected %d, actual %d",
				tt.k, tt.x, tt.expected, actual)
		}
	}
}
