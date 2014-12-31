// Problem 12.14

package chapter12

import (
	"container/list"
	"math"
)

func SmallestSubarrayCoveringSet(A, Q []string) []string {
	queue := list.New()

	occ := make(map[string]*list.Element)

	for _, q := range Q {
		occ[q] = nil
	}

	n := len(occ)

	min, i, j := math.MaxInt32, 0, -1

	for p, w := range A {
		if _, ok := occ[w]; ok {
			if occ[w] != nil {
				queue.Remove(occ[w])
			}
			e := queue.PushBack(p)
			occ[w] = e

			if queue.Len() == n {
				front := queue.Front().Value.(int)
				if szCover := p - front + 1; szCover < min {
					i, j = front, p
					min = j - i + 1
				}
			}
		}
	}

	return A[i : j+1]
}
