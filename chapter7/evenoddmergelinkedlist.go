// problem 7.6

package chapter7

func EvenOddMergeLinkedList(head *Element) *Element {
	curr := head

	var evenHead, oddHead *Element = nil, nil
	var prevEven, prevOdd *Element = nil, nil

	isEven := true
	for curr != nil {
		switch isEven {
		case true:
			if prevEven != nil {
				prevEven.Next = curr
			} else {
				evenHead = curr
			}
			prevEven = curr
		case false:
			if prevOdd != nil {
				prevOdd.Next = curr
			} else {
				oddHead = curr
				prevOdd = oddHead
			}
			prevOdd = curr
		}
		curr = curr.Next
		isEven = !isEven
	}

	if prevEven != nil {
		prevEven.Next = oddHead
	}
	if prevOdd != nil {
		prevOdd.Next = nil
	}

	return evenHead
}
