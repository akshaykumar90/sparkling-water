package chapter9

import "fmt"

type TreeNodeWithParent struct {
	Value  int
	Left   *TreeNodeWithParent
	Right  *TreeNodeWithParent
	Parent *TreeNodeWithParent
}

const (
	up   = iota
	down = iota
)

func InorderTraversalWithParent(root *TreeNodeWithParent) {
	var cur, prev *TreeNodeWithParent = root, nil
	var dir int = down

	for cur != nil {
		switch dir {
		case down:
			if cur.Left != nil {
				cur, prev = cur.Left, cur
			} else {
				fmt.Println(cur.Value)
				if cur.Right != nil {
					cur, prev = cur.Right, cur
				} else {
					cur, prev = prev, cur
					dir = up
				}
			}
		case up:
			if cur.Left == prev {
				fmt.Println(cur.Value)
				if cur.Right != nil {
					cur, prev = cur.Right, cur
					dir = down
				} else {
					cur, prev = cur.Parent, cur
				}
			} else {
				cur, prev = cur.Parent, cur
			}
		}
	}
}
