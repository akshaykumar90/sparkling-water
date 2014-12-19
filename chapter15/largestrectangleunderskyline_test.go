package chapter15

import "testing"

var largestRectangleUnderSkylineTests = []struct {
	hs       []int
	expected int
}{
	{[]int{4, 5, 6}, 12},
	{[]int{6, 5, 4}, 12},
	{[]int{7, 5, 6}, 15},
	{[]int{1, 3, 3, 3, 4, 4, 4, 3, 6, 3, 3, 2, 2, 1}, 30},
}

func TestLargestRectangleUnderSkyline(t *testing.T) {
	for _, tt := range largestRectangleUnderSkylineTests {
		actual := LargestRectangleUnderSkyline(tt.hs)
		if actual != tt.expected {
			t.Errorf("LargestRectangleUnderSkyline(%v): expected %d, actual %d",
				tt.hs, tt.expected, actual)
		}
	}
}
