// Problem 9.9

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

func ReconstructPreorderWithNull2(vs []string) *TreeNode {
	stack := make([]*TreeNode, 0)

	for i := len(vs) - 1; i >= 0; i-- {
		if vs[i] != "null" {
			sp := len(stack)
			node := TreeNode{vs[i], stack[sp-1], stack[sp-2]}
			stack = stack[:sp-2]
			stack = append(stack, &node)
		} else {
			stack = append(stack, nil)
		}
	}

	return stack[len(stack)-1]
}
