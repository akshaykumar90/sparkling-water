// problem 9.2

package chapter9

type TreeElement struct {
	Value int
	Left  *TreeElement
	Right *TreeElement
}

func visitTreeElement(elem *TreeElement, k int) (int, *TreeElement) {
	if elem == nil {
		return 0, nil
	} else {
		ln, ld := visitTreeElement(elem.Left, k)
		if ld != nil {
			return 0, ld
		}

		rn, rd := visitTreeElement(elem.Right, k)
		if rd != nil {
			return 0, rd
		}

		diff := ln - rn
		if diff < 0 {
			diff = -diff
		}
		if diff > k {
			return 0, elem
		} else {
			return ln + rn + 1, nil
		}
	}
}

func KBalancedBinaryTree(root *TreeElement, k int) *TreeElement {
	_, found := visitTreeElement(root, k)
	return found
}
