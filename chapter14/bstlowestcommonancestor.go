// Problem 14.12

package chapter14

func BSTLowestCommonAncestor(root *TreeNode, s, b int) int {
	if root.Value < s {
		return BSTLowestCommonAncestor(root.Right, s, b)
	} else if root.Value > b {
		return BSTLowestCommonAncestor(root.Left, s, b)
	} else {
		return root.Value
	}
}
