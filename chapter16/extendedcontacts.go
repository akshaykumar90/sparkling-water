// problem 16.5

package chapter16

func contactsSearchHelper(graph [][]int) (*[]int, func(int)) {
	contacts := make([]int, 0)

	contactsPtr := &contacts

	state := make([]int, len(graph))

	var dfs func(int)
	dfs = func(v int) {
		state[v] = DISCOVERED
		for _, e := range graph[v] {
			if state[e] == UNDISCOVERED {
				*contactsPtr = append(*contactsPtr, e)
				dfs(e)
			}
		}
	}

	return contactsPtr, dfs
}

func ExtendedContacts(graph [][]int) [][]int {
	extendedContacts := make([][]int, len(graph))

	for i := range graph {
		contactsPtr, dfsFn := contactsSearchHelper(graph)
		dfsFn(i)
		extendedContacts[i] = *contactsPtr
	}

	return extendedContacts
}
