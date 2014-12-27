// Problem 8.9

package chapter8

func BinaryTreeLevelOrder(root *BST) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	q := []*BST{root}
	cur, next := 1, 0
	for cur > 0 {
		level := make([]int, 0, cur)
		for ; cur > 0; cur-- {
			x := q[0]
			q = q[1:]

			level = append(level, x.Value)

			if x.Left != nil {
				q = append(q, x.Left)
				next++
			}
			if x.Right != nil {
				q = append(q, x.Right)
				next++
			}
		}
		cur, next = next, 0
		result = append(result, level)
	}

	return result
}
