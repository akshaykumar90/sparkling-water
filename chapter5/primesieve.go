// problem 5.11

package chapter5

func PrimeSieve(n int) []int {
	table := make([]bool, n+1)
	for i := range table {
		table[i] = true
	}

	primes := make([]int, 0)

	for i := 2; i < len(table); i++ {
		if table[i] {
			primes = append(primes, i)
			for j := i * i; j < len(table); j += i {
				table[j] = false
			}
		}
	}

	return primes
}
