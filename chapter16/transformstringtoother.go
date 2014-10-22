// Problem 16.2

package chapter16

func transformHelper(D map[string]bool, s, t string) map[string]string {
	parent := make(map[string]string)
	visited := make(map[string]bool)
	q := make([]string, 0)

	parent[s] = ""
	visited[s] = true
	q = append(q, s)

	for len(q) > 0 {
		var x string
		x, q = q[0], q[1:]

		runes := []rune(x)
		alphabet := []rune("abcdefghijklmnopqrstuvwxyz")
		for i, r := range runes {
			for _, a := range alphabet {
				if r != a {
					runes[i] = a
					temp := string(runes)
					if D[temp] && !visited[temp] {
						parent[temp] = x
						if temp == t {
							return parent
						} else {
							visited[temp] = true
							q = append(q, temp)
						}
					}
					runes[i] = r
				}
			}
		}
	}

	return parent
}

func TransformStringToOther(D map[string]bool, s, t string) int {
	parent := transformHelper(D, s, t)
	if p, ok := parent[t]; ok {
		i := 0
		for ; p != ""; i++ {
			p = parent[p]
		}
		return i
	} else {
		return -1
	}
}
