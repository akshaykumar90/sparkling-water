package chapter14

import "testing"

// Figure 14.3 - A min-first BST.
var minFirstTree = &TreeNode{2,
	&TreeNode{3,
		nil,
		&TreeNode{5,
			&TreeNode{Value: 7},
			&TreeNode{Value: 11},
		},
	},
	&TreeNode{13,
		&TreeNode{Value: 17},
		&TreeNode{19,
			&TreeNode{Value: 23},
			nil,
		},
	},
}

var searchMinFirstBSTTests = []struct {
	k        int
	expected bool
}{
	{2, true},
	{11, true},
	{13, true},
	{23, true},
	{3, true},
	{1, false},
	{4, false},
	{20, false},
}

func TestSearchMinFirstBST(t *testing.T) {
	for _, tt := range searchMinFirstBSTTests {
		actual := SearchMinFirstBST(minFirstTree, tt.k)
		if actual != tt.expected {
			t.Errorf("SearchMinFirstBST(%d): expected %t, actual %t", tt.k, tt.expected, actual)
		}
	}
}
