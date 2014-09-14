// problem 16.6

package chapter16

type Pair struct {
	fst, snd int
}

func partitionHelper(adj map[int]map[int]bool) ([]int, func(int, int)) {
	partition := make([]int, len(adj)+1)
	for i := range partition {
		partition[i] = -1
	}

	var dfs func(int, int)
	dfs = func(v int, n int) {
		partition[v] = n
		for e := range adj[v] {
			if partition[e] == -1 {
				dfs(e, n)
			}
		}
	}

	return partition, dfs
}

func AreConstraintsSatisfied(eq []Pair, ieq []Pair) bool {
	graph := make(map[int]map[int]bool)

	for _, p := range eq {
		if _, ok := graph[p.fst]; !ok {
			graph[p.fst] = make(map[int]bool)
		}
		graph[p.fst][p.snd] = true

		if _, ok := graph[p.snd]; !ok {
			graph[p.snd] = make(map[int]bool)
		}
		graph[p.snd][p.fst] = true
	}

	partition, dfsFn := partitionHelper(graph)

	n := 0
	for v := range graph {
		if partition[v] == -1 {
			dfsFn(v, n)
			n++
		}
	}

	for _, p := range ieq {
		if partition[p.fst] == partition[p.snd] {
			return false
		}
	}

	return true

}
