package chapter14

import (
	"sort"
	"testing"
)

func TestFindKLargestBST(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53}
	sz := len(expected)

	for k := 1; k <= sz; k++ {
		actual := FindKLargestBST(&tree, k)
		if k != len(actual) {
			t.Fatalf("FindKLargestBST(%v, %d): expected length %d, actual length %d",
				tree, k, k, len(actual))
		} else {
			sort.Sort(actual)
			for i, v := range expected[sz-k:] {
				if v != actual[i].Value {
					t.Fatalf("FindKLargestBST(%v, %d): at %d: expected %d, actual %d",
						tree, k, i+1, v, actual[i].Value)
				}
			}
		}
	}
}
