package chapter14

import (
	"fmt"
	"strings"
	"testing"
)

func preOrderPrettyPrint(node *TreeNode, depth int) {
	if node == nil {
		fmt.Printf("%s%s\n", strings.Repeat(">", depth), "nil")
	} else {
		fmt.Printf("%s%d\n", strings.Repeat(">", depth), node.Value)
		preOrderPrettyPrint(node.Left, depth+1)
		preOrderPrettyPrint(node.Right, depth+1)
	}
}

func TestBuildBSTFromSortedArray(t *testing.T) {
	root := BuildBSTFromSortedArray([]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53})

	fmt.Println("TestBuildBSTFromSortedArray:")
	preOrderPrettyPrint(root, 0)
	fmt.Println()
}
