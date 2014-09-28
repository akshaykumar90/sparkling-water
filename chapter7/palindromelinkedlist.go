// Problem 7.10

package chapter7

func PalindromeLinkedList(head *Element) bool {
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}

	reverse := ReverseLinkedList(slow)

	for head != nil && reverse != nil {
		if head.Value != reverse.Value {
			return false
		}
		head = head.Next
		reverse = reverse.Next
	}
	return true
}
