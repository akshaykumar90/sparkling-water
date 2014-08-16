// problem 7.8

package chapter7

import "errors"

func RemoveKthLastList(head *Element, k int) (*Element, error) {
	var ahead *Element

	for ahead = head; k > 0; k-- {
		if ahead != nil {
			ahead = ahead.Next
		} else {
			return nil, errors.New("RemoveKthLastList: list has fewer than k elements")
		}
	}

	curr := head
	var prev *Element
	for ahead != nil {
		prev, curr = curr, curr.Next
		ahead = ahead.Next
	}

	if prev != nil {
		prev.Next = curr.Next
	} else {
		head = curr.Next
	}

	return head, nil
}
