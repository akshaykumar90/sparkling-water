package chapter8

import (
	"fmt"
	"testing"
)

func TestSearchPostingsListRecursive(t *testing.T) {
	a := &PostingsListElement{Value: 1, Order: -1}
	b := &PostingsListElement{Value: 2, Order: -1}
	c := &PostingsListElement{Value: 3, Order: -1}
	d := &PostingsListElement{Value: 4, Order: -1}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = nil

	a.Jump = c
	b.Jump = d
	c.Jump = b
	d.Jump = d

	fmt.Println("TestSearchPostingsListRecursive:")
	SearchPostingsListRecursive(a)

	curr := a
	for curr != nil {
		fmt.Println(curr.Order)
		curr = curr.Next
	}

	fmt.Println()
}

func TestSearchPostingsListIterative(t *testing.T) {
	a := &PostingsListElement{Value: 1, Order: -1}
	b := &PostingsListElement{Value: 2, Order: -1}
	c := &PostingsListElement{Value: 3, Order: -1}
	d := &PostingsListElement{Value: 4, Order: -1}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = nil

	a.Jump = c
	b.Jump = d
	c.Jump = b
	d.Jump = d

	fmt.Println("TestSearchPostingsListIterative:")
	SearchPostingsListIterative(a)

	curr := a
	for curr != nil {
		fmt.Println(curr.Order)
		curr = curr.Next
	}

	fmt.Println()
}
