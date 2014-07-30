// problem 13.5

package chapter13

func IntersectSortedArrays(a, b []int) []int {
	c := make([]int, 0)

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			c = append(c, a[i])
			i++
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}

	return c
}
