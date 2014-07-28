// problem 11.3

package chapter11

import "errors"
import "math/rand"

func partition(a []int) int {
	p_i := rand.Intn(len(a))
	a[p_i], a[len(a)-1] = a[len(a)-1], a[p_i]
	p_i = len(a) - 1

	firstlow := 0

	for i, _ := range a[:p_i] {
		if a[i] > a[p_i] {
			a[i], a[firstlow] = a[firstlow], a[i]
			firstlow++
		}
	}

	a[p_i], a[firstlow] = a[firstlow], a[p_i]

	return firstlow
}

func KthLargestElement(a []int, k int) (int, error) {
	lo, hi := 0, len(a)

	for lo < hi {
		p_i := partition(a[lo:hi])

		if p_i == k-1 {
			return a[lo+p_i], nil
		} else if p_i < k-1 {
			lo += p_i + 1
			k -= p_i + 1
		} else {
			hi = lo + p_i
		}
	}

	return 0, errors.New("KthLargestElement: invalid input")
}
