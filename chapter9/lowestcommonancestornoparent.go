// Problem 9.12

package chapter9

func LowestCommonAncestorNoParent(root, a, b *TreeElement) *TreeElement {
	if root == nil {
		return nil
	} else if root == a || root == b {
		return root
	}

	left, right := LowestCommonAncestorNoParent(root.Left, a, b), LowestCommonAncestorNoParent(root.Right, a, b)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}
}
