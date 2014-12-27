// Problem 9.5

package chapter9

type TreeNodeWithParent struct {
	Value  int
	Left   *TreeNodeWithParent
	Right  *TreeNodeWithParent
	Parent *TreeNodeWithParent
}

const (
	left   = iota
	right  = iota
	parent = iota
)

func InorderTraversalWithParent(root *TreeNodeWithParent) []int {
	result := make([]int, 0)

	getNextNode := func(node *TreeNodeWithParent, nodeType int) *TreeNodeWithParent {
		switch nodeType {
		case left:
			if node.Left != nil {
				return node.Left
			}
			fallthrough
		case right:
			if node.Right != nil {
				return node.Right
			}
			fallthrough
		default:
			return node.Parent
		}
	}

	var curr, prev *TreeNodeWithParent = root, nil

	for curr != nil {
		var nextNodeType int
		switch prev {
		case nil:
			fallthrough
		case curr.Parent:
			if curr.Left == nil {
				result = append(result, curr.Value)
			}
			nextNodeType = left
		case curr.Left:
			result = append(result, curr.Value)
			nextNodeType = right
		case curr.Right:
			nextNodeType = parent
		}
		prev, curr = curr, getNextNode(curr, nextNodeType)
	}

	return result
}
