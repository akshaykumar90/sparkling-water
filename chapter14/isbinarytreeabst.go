// problem 14.1

package chapter14

import "math"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func isBinaryTreeWithinLimits(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if root.Value < left || root.Value > right {
		return false
	} else {
		return isBinaryTreeWithinLimits(root.Left, left, root.Value) &&
			isBinaryTreeWithinLimits(root.Right, root.Value, right)
	}
}

func IsBinaryTreeABST(root *TreeNode) bool {
	return isBinaryTreeWithinLimits(root, math.MinInt32, math.MaxInt32)
}
