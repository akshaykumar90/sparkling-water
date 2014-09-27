package chapter10

import "testing"
import "sort"

var closestToMedianTests = []struct {
	a        []int
	k        int
	expected []int
}{
	{
		[]int{7, 14, 10, 12, 2, 11, 29, 3, 4},
		5,
		[]int{7, 10, 11, 12, 14},
	},
	{
		[]int{7, 14, 10, 12, 2},
		5,
		[]int{2, 7, 10, 12, 14},
	},
}

func TestClosestToMedian(t *testing.T) {
	for _, tt := range closestToMedianTests {
		actual := ClosestToMedian(tt.a, tt.k)
		if tt.k != len(actual) {
			t.Fatalf("ClosestToMedian(%v, %d): expected length %d, actual length %d",
				tt.a, tt.k, tt.k, len(actual))
		} else {
			sort.Ints(actual)
			for i, v := range tt.expected {
				if v != actual[i] {
					t.Fatalf("ClosestToMedian(%v, %d): at %d: expected %d, actual %d",
						tt.a, tt.k, i+1, v, actual[i])
				}
			}
		}
	}
}
