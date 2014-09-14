package chapter16

import (
	"sort"
	"testing"
)

var extendedContactsTests = []struct {
	graph    [][]int
	expected [][]int
}{
	{
		[][]int{{1}, {0, 2}, {1}},
		[][]int{{1, 2}, {0, 2}, {0, 1}},
	},
	{
		[][]int{{}, {0}, {0}, {0}},
		[][]int{{}, {0}, {0}, {0}},
	},
	{
		[][]int{{1}, {0, 2, 3}, {1, 3}, {1, 2}},
		[][]int{{1, 2, 3}, {0, 2, 3}, {0, 1, 3}, {0, 1, 2}},
	},
}

func TestExtendedContacts(t *testing.T) {
	for _, tt := range extendedContactsTests {
		actual := ExtendedContacts(tt.graph)

		for i, ec := range tt.expected {
			if len(ec) != len(actual[i]) {
				t.Errorf("ExtendedContacts(%v): for %d: expected length %d, actual length %d",
					tt.graph, i, len(ec), len(actual[i]))
			} else {
				sort.Ints(actual[i])
				for j, v := range ec {
					if v != actual[i][j] {
						t.Fatalf("ExtendedContacts(%v): for %d: at %d: expected %d, actual %d",
							tt.graph, i, j+1, v, actual[i][j])
					}
				}
			}
		}
	}
}
