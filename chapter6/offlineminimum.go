// Problem 6.8

package chapter6

func findSet(set []int, x int) int {
	if set[x] != x {
		set[x] = findSet(set, set[x])
	}
	return set[x]
}

func unionSet(set []int, x, y int) {
	rx, ry := findSet(set, x), findSet(set, y)

	var min, max int
	if rx < ry {
		min, max = rx, ry
	} else {
		min, max = ry, rx
	}

	set[min] = max
}

func OfflineMinimum(A, E []int) []int {
	R := make([]int, len(A))
	for i := range R {
		R[i] = len(E)
	}

	pre := 1
	for i, x := range E {
		for j := pre; j <= x; j++ {
			R[A[j-1]] = i
		}
		pre = x + 1
	}

	ret := make([]int, len(E))

	set := make([]int, len(E)+1)
	for i := range set {
		set[i] = i
	}

	for i := 0; i < len(A); i++ {
		if findSet(set, R[i]) != len(E) {
			ret[set[R[i]]] = i
			unionSet(set, set[R[i]], set[R[i]]+1)
		}
	}

	return ret
}
