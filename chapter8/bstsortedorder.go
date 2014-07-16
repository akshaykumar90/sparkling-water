package chapter8

import (
	"errors"
	"fmt"
)

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

type Stack []*BST

func (st *Stack) Push(x *BST) {
	*st = append(*st, x)
}

func (st *Stack) Pop() (*BST, error) {
	if len(*st) == 0 {
		return nil, errors.New("Pop: empty stack")
	} else {
		var x *BST
		x, *st = (*st)[len(*st)-1], (*st)[:len(*st)-1]
		return x, nil
	}
}

func BSTSortedOrder(root *BST) {
	var path Stack

	for root != nil || len(path) > 0 {
		switch root {
		case nil:
			parent, _ := path.Pop()
			fmt.Println(parent.Value)
			root = parent.Right
		default:
			path.Push(root)
			root = root.Left
		}
	}
}
