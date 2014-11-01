// Problem 9.3

package chapter9

func symmetricBinaryTreeHelper(left, right *TreeElement) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil {
		return left.Value == right.Value &&
			symmetricBinaryTreeHelper(left.Left, right.Right) &&
			symmetricBinaryTreeHelper(left.Right, right.Left)
	} else {
		return false
	}
}

func SymmetricBinaryTree(root *TreeElement) bool {
	return root == nil || symmetricBinaryTreeHelper(root.Left, root.Right)
}
