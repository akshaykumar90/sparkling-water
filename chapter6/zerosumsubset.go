// Problem 6.5

package chapter6

func ZeroSumSubset(arr []int) []int {
	n := len(arr)

	prefixSum := make([]int, n)
	copy(prefixSum, arr)

	for i := range arr {
		if i > 0 {
			prefixSum[i] += prefixSum[i-1]
		}
		prefixSum[i] %= n
	}

	index := make([]int, n)
	for i := range index {
		index[i] = -1
	}

	result := make([]int, 0)

	for i, ps := range prefixSum {
		if ps == 0 {
			for j := 0; j <= i; j++ {
				result = append(result, j)
			}
			return result
		}
		if index[ps] != -1 {
			for j := index[ps] + 1; j <= i; j++ {
				result = append(result, j)
			}
			return result
		}
		index[ps] = i
	}
	return result
}
