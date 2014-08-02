// problem 14.4

package chapter14

func search(root *TreeNode, k int, found_k bool) (*TreeNode, bool) {
	if root == nil {
		return nil, found_k
	}

	var succ *TreeNode

	if !found_k {
		if k == root.Value {
			return search(root.Right, k, true)
		} else if k < root.Value {
			succ, found_k = search(root.Left, k, found_k)
		} else {
			return search(root.Right, k, found_k)
		}
	} else {
		succ, found_k = search(root.Left, k, found_k)
	}

	if found_k && succ != nil {
		return succ, found_k
	} else if found_k {
		return root, found_k
	} else {
		return nil, found_k
	}
}

func SearchBSTFirstLargerK(root *TreeNode, k int) *TreeNode {
	succ, _ := search(root, k, false)
	return succ
}
