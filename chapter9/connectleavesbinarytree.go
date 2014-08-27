// problem 9.10

package chapter9

type LeavesList []*TreeElement

func visitTreeElementInorder(elem *TreeElement, list *LeavesList) {
	if elem == nil {
		return
	} else {
		if elem.Left == nil && elem.Right == nil {
			*list = append(*list, elem)
		} else {
			visitTreeElementInorder(elem.Left, list)
			visitTreeElementInorder(elem.Right, list)
		}
	}
}

func ConnectLeavesBinaryTree(root *TreeElement) LeavesList {
	list := make(LeavesList, 0)
	visitTreeElementInorder(root, &list)
	return list
}
