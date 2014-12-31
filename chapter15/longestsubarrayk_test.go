package chapter15

import "testing"

var longestSubarrayKTests = []struct {
	arr      []int
	k        int
	expected int
}{
	{[]int{431, -15, 639, 342, -14, 565, -924, 635, 167, -70}, 184, 4},
}

func TestLongestSubarrayK(t *testing.T) {
	for _, tt := range longestSubarrayKTests {
		actual := LongestSubarrayK(tt.arr, tt.k)
		if actual != tt.expected {
			t.Errorf("LongestSubarrayK(%v, %d): expected %d, actual %d",
				tt.arr, tt.k, tt.expected, actual)
		}
	}
}
