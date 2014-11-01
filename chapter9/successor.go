// Problem 9.6

package chapter9

func Successor(node *TreeNodeWithParent) *TreeNodeWithParent {
	var succ *TreeNodeWithParent
	if node.Right != nil {
		succ = node.Right
		for succ.Left != nil {
			succ = succ.Left
		}
	} else {
		succ = node.Parent
		for ; succ != nil && succ.Left != node; node, succ = succ, succ.Parent {
		}
	}
	return succ
}
