// Problem 16.2

package chapter16

func TransformStringToOther(D map[string]bool, s, t string) int {
	lengths := make(map[string]int)
	visited := make(map[string]bool)

	lengths[s] = 0
	visited[s] = true

	q := []string{s}

	for len(q) > 0 {
		curr := q[0]

		dist := lengths[curr]

		if curr == t {
			return dist
		}

		runes := []rune(curr)

		for i, r := range runes {
			for j := 0; j < 26; j++ {
				a := 'a' + rune(j)
				if r != a {
					runes[i] = a

					if next := string(runes); D[next] && !visited[next] {
						lengths[next] = dist + 1
						visited[next] = true
						q = append(q, next)
					}

					runes[i] = r
				}
			}
		}

		q = q[1:]
	}

	return -1
}
