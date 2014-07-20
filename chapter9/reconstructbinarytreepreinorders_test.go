package chapter9

import (
	"fmt"
	"testing"
)

func TestReconstructBinaryTreePreInOrders(t *testing.T) {
	var inOrderTraversal func(root *TreeNode)
	inOrderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		} else {
			inOrderTraversal(root.Left)
			fmt.Println(root.Value)
			inOrderTraversal(root.Right)
		}
	}

	in := []string{"F", "B", "A", "E", "H", "C", "D", "I", "G"}
	pre := []string{"H", "B", "F", "E", "A", "C", "D", "G", "I"}

	root := ReconstructBinaryTreePreInOrders(pre, in)

	inOrderTraversal(root)
	fmt.Println()
}
