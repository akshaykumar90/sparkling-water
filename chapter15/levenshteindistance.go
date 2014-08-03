// problem 15.11

package chapter15

func LevenshteinDistance(a, b string) int {
	table := make([][]int, len(a)+1)

	for i := 0; i <= len(a); i++ {
		table[i] = make([]int, len(b)+1)
		table[i][0] = i
	}

	for j := 0; j <= len(b); j++ {
		table[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				table[i][j] = table[i-1][j-1]
			} else {
				options := []int{
					table[i-1][j],
					table[i-1][j-1],
					table[i][j-1],
				}
				min := options[0]
				for _, x := range options[1:] {
					if x < min {
						min = x
					}
				}
				table[i][j] = min + 1
			}
		}
	}

	return table[len(a)][len(b)]
}
