// Problem 11.4

package chapter11

func searchPosPair(arr []int, k int) (int, int) {
	n := len(arr)

	lo, hi := 0, n-1

	for ; lo < n && arr[lo] < 0; lo++ {
	}
	for ; hi >= 0 && arr[hi] < 0; hi-- {
	}

	for lo < hi {
		if sum := arr[lo] + arr[hi]; sum == k {
			return lo, hi
		} else if sum < k {
			lo++
			for ; lo < n && arr[lo] < 0; lo++ {
			}
		} else {
			hi--
			for ; hi >= 0 && arr[hi] < 0; hi-- {
			}
		}
	}

	return -1, -1
}

func searchNegPair(arr []int, k int) (int, int) {
	n := len(arr)

	lo, hi := 0, n-1

	for ; lo < n && arr[lo] > 0; lo++ {
	}
	for ; hi >= 0 && arr[hi] > 0; hi-- {
	}

	for lo < hi {
		if sum := arr[lo] + arr[hi]; sum == k {
			return lo, hi
		} else if sum < k {
			hi--
			for ; hi >= 0 && arr[hi] > 0; hi-- {
			}
		} else {
			lo++
			for ; lo < n && arr[lo] > 0; lo++ {
			}
		}
	}

	return -1, -1
}

func searchPosNegPair(arr []int, k int) (int, int) {
	n := len(arr)

	lo, hi := n-1, n-1

	for ; lo >= 0 && arr[lo] >= 0; lo-- {
	}
	for ; hi >= 0 && arr[hi] < 0; hi-- {
	}

	for lo >= 0 && hi >= 0 {
		if sum := arr[lo] + arr[hi]; sum == k {
			return lo, hi
		} else if sum < k {
			lo--
			for ; lo >= 0 && arr[lo] >= 0; lo-- {
			}
		} else {
			hi--
			for ; hi >= 0 && arr[hi] < 0; hi-- {
			}
		}
	}

	return -1, -1
}

func SearchAPairSortedArray(arr []int, k int) (int, int) {
	cases := []func([]int, int) (int, int){
		searchPosNegPair,
		searchPosPair,
		searchNegPair,
	}

	for _, fn := range cases {
		x, y := fn(arr, k)
		if x != -1 && y != -1 {
			return x, y
		}
	}

	return -1, -1
}
