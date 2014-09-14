// problem 16.9

package chapter16

import "math"

type edgenode struct {
	y, weight int
}

type wdist struct {
	dist, edges int
}

func (p wdist) Less(than wdist) bool {
	return p.dist < than.dist || (p.dist == than.dist && p.edges < than.edges)
}

func (p wdist) Plus(weight int) wdist {
	return wdist{p.dist + weight, p.edges + 1}
}

func ShortestPathFewestEdges(graph [][]edgenode, start, end int) []int {
	distance := make([]wdist, len(graph))
	parent := make([]int, len(graph))
	intree := make([]bool, len(graph))

	for i := range distance {
		distance[i] = wdist{math.MaxInt32, math.MaxInt32}
		parent[i] = -1
	}

	distance[start] = wdist{0, 0}

	for v := start; v != end; {
		intree[v] = true

		for _, edge := range graph[v] {
			w := edge.y
			if distance[v].Plus(edge.weight).Less(distance[w]) {
				distance[w] = distance[v].Plus(edge.weight)
				parent[w] = v
			}
		}

		min := wdist{math.MaxInt32, math.MaxInt32}
		for i := range distance {
			if !intree[i] && distance[i].Less(min) {
				min = distance[i]
				v = i
			}
		}
	}

	path := make([]int, 0)
	for v := end; v != -1; v = parent[v] {
		path = append(path, v)
	}

	n := len(path)
	for i := 0; i < n/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}

	return path

}
