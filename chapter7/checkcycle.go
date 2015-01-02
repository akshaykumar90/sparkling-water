// problem 7.2

package chapter7

func CheckCycle(head *Element) *Element {
	i, j := head, head

	for i != nil && j != nil {
		i, j = i.Next, j.Next

		if j != nil {
			j = j.Next
		}

		if i == j {
			for i = head; i != j; i, j = i.Next, j.Next {
			}
			return i
		}
	}

	return nil
}
