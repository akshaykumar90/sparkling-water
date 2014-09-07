// problem 14.6

package chapter14

func BuildBSTFromSortedArray(a []int) *TreeNode {
	if len(a) > 0 {
		mid := len(a) / 2
		return &TreeNode{a[mid],
			BuildBSTFromSortedArray(a[:mid]),
			BuildBSTFromSortedArray(a[mid+1:]),
		}
	} else {
		return nil
	}
}
