package chapter14

import "testing"

func treeHeight(node *TreeNode) int {
	if node == nil {
		return 0
	} else {
		left := treeHeight(node.Left)
		right := treeHeight(node.Right)

		max := 0
		if left < right {
			max = right
		} else {
			max = left
		}

		return max + 1
	}
}

func TestBuildBSTFromSortedArray(t *testing.T) {
	root := BuildBSTFromSortedArray([]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53})

	if !IsBinaryTreeABST(root) {
		t.Fatalf("BuildBSTFromSortedArray: Not a BST!")
	}

	if actual, expected := treeHeight(root), 5; actual > expected {
		t.Fatalf("BuildBSTFromSortedArray: expected height %d, actual height %d", expected, actual)
	}
}
