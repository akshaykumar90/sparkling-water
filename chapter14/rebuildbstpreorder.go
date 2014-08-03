// problem 14.11

package chapter14

import "math"

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

func rebuildBSTPreorderWithinLimits(pre []int, i, min, max int) (int, *TreeNode) {
	if i == len(pre) {
		return 0, nil
	}

	if pre[i] < min || pre[i] > max {
		return 0, nil
	} else {
		n := new(TreeNode)
		n.Value = pre[i]
		j, left := rebuildBSTPreorderWithinLimits(pre, i+1, min, pre[i])
		k, right := rebuildBSTPreorderWithinLimits(pre, i+j+1, pre[i], max)
		n.Left = left
		n.Right = right
		return j + k + 1, n
	}
}

func RebuildBSTPreorderBetter(pre []int) *TreeNode {
	_, t := rebuildBSTPreorderWithinLimits(pre, 0, math.MinInt32, math.MaxInt32)
	return t
}
