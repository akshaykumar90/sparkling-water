// problem 16.7

package chapter16

import "github.com/akshaykumar90/sparkling-water/chapter13"

var depths []int

func dfsTeam(edges [][]bool, state []int, st int) {
	state[st] = DISCOVERED

	for i := 0; i < len(edges[st]); i++ {
		if edges[st][i] && state[i] == UNDISCOVERED {
			dfsTeam(edges, state, i)
		}
	}

	state[st] = VISITED

	maxDepth := 0
	for i := 0; i < len(edges[st]); i++ {
		if edges[st][i] && state[i] == VISITED {
			if maxDepth < depths[i] {
				maxDepth = depths[i]
			}
		}
	}

	depths[st] = maxDepth + 1
}

func TeamPhoto2(teamHeights [][]int) int {
	edges := make([][]bool, len(teamHeights))

	for i := range edges {
		edges[i] = make([]bool, len(teamHeights))
	}

	for i, fst := range teamHeights {
		for j, snd := range teamHeights[i+1:] {
			if edge := chapter13.TeamPhoto1(fst, snd); edge {
				edges[i+j+1][i] = true
			} else if reverseEdge := chapter13.TeamPhoto1(snd, fst); reverseEdge {
				edges[i][i+j+1] = true
			}
		}
	}

	depths = make([]int, len(teamHeights))
	state := make([]int, len(teamHeights))

	for i := range teamHeights {
		if state[i] == UNDISCOVERED {
			dfsTeam(edges, state, i)
		}
	}

	longestPath := 0
	for _, x := range depths {
		if longestPath < x {
			longestPath = x
		}
	}

	return longestPath
}
