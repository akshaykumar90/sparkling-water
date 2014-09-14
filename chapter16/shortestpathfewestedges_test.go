package chapter16

import "testing"

// Figure 16.1 - A directed graph with weights on edges.
var graph = [][]edgenode{
	{{1, 3}, {2, 2}},
	{{0, 4}, {10, 1}},
	{{4, 8}},
	{{2, 5}, {7, 5}},
	{{3, 7}},
	{{6, 6}},
	{{5, 7}, {7, 4}},
	{},
	{{9, 6}},
	{{11, 7}, {5, 1}},
	{{8, 1}},
	{{8, 9}},
	{{13, 5}},
	{{12, 12}},
}

var shortestPathFewestEdgesTests = []struct {
	start    int
	end      int
	expected []int
}{
	{0, 7, []int{0, 2, 4, 3, 7}},
}

func TestShortestPathFewestEdges(t *testing.T) {
	for _, tt := range shortestPathFewestEdgesTests {
		actual := ShortestPathFewestEdges(graph, tt.start, tt.end)

		if len(tt.expected) != len(actual) {
			t.Errorf("ShortestPathFewestEdges(%d,%d): expected length %d, actual length %d",
				tt.start, tt.end, len(tt.expected), len(actual))
		} else {
			for i, v := range tt.expected {
				if v != actual[i] {
					t.Fatalf("ShortestPathFewestEdges(%d,%d): at %d: expected %d, actual %d",
						tt.start, tt.end, i+1, v, actual[i])
				}
			}
		}
	}
}
