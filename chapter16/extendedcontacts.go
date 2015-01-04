// Problem 16.5

package chapter16

func ExtendedContacts(graph [][]int) [][]int {
	extendedContacts := make([][]int, 0)

	var dfs func([]bool, int, *[]int)
	dfs = func(state []bool, v int, contacts *[]int) {
		state[v] = true
		for _, e := range graph[v] {
			if state[e] == false {
				*contacts = append(*contacts, e)
				dfs(state, e, contacts)
			}
		}
	}

	for i := range graph {
		state := make([]bool, len(graph))
		contacts := make([]int, 0)
		dfs(state, i, &contacts)
		extendedContacts = append(extendedContacts, contacts)
	}

	return extendedContacts
}
