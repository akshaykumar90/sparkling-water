// problem 9.9

package chapter9

func consumePreorder(vs []string, st int) (*TreeNode, int) {
	if vs[st] == "null" {
		return nil, 1
	} else {
		elem := &TreeNode{Value: vs[st]}
		left, lc := consumePreorder(vs, st+1)
		elem.Left = left
		right, rc := consumePreorder(vs, st+1+lc)
		elem.Right = right
		return elem, lc + rc + 1
	}
}

func ReconstructPreorderWithNull(vs []string) *TreeNode {
	root, _ := consumePreorder(vs, 0)
	return root
}
