// Problem 6.10

package chapter6

func PermutationArray(P, A []int) []int {
	for i := 0; i < len(A); i++ {
		if A[i] >= 0 {
			dest, src := i, P[i]
			start := A[dest]
			for A[src] >= 0 {
				A[dest] = -(A[src] + 1)
				dest, src = src, P[src]
			}
			A[dest] = -(start + 1)
		}
	}

	for i, a := range A {
		A[i] = -(a + 1)
	}

	return A
}
