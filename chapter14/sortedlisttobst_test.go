package chapter14

import "testing"

func TestSortedListToBST(t *testing.T) {
	head := &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{5, nil, &TreeNode{7, nil, &TreeNode{11, nil, nil}}}}}

	root := SortedListToBST(head, 5)

	if !IsBinaryTreeABST(root) {
		t.Fatalf("SortedListToBST: Not a BST")
	}
}
