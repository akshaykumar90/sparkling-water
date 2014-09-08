// problem 14.10

package chapter14

type TreeNodeList []*TreeNode

func (p TreeNodeList) Len() int           { return len(p) }
func (p TreeNodeList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p TreeNodeList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func extendKLargestBST(node *TreeNode, k int, list *TreeNodeList) {
	if node != nil {
		extendKLargestBST(node.Right, k, list)
		if len(*list) < k {
			*list = append(*list, node)
			extendKLargestBST(node.Left, k, list)
		}
	}
}

func FindKLargestBST(root *TreeNode, k int) TreeNodeList {
	list := make(TreeNodeList, 0)
	extendKLargestBST(root, k, &list)
	return list
}
