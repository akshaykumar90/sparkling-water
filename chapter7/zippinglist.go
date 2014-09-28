// Problem 7.11

package chapter7

func ZippingList(head *Element) *Element {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}

	reverse := ReverseLinkedList(slow)

	retVal := head

	if head != nil {
		for head.Next != nil && reverse.Next != nil {
			left, right := head.Next, reverse.Next
			head.Next = reverse
			reverse.Next = left
			head, reverse = left, right
		}
	}

	return retVal
}
