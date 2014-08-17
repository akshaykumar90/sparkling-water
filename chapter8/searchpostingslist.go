// problem 8.4

package chapter8

import "errors"

type PostingsListElement struct {
	Value int
	Order int
	Next  *PostingsListElement
	Jump  *PostingsListElement
}

func jumpFirstSearch(curr *PostingsListElement, order *int) {
	if curr != nil && curr.Order == -1 {
		*order++
		curr.Order = *order
		jumpFirstSearch(curr.Jump, order)
		jumpFirstSearch(curr.Next, order)
	}
}

func SearchPostingsListRecursive(head *PostingsListElement) {
	order := -1
	jumpFirstSearch(head, &order)
}

type PostingsListElementStack []*PostingsListElement

func (st *PostingsListElementStack) Push(x *PostingsListElement) {
	*st = append(*st, x)
}

func (st *PostingsListElementStack) Pop() (*PostingsListElement, error) {
	if len(*st) == 0 {
		return nil, errors.New("Pop: empty stack")
	} else {
		var x *PostingsListElement
		x, *st = (*st)[len(*st)-1], (*st)[:len(*st)-1]
		return x, nil
	}
}

func SearchPostingsListIterative(head *PostingsListElement) {
	st := make(PostingsListElementStack, 0)
	order := -1

	st.Push(head)
	for len(st) > 0 {
		curr, _ := st.Pop()
		if curr != nil && curr.Order == -1 {
			order++
			curr.Order = order
			st.Push(curr.Next)
			st.Push(curr.Jump)
		}
	}
}
