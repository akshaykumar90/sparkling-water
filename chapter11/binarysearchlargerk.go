//problem 11.2
package chapter11

func BinarySearchLargerK(arr []int, k int) int {
	// arr[:lo] has all x <= k
	// arr[hi:] has all x > k
	lo, hi := 0, len(arr)

	for lo < hi {
		mid := (lo + hi - 1) / 2
		if arr[mid] <= k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if hi < len(arr) {
		return hi
	} else {
		return -1
	}

}
