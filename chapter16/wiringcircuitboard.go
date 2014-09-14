// problem 16.3

package chapter16

func dfsHelper(adj [][]int) ([]bool, func(int) bool) {
	n := len(adj)
	state := make([]int, n)
	parent := make([]int, n)
	color := make([]bool, n)

	var dfs func(int) bool
	dfs = func(pin int) bool {
		state[pin] = DISCOVERED

		for _, v := range adj[pin] {
			if state[v] == UNDISCOVERED {
				parent[v] = pin
				color[v] = !color[pin]
				if !dfs(v) {
					return false
				}
			} else if state[v] == DISCOVERED && color[v] == color[pin] {
				return false
			}
		}

		state[pin] = VISITED
		return true
	}

	return color, dfs
}

func WiringCircuitBoard(board [][]int) bool {
	_, dfsFn := dfsHelper(board)
	return dfsFn(0)
}
