// problem 13.6

package chapter13

import "sort"

func TeamPhoto1(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)

	for i := 0; i < len(a); i++ {
		if !(a[i] < b[i]) {
			return false
		}
	}

	return true
}
