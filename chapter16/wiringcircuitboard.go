// problem 16.3

package chapter16

func dfsHelper(adj [][]int) ([]bool, func(int) bool) {
	color := make([]bool, len(adj))

	state := make([]int, len(adj))

	var dfs func(int) bool
	dfs = func(pin int) bool {
		state[pin] = DISCOVERED
		for _, v := range adj[pin] {
			if state[v] == UNDISCOVERED {
				color[v] = !color[pin]
				if !dfs(v) {
					return false
				}
			} else if color[v] == color[pin] {
				return false
			}
		}
		return true
	}

	return color, dfs
}

func WiringCircuitBoard(board [][]int) bool {
	_, dfsFn := dfsHelper(board)
	return dfsFn(0)
}
