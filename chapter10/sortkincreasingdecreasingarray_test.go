package chapter10

import "testing"

func TestSortKIncreasingDecreasingArray(t *testing.T) {
	arr := []int{57, 131, 493, 294, 221, 339, 418, 452, 442, 190}

	result := SortKIncreasingDecreasingArray(arr)

	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			t.Errorf("SortKIncreasingDecreasingArray(%v) => %v, unexpected at [%d]%d, [%d]%d",
				arr, result, i-1, result[i-1], i, result[i])
		}
	}
}
