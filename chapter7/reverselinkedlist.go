// problem 7.9

package chapter7

func ReverseLinkedList(head *Element) *Element {
	var prev *Element = nil
	var next *Element

	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
