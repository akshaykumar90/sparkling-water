// problem 15.15

package chapter15

func ScoreCombination(W []int, s int) int {
	combinations := make([]int, s+1)

	combinations[0] = 1

	for _, w := range W {
		for j := w; j <= s; j++ {
			combinations[j] += combinations[j-w]
		}
	}

	return combinations[s]
}
