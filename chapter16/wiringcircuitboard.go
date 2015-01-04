// Problem 16.3

package chapter16

func WiringCircuitBoard(board [][]int) bool {
	dist := make([]int, len(board))

	for i := range dist {
		dist[i] = -1
	}

	bfs := func(v int) bool {
		q := []int{v}

		for len(q) > 0 {
			curr := q[0]

			for _, e := range board[curr] {
				if dist[e] == -1 {
					dist[e] = dist[curr] + 1
					q = append(q, e)
				} else if dist[e] == dist[curr] {
					return false
				}
			}

			q = q[1:]
		}

		return true
	}

	for v := range board {
		if dist[v] == -1 {
			dist[v] = 0
			if bfs(v) == false {
				return false
			}
		}
	}

	return true
}
