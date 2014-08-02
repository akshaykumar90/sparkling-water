package chapter14

import "testing"

func TestIsBinaryTreeABSTValid(t *testing.T) {
	tree := TreeNode{19,
		&TreeNode{7,
			&TreeNode{Value: 3},
			&TreeNode{Value: 11},
		},
		&TreeNode{43,
			&TreeNode{Value: 23},
			nil,
		},
	}

	actual := IsBinaryTreeABST(&tree)
	if actual != true {
		t.Errorf("IsBinaryTreeABST(%v): expected %t, actual %t",
			tree, true, actual)
	}
}

func TestIsEmptyBinaryTreeABST(t *testing.T) {
	actual := IsBinaryTreeABST(nil)
	if actual != true {
		t.Errorf("IsBinaryTreeABST(nil): expected %t, actual %t",
			true, actual)
	}
}

func TestIsBinaryTreeABSTInvalid(t *testing.T) {
	tree := TreeNode{19,
		&TreeNode{7,
			&TreeNode{Value: 3},
			&TreeNode{Value: 22},
		},
		&TreeNode{43,
			&TreeNode{Value: 12},
			nil,
		},
	}

	actual := IsBinaryTreeABST(&tree)
	if actual != false {
		t.Errorf("IsBinaryTreeABST(%v): expected %t, actual %t",
			tree, false, actual)
	}
}
