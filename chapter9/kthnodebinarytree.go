// problem 9.7

package chapter9

type TreeElementWithSize struct {
	Value int
	Size  int
	Left  *TreeElementWithSize
	Right *TreeElementWithSize
}

func KThNodeBinaryTree(root *TreeElementWithSize, k int) *TreeElementWithSize {
	l := 0
	if root.Left != nil {
		l = root.Left.Size
	}
	if k <= l {
		return KThNodeBinaryTree(root.Left, k)
	} else if k == l+1 {
		return root
	} else {
		return KThNodeBinaryTree(root.Right, k-l-1)
	}
}
