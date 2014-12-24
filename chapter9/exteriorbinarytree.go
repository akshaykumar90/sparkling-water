// Problem 9.11

package chapter9

func ExteriorBinaryTree(root *TreeElement) []int {
	result := make([]int, 0)

	var visitLeft, visitRight func(*TreeElement, bool)

	visitLeft = func(elem *TreeElement, isBoundary bool) {
		if elem != nil {
			if isBoundary || (elem.Left == nil && elem.Right == nil) {
				result = append(result, elem.Value)
			}
			visitLeft(elem.Left, isBoundary)
			visitLeft(elem.Right, isBoundary && elem.Left == nil)
		}
	}

	visitRight = func(elem *TreeElement, isBoundary bool) {
		if elem != nil {
			visitRight(elem.Left, isBoundary && elem.Right == nil)
			visitRight(elem.Right, isBoundary)
			if isBoundary || (elem.Left == nil && elem.Right == nil) {
				result = append(result, elem.Value)
			}
		}
	}

	if root != nil {
		result = append(result, root.Value)
		visitLeft(root.Left, true)
		visitRight(root.Right, true)
	}

	return result
}
