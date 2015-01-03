package common

func ReverseInts(arr []int) {
	n := len(arr)

	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

func MinInt(a, b int) int {
	if b < a {
		return b
	} else {
		return a
	}
}

func MaxInt(a, b int) int {
	if b > a {
		return b
	} else {
		return a
	}
}
