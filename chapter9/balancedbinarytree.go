// Problem 9.1

package chapter9

import "github.com/akshaykumar90/sparkling-water/common"

func balancedBinaryTreeHelper(elem *TreeElement) int {
	if elem == nil {
		return -1
	}

	leftHeight := balancedBinaryTreeHelper(elem.Left)
	if leftHeight == -2 {
		return -2
	}

	rightHeight := balancedBinaryTreeHelper(elem.Right)
	if rightHeight == -2 {
		return -2
	}

	if common.Abs(leftHeight-rightHeight) > 1 {
		return -2
	}

	return 1 + common.MaxInt(leftHeight, rightHeight)
}

func BalancedBinaryTree(root *TreeElement) bool {
	height := balancedBinaryTreeHelper(root)
	return height != -2
}
