// problem 11.5

package chapter11

func BinarySearchCircularArray(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		if mid := (lo + hi) / 2; arr[mid] > arr[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
