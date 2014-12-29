// Problem 14.4

package chapter14

func SearchBSTFirstLargerK(root *TreeNode, k int) *TreeNode {
	var first *TreeNode
	found := false

	for root != nil {
		if root.Value > k {
			first = root
			root = root.Left
		} else {
			if root.Value == k {
				found = true
			}
			root = root.Right
		}
	}

	if found {
		return first
	} else {
		return nil
	}
}
