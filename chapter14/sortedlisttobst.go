// Problem 14.8

package chapter14

func sortedListToBSTHelper(head **TreeNode, lo, hi int) *TreeNode {
	if lo <= hi {
		mid := (lo + hi) / 2
		left := sortedListToBSTHelper(head, lo, mid-1)
		curr := *head
		*head = (*head).Right
		curr.Left = left
		curr.Right = sortedListToBSTHelper(head, mid+1, hi)
		return curr
	} else {
		return nil
	}
}

func SortedListToBST(head *TreeNode, n int) *TreeNode {
	return sortedListToBSTHelper(&head, 0, n-1)
}
