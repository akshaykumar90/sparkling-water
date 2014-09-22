// Problem 15.14

package chapter15

func ComputeBinomialCoefficients(n, k int) int {
	table := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		table[i] = 1
		for j := i - 1; j > 0; j-- {
			table[j] += table[j-1]
		}
	}

	return table[k]
}
