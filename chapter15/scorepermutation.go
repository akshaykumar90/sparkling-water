// problem 15.15

package chapter15

func ScorePermutation(W []int, s int) int {
	permutations := make([]int, s+1)

	permutations[0] = 1

	for i := range permutations {
		for _, w := range W {
			if i >= w {
				permutations[i] += permutations[i-w]
			}
		}
	}

	return permutations[s]
}
