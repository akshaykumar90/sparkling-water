//problem 11.3

package chapter11

func BinarySearchAiEqI(a []int) int {
	lo, hi := 0, len(a)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		if a[mid] == mid {
			return mid
		} else if a[mid] < mid {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1
}
