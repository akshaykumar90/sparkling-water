// Problem 14.7

package chapter14

func BSTToSortedDoublyList(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	} else {
		lf, lb := BSTToSortedDoublyList(root.Left)
		rf, rb := BSTToSortedDoublyList(root.Right)

		var f, b *TreeNode

		if lb != nil {
			lb.Right = root
			root.Left = lb
			f = lf
		} else {
			f = root
		}

		if rf != nil {
			rf.Left = root
			root.Right = rf
			b = rb
		} else {
			b = root
		}

		return f, b
	}
}
