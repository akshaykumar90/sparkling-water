// Problem 9.1

package chapter9

func balancedBinaryTreeHelper(elem *TreeElement) (int, bool) {
	if elem == nil {
		return -1, true
	}

	leftHeight, leftBalanced := balancedBinaryTreeHelper(elem.Left)
	if !leftBalanced {
		return -2, false
	}

	rightHeight, rightBalanced := balancedBinaryTreeHelper(elem.Right)
	if !rightBalanced {
		return -2, false
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		} else {
			return x
		}
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -2, false
	}

	max := func(x, y int) int {
		if x < y {
			return y
		} else {
			return x
		}
	}

	return 1 + max(leftHeight, rightHeight), true
}

func BalancedBinaryTree(root *TreeElement) bool {
	_, balanced := balancedBinaryTreeHelper(root)
	return balanced
}
