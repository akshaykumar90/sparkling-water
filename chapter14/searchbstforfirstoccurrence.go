// problem 14.3

package chapter14

func SearchBSTForFirstOccurrenceRecursive(root *TreeNode, k int) *TreeNode {
	if root != nil {
		if root.Value == k {
			first := SearchBSTForFirstOccurrenceRecursive(root.Left, k)
			if first != nil {
				return first
			} else {
				return root
			}
		} else if k < root.Value {
			return SearchBSTForFirstOccurrenceRecursive(root.Left, k)
		} else {
			return SearchBSTForFirstOccurrenceRecursive(root.Right, k)
		}
	} else {
		return nil
	}
}

func SearchBSTForFirstOccurrenceIterative(root *TreeNode, k int) *TreeNode {
	var first *TreeNode

	for root != nil {
		if root.Value == k {
			first = root
			root = root.Left
		} else if k < root.Value {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return first
}
