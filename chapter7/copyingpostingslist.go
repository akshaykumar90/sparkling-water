// Problem 7.12

package chapter7

type PostingsListElement struct {
	Value int
	Next  *PostingsListElement
	Jump  *PostingsListElement
}

func CopyingPostingsList(head *PostingsListElement) *PostingsListElement {
	if head == nil {
		return nil
	}

	for curr := head; curr != nil; {
		copy := &PostingsListElement{Value: curr.Value, Next: curr.Next}
		curr.Next = copy
		curr = copy.Next
	}

	for curr := head; curr != nil; curr = curr.Next.Next {
		curr.Next.Jump = curr.Jump.Next
	}

	headCopy := head.Next

	for curr := head; curr.Next != nil; {
		temp := curr.Next
		curr.Next = temp.Next
		curr = temp
	}

	return headCopy
}
