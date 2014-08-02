package chapter14

import (
	"fmt"
	"testing"
)

func TestRebuildBSTPreorder(t *testing.T) {
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

	root := RebuildBSTPreorder([]int{19, 7, 3, 2, 5, 11, 17, 13, 43, 23, 37, 29, 31, 41, 47})

	inOrderTraversal(root)
	fmt.Println()
}
