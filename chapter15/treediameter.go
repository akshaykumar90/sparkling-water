// problem 15.4

package chapter15

type WeightedTreeNode struct {
	Value    int
	Children []*WeightedTreeNode
}

func searchTree(root *WeightedTreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	first, second := 0, 0
	maxTreeDiameter := 0

	for _, ch := range root.Children {
		depth, maxSubtreeDiameter := searchTree(ch)
		if first < depth {
			first, second = depth, first
		} else if second < depth {
			second = depth
		}
		if maxTreeDiameter < maxSubtreeDiameter {
			maxTreeDiameter = maxSubtreeDiameter
		}
	}

	if maxTreeDiameter < first+second {
		maxTreeDiameter = first + second
	}

	return first + root.Value, maxTreeDiameter
}

func TreeDiameter(root *WeightedTreeNode) int {
	_, diameter := searchTree(root)
	return diameter
}
