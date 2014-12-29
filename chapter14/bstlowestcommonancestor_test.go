package chapter14

import "testing"

var BSTLowestCommonAncestorTests = []struct {
	s        int
	b        int
	expected int
}{
	{3, 17, 7},
	{11, 37, 19},
	{19, 31, 19},
	{11, 13, 11},
	{23, 43, 43},
	{2, 53, 19},
}

func TestBSTLowestCommonAncestor(t *testing.T) {
	for _, tt := range BSTLowestCommonAncestorTests {
		actual := BSTLowestCommonAncestor(tree, tt.s, tt.b)
		if actual != tt.expected {
			t.Errorf("BSTLowestCommonAncestor(%d, %d): expected %d, actual %d",
				tt.s, tt.b, tt.expected, actual)
		}
	}
}
