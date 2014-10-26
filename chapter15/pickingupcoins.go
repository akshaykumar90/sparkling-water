// Problem 15.18

package chapter15

func PickingUpCoins(C []int) int {
	mm := make([][]int, len(C))

	for i := range mm {
		mm[i] = make([]int, len(C))
	}

	max := func(a, b int) int {
		if a < b {
			return b
		} else {
			return a
		}
	}

	min := func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	var helper func(int, int) int
	helper = func(a, b int) int {
		if a <= b {
			if mm[a][b] == 0 {
				mm[a][b] = max(
					C[a]+min(helper(a+2, b), helper(a+1, b-1)),
					C[b]+min(helper(a+1, b-1), helper(a, b-2)),
				)
			}
			return mm[a][b]
		} else {
			return 0
		}
	}

	return helper(0, len(C)-1)
}
