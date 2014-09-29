// Problem 6.12

package chapter6

func reverseInts(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func NextPermutation(p []int) []int {
	i := len(p) - 2
	for ; i >= 0 && p[i] >= p[i+1]; i-- {
	}

	if i == -1 {
		return nil
	} else {
		j := i + 1
		for ; j < len(p) && p[j] > p[i]; j++ {
		}

		p[i], p[j-1] = p[j-1], p[i]

		reverseInts(p[i+1:])

		return p
	}
}
