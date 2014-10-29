// Problem 14.5

package chapter14

func SearchMinFirstBST(root *TreeNode, k int) bool {
	if root != nil {
		if k == root.Value {
			return true
		} else if k < root.Value {
			return false
		} else if root.Right != nil && k >= root.Right.Value {
			return SearchMinFirstBST(root.Right, k)
		} else {
			return SearchMinFirstBST(root.Left, k)
		}
	} else {
		return false
	}
}
