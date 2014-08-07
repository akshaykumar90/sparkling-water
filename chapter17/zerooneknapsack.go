// problem 17.2

package chapter17

type item struct {
	w, v int
}

func ZeroOneKnapsack(items []item, W int) int {
	table := make([][]int, len(items)+1)
	for i := range table {
		table[i] = make([]int, W+1)
	}

	for i := 1; i <= len(items); i++ {
		for j := 1; j <= W; j++ {
			if items[i-1].w <= j {
				withItem := table[i-1][j-items[i-1].w] + items[i-1].v
				withoutItem := table[i-1][j]
				if withItem > withoutItem {
					table[i][j] = withItem
				} else {
					table[i][j] = withoutItem
				}
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}

	return table[len(items)][W]
}
