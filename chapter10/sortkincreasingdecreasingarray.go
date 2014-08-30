// problem 10.2

package chapter10

func sgn(a, b int) int {
	if a < b {
		return 1
	} else {
		return -1
	}
}

func reverse(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

func SortKIncreasingDecreasingArray(arr []int) []int {
	arrs := make([][]int, 0)
	currSgn := sgn(arr[0], arr[1])
	begin := 0
	for i := 1; i < len(arr); i++ {
		if sgn(arr[i-1], arr[i]) != currSgn {
			if currSgn == -1 {
				reverse(arr[begin:i])
			}
			arrs = append(arrs, arr[begin:i])
			begin = i
			currSgn = sgn(arr[i-1], arr[i])
		}
	}
	if currSgn == -1 {
		reverse(arr[begin:])
	}
	arrs = append(arrs, arr[begin:])
	return MergeSortedArrays(arrs)
}
