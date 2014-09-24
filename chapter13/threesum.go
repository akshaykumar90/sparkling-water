// Problem 13.14

package chapter13

import "sort"

func ThreeSum(a []int, t int) bool {
	sort.Ints(a)

	for _, n := range a {
		s := t - n
		for i, j := 0, len(a)-1; i <= j; {
			if sum := a[i] + a[j]; sum == s {
				return true
			} else if sum < s {
				i++
			} else {
				j--
			}
		}
	}

	return false
}
