// Problem 15.7

package chapter15

import "math"

func binarySearchSmallerK(arr []int, k int) int {
	res := -1

	lo, hi := 0, len(arr)-1

	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] >= k {
			hi = mid - 1
		} else {
			res = mid
			lo = mid + 1
		}
	}

	return res
}

func LongestSubarrayK(arr []int, k int) int {
	prefixSum := make([]int, 0, len(arr))

	sum := 0

	for _, a := range arr {
		sum += a
		prefixSum = append(prefixSum, sum)
	}

	minPrefixSum := make([]int, len(arr))

	currMin := math.MaxInt32

	for i := len(prefixSum) - 1; i >= 0; i-- {
		if prefixSum[i] < currMin {
			currMin = prefixSum[i]
		}
		minPrefixSum[i] = currMin
	}

	maxLength := 0

	for i, s := range prefixSum {
		j := binarySearchSmallerK(minPrefixSum[i+1:], k+s)
		if length := j + 1; length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
