package chapter11

import "testing"

func TestCompletionSearch(t *testing.T) {
	var completionSearchTests = []struct {
		a        []int
		total    int
		expected int
	}{
		{
			[]int{20, 30, 40, 90, 100},
			210,
			60,
		},
	}

	for _, tt := range completionSearchTests {
		actual := CompletionSearch(tt.a, tt.total)
		if actual != tt.expected {
			t.Errorf("CompletionSearch(%v, %d): expected %d, actual %d",
				tt.a, tt.total, tt.expected, actual)
		}
	}
}
