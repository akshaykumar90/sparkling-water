package chapter9

import (
	"fmt"
	"testing"
)

func inOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	} else {
		inOrderTraversal(root.Left)
		fmt.Println(root.Value)
		inOrderTraversal(root.Right)
	}
}

func TestReconstructBinaryTreePreInOrders(t *testing.T) {
	in := []string{"F", "B", "A", "E", "H", "C", "D", "I", "G"}
	pre := []string{"H", "B", "F", "E", "A", "C", "D", "G", "I"}

	root := ReconstructBinaryTreePreInOrders(pre, in)

	inOrderTraversal(root)
	fmt.Println()
}
