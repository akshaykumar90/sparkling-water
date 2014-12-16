// Problem 15.6

package chapter15

import "github.com/akshaykumar90/sparkling-water/chapter11"

func LongestNondecreasingSubsequence(arr []int) int {
	M := make([]int, 0, len(arr))

	for _, a := range arr {
		if i := chapter11.BinarySearchLargerK(M, a); i == -1 {
			M = append(M, a)
		} else {
			M[i] = a
		}
	}

	return len(M)
}
