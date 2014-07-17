package chapter8

import "fmt"

func BinaryTreeLevelOrder(root *BST) {
	q := []*BST{root}
	cur, next := 1, 0
	for cur > 0 {
		for ; cur > 0; cur-- {
			x := q[0]
			q = q[1:]
			fmt.Print(x.Value)
			fmt.Print(" ")

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
		fmt.Println()
	}
}
