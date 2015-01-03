// Problem 17.2

package chapter17

import "github.com/akshaykumar90/sparkling-water/common"

type item struct {
	w, v int
}

func ZeroOneKnapsack(items []item, W int) int {
	table := make([]int, W+1)

	for i := 0; i < len(items); i++ {
		for j := W; j >= items[i].w; j-- {
			table[j] = common.MaxInt(table[j], table[j-items[i].w]+items[i].v)
		}
	}

	return table[W]
}
