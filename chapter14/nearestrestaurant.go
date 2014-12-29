// Problem 14.14

package chapter14

func rangeQuery(root *TreeNode, lo, hi int, result *[]int) {
	if root != nil {
		if hi < root.Value {
			rangeQuery(root.Left, lo, hi, result)
		} else if lo > root.Value {
			rangeQuery(root.Right, lo, hi, result)
		} else {
			rangeQuery(root.Left, lo, hi, result)
			*result = append(*result, root.Value)
			rangeQuery(root.Right, lo, hi, result)
		}
	}
}

func NearestRestaurant(root *TreeNode, lo, hi int) []int {
	result := make([]int, 0)
	rangeQuery(root, lo, hi, &result)
	return result
}
