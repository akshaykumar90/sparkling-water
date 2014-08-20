// Problem 9.8

package chapter9

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

func ReconstructBinaryTreePreInOrders(pre, in []string) *TreeNode {
	indexInSlice := func(val string, slice []string) int {
		for i, v := range slice {
			if v == val {
				return i
			}
		}
		return -1
	}

	if len(pre) == 0 {
		return nil
	} else {
		mid := indexInSlice(pre[0], in)

		n := new(TreeNode)
		n.Value = pre[0]
		n.Left = ReconstructBinaryTreePreInOrders(pre[1:mid+1], in[:mid])
		n.Right = ReconstructBinaryTreePreInOrders(pre[mid+1:], in[mid+1:])
		return n
	}
}
