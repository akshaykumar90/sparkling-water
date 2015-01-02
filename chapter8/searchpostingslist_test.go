package chapter8

import "testing"

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

	expected := []int{0, 2, 1, 3}

	SearchPostingsListRecursive(a)

	curr := a
	i := 0
	for curr != nil {
		if actual := curr.Order; actual != expected[i] {
			t.Fatalf("SearchPostingsListRecursive: at %d: expected %d, actual %d", i, expected[i], actual)
		}
		curr = curr.Next
		i++
	}
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

	expected := []int{0, 2, 1, 3}

	SearchPostingsListIterative(a)

	curr := a
	i := 0
	for curr != nil {
		if actual := curr.Order; actual != expected[i] {
			t.Fatalf("SearchPostingsListIterative: at %d: expected %d, actual %d", i, expected[i], actual)
		}
		curr = curr.Next
		i++
	}
}
