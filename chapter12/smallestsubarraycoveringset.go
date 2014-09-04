// problem 12.14

package chapter12

import (
	"container/list"
	"math"
)

type MatchValue struct {
	kw     string
	offset int
}

func SmallestSubarrayCoveringSet(A, Q []string) []string {
	bag := make(map[string]bool)
	for _, q := range Q {
		bag[q] = true
	}
	szBag := len(bag)

	queue := list.New()
	occ := make(map[string]int)

	min, i, j := math.MaxInt32, 0, -1

	for p, w := range A {
		if bag[w] {
			occ[w] += 1
			queue.PushBack(MatchValue{w, p})
			for len(occ) == szBag {
				front := queue.Remove(queue.Front()).(MatchValue)
				if szCover := p - front.offset + 1; szCover < min {
					i, j = front.offset, p
					min = j - i + 1
				}
				occ[front.kw] -= 1
				if occ[front.kw] == 0 {
					delete(occ, front.kw)
				}
			}
		}
	}
	return A[i : j+1]
}
