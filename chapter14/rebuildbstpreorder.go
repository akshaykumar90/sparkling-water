// problem 14.11

package chapter14

func RebuildBSTPreorder(pre []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	} else {
		i := 1
		for ; i < len(pre) && pre[i] < pre[0]; i++ {
		}

		n := new(TreeNode)
		n.Value = pre[0]
		n.Left = RebuildBSTPreorder(pre[1:i])
		n.Right = RebuildBSTPreorder(pre[i:])

		return n
	}
}
