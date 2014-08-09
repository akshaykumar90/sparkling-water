// problem 5.2

package chapter5

func SwapBits(n uint64, i uint, j uint) uint64 {
	if n>>i&1 != n>>j&1 {
		n ^= 1<<i | 1<<j
	}

	return n
}
