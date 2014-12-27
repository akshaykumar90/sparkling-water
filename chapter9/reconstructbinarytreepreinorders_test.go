package chapter9

import "testing"

// Figure 9.3: A binary tree - edges that do not terminate in nodes denote empty subtrees.
var ReconstructThisBinaryTree = &TreeNode{"H",
	&TreeNode{"B",
		&TreeNode{"F", nil, nil},
		&TreeNode{"E",
			&TreeNode{"A", nil, nil},
			nil,
		},
	},
	&TreeNode{"C",
		nil,
		&TreeNode{"D",
			nil,
			&TreeNode{"G",
				&TreeNode{"I", nil, nil},
				nil,
			},
		},
	},
}

func assertEqualBinaryTrees(expected, actual *TreeNode) bool {
	if expected == nil && actual == nil {
		return true
	} else if expected != nil && actual != nil {
		return expected.Value == actual.Value &&
			assertEqualBinaryTrees(expected.Left, actual.Left) &&
			assertEqualBinaryTrees(expected.Right, actual.Left)
	} else {
		return false
	}
}

func TestReconstructBinaryTreePreInOrders(t *testing.T) {
	in := []string{"F", "B", "A", "E", "H", "C", "D", "I", "G"}
	pre := []string{"H", "B", "F", "E", "A", "C", "D", "G", "I"}

	actual := ReconstructBinaryTreePreInOrders(pre, in)

	assertEqualBinaryTrees(ReconstructThisBinaryTree, actual)
}
