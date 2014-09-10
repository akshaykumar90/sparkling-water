// problem 15.2

package chapter15

func merge(arr []int, mid int) int {
	left := make([]int, mid)
	copy(left, arr[:mid])
	right := arr[mid:]

	inv := 0

	i, j, k := 0, 0, 0
	for ; i < len(left) && j < len(right); k++ {
		if left[i] > right[j] {
			arr[k] = right[j]
			j++
			inv += len(left) - i
		} else {
			arr[k] = left[i]
			i++
		}
	}

	if i < len(left) {
		copy(arr[k:], left[i:])
	}

	return inv
}

func CountInversions(arr []int) int {
	if len(arr) > 1 {
		mid := len(arr) / 2
		left := CountInversions(arr[:mid])
		right := CountInversions(arr[mid:])
		return left + right + merge(arr, mid)
	} else {
		return 0
	}
}
