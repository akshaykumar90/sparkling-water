// Problem 17.1

package chapter17

func TiesElection(V []int) int {
	total := 0
	for _, v := range V {
		total += v
	}

	if total&1 != 0 {
		return 0
	}

	table := make([]int, total+1)

	table[0] = 1

	for _, v := range V {
		for j := total; j >= v; j-- {
			table[j] += table[j-v]
		}
	}

	return table[total>>1]
}
