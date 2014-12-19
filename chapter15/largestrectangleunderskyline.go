// Problem 15.8

package chapter15

func largestRectangleHelper(hs []int) []int {
	stack := make([]int, 0, len(hs))

	longest := make([]int, len(hs))

	for i, h := range hs {
		j := len(stack) - 1

		for ; j >= 0 && hs[stack[j]] >= h; j-- {
		}

		if j < 0 {
			longest[i] = i + 1
		} else {
			longest[i] = i - stack[j]
		}

		stack = stack[:j+1]
		stack = append(stack, i)
	}

	return longest
}

func LargestRectangleUnderSkyline(hs []int) int {
	reversed := func(arr []int) []int {
		n := len(arr)

		res := make([]int, n)

		for i := 0; i < n; i++ {
			res[i] = arr[n-1-i]
		}

		return res
	}

	left := largestRectangleHelper(hs)

	right := reversed(largestRectangleHelper(reversed(hs)))

	maximum := 0

	for i, h := range hs {
		if area := h * (left[i] + right[i] - 1); area > maximum {
			maximum = area
		}
	}

	return maximum
}
