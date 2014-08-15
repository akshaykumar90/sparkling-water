// problem 7.7

package chapter7

func DeletionList(elem *Element) {
	elem.Value = elem.Next.Value
	elem.Next = elem.Next.Next
}
