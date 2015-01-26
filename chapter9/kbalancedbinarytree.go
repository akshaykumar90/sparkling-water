// Problem 9.2

package chapter9

import "github.com/akshaykumar90/sparkling-water/common"

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

		if diff := common.Abs(ln - rn); diff > k {
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
