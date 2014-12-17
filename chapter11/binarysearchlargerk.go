// Problem 11.2

package chapter11

func BinarySearchLargerK(arr []int, k int) int {
	lo, hi, res := 0, len(arr)-1, -1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] <= k {
			lo = mid + 1
		} else {
			res = mid
			hi = mid - 1
		}
	}

	return res
}
