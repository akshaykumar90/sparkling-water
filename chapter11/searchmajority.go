// Problem 11.19

package chapter11

func SearchMajority(arr []int) int {
	var cand int

	count := 0

	for _, a := range arr {
		if count == 0 {
			cand = a
			count = 1
		} else if a == cand {
			count++
		} else {
			count--
		}
	}

	return cand
}
