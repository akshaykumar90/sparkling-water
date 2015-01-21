// Problem 17.3

package chapter17

func MinimumSubarrayDifference(arr []int) int {
	sum := 0
	for _, a := range arr {
		sum += a
	}

	table := make([]bool, sum+1)

	table[0] = true

	for _, item := range arr {
		for v := sum >> 1; v >= item; v-- {
			if table[v-item] {
				table[v] = true
			}
		}
	}

	for i := sum >> 1; i > 0; i-- {
		if table[i] {
			return (sum - i) - i
		}
	}

	return sum
}
