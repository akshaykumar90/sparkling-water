// Problem 11.7

package chapter11

func CompletionSearch(a []int, total int) int {
	curr := 0
	n := len(a)
	for k, s := range a {
		if curr+(n-k)*s > total {
			return (total - curr) / (n - k)
		} else {
			curr += s
		}
	}
	return 0
}
