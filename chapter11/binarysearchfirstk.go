// problem 11.1

package chapter11

func BinarySearchFirstK(arr []int, k int) int {
	lo, hi := 0, len(arr)-1
	last := -1
	for lo <= hi {
		mid := (lo + hi) / 2
		if arr[mid] == k {
			last = mid
			hi = mid - 1
		} else if arr[mid] < k {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return last
}
