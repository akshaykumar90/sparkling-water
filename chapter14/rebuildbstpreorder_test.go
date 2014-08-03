package chapter14

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

func TestRebuildBSTPreorder(t *testing.T) {
	root := RebuildBSTPreorder([]int{19, 7, 3, 2, 5, 11, 17, 13, 43, 23, 37, 29, 31, 41, 47})

	fmt.Println("TestRebuildBSTPreorder:")
	inOrderTraversal(root)
	fmt.Println()
}

func TestRebuildBSTPreorderBetter(t *testing.T) {
	root := RebuildBSTPreorderBetter([]int{19, 7, 3, 2, 5, 11, 17, 13, 43, 23, 37, 29, 31, 41, 47})

	fmt.Println("TestRebuildBSTPreorderBetter:")
	inOrderTraversal(root)
	fmt.Println()
}
