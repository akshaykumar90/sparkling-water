package chapter6

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestOfflineMinimum(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	size_A, size_E := 100, 10

	for tc := 0; tc < 100; tc++ {
		A := rand.Perm(size_A)

		E := make([]int, size_E)
		for i := range E {
			E[i] = rand.Intn(size_A) + 1
		}

		sort.Ints(E)

		expected := make([]int, size_E)

		curr := make([]int, 0, len(A))
		for i := 0; i < size_E; i++ {
			// read new elements
			if i == 0 {
				curr = append(curr, A[:E[i]]...)
			} else {
				curr = append(curr, A[E[i-1]:E[i]]...)
			}
			sort.Ints(curr)
			expected[i] = curr[0] // curr[0] is min element
			curr = curr[1:]       // remove min element
		}

		actual := OfflineMinimum(A, E)

		for i, v := range expected {
			if v != actual[i] {
				t.Errorf("OfflineMinimum(%v, %v): at %d: expected %d, actual %d", A, E, i+1, v, actual[i])
			}
		}
	}
}
