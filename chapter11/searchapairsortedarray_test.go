package chapter11

import "testing"

func TestSearchAPairSortedArray(t *testing.T) {
	arr := []int{-49, 75, 103, -147, 164, -197, -238, 314, 348, -422}

	searchAPairSortedArrayTests := []struct {
		k int
		i int
		j int
	}{
		{167, 3, 7},
		{239, 1, 4},
		{-619, 5, 9},
		{104, -1, -1},
	}

	for _, tt := range searchAPairSortedArrayTests {
		x, y := SearchAPairSortedArray(arr, tt.k)
		if x != tt.i || y != tt.j {
			t.Errorf("SearchAPairSortedArray(%v, %d): expected (%d,%d), actual (%d,%d)",
				arr, tt.k, tt.i, tt.j, x, y)
		}
	}
}
