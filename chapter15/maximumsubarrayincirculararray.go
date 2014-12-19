// Problem 15.5

package chapter15

func optimalSubarrayUsingCmp(arr []int, cmp func(int, int) int) int {
	var till, overall int

	for _, a := range arr {
		till = cmp(a, a+till)
		overall = cmp(overall, till)
	}

	return overall
}

func MaximumSubarrayInCircularArray(arr []int) int {
	sum := 0
	for _, a := range arr {
		sum += a
	}

	min := optimalSubarrayUsingCmp(arr, func(a, b int) int {
		if a <= b {
			return a
		} else {
			return b
		}
	})

	max := optimalSubarrayUsingCmp(arr, func(a, b int) int {
		if a >= b {
			return a
		} else {
			return b
		}
	})

	if csum := sum - min; csum > max {
		return csum
	} else {
		return max
	}
}
