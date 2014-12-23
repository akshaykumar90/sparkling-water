package chapter11

import "testing"

var findElementAppearsOnceTests = []struct {
	arr      []int
	expected int
}{
	{[]int{1, 2, 3, 4, 1, 2, 3, 1, 2, 3}, 4},
	{[]int{1, 1, 1, 2}, 2},
	{[]int{5, 2, 2, 2, 8, 5, 5}, 8},
}

func TestFindElementAppearsOnce(t *testing.T) {
	for _, tt := range findElementAppearsOnceTests {
		actual := FindElementAppearsOnce(tt.arr)
		if actual != tt.expected {
			t.Errorf("FindElementAppearsOnce(%v): expected %d, actual %d",
				tt.arr, tt.expected, actual)
		}
	}
}
