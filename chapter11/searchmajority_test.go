package chapter11

import "testing"

var searchMajorityTests = []struct {
	arr      []int
	expected int
}{
	{[]int{1, 2, 1, 2, 1, 2, 2, 2}, 2},
	{[]int{1, 1, 1, 1, 2, 2, 2}, 1},
	{[]int{1, 1, 1, 2, 2, 2, 2}, 2},
	{[]int{1, 4, 3, 4, 4, 3, 1, 1, 4, 4, 4}, 4},
	{[]int{1, 1, 1, 1, 1, 1, 1}, 1},
	{[]int{1, 2, 1, 2, 1, 1, 1, 1, 1}, 1},
	{[]int{3}, 3},
}

func TestSearchMajority(t *testing.T) {
	for _, tt := range searchMajorityTests {
		actual := SearchMajority(tt.arr)
		if actual != tt.expected {
			t.Errorf("SearchMajority(%v): expected %d, actual %d",
				tt.arr, tt.expected, actual)
		}
	}
}
