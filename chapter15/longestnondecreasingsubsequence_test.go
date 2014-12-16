package chapter15

import "testing"

var longestNondecreasingSubsequenceTests = []struct {
	arr      []int
	expected int
}{
	{[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9}, 4},
	{[]int{1, 2, 3, 4}, 4},
	{[]int{2, 3, 1, 6}, 3},
	{[]int{8, 3, 1}, 1},
}

func TestLongestNondecreasingSubsequence(t *testing.T) {
	for _, tt := range longestNondecreasingSubsequenceTests {
		actual := LongestNondecreasingSubsequence(tt.arr)
		if actual != tt.expected {
			t.Errorf("LongestNondecreasingSubsequence(%v): expected %d, actual %d",
				tt.arr, tt.expected, actual)
		}
	}
}
