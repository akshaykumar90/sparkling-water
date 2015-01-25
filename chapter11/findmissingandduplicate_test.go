package chapter11

import "testing"

var findMissingAndDuplicateTests = []struct {
	arr      []int
	mis, dup int
}{
	{[]int{0, 4, 3, 1, 1}, 2, 1},
	{[]int{1, 1}, 0, 1},
	{[]int{5, 3, 6, 2, 4, 3, 1}, 0, 3},
}

func TestFindMissingAndDuplicate(t *testing.T) {
	for _, tt := range findMissingAndDuplicateTests {
		mis, dup := FindMissingAndDuplicate(tt.arr)
		if mis != tt.mis || dup != tt.dup {
			t.Errorf("FindMissingAndDuplicate(%v): expected (%d,%d), actual (%d,%d)",
				tt.arr, tt.mis, tt.dup, mis, dup)
		}
	}
}
