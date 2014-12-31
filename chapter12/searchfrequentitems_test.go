package chapter12

import (
	"sort"
	"testing"
)

func TestSearchFrequentItems(t *testing.T) {
	arr := []int{
		7, 2, 4, 4, 2, 9, 10, 7, 7, 2,
		2, 4, 2, 2, 10, 4, 7, 7, 4, 10,
		5, 2, 8, 9, 2,
	}

	expected := []int{2, 4, 7}

	actual := SearchFrequentItems(arr, 5)

	sort.Ints(actual)

	if len(expected) != len(actual) {
		t.Fatalf("SearchFrequentItems: expected length %d, actual length %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Fatalf("SearchFrequentItems: at %d: expected %d, actual %d", i, v, actual[i])
		}
	}
}
