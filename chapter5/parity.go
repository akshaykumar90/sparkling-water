// problem 5.1

package chapter5

var precomputedParity = buildParityTable()

func buildParityTable() []uint8 {
	table := make([]uint8, 1 << 16)

	for i := 0; i < (1 << 16); i++ {
		table[i] = uint8(parity(i))
	}

	return table
}

func parity(n int) int {
	var p int
	
	for ; n > 0; n &= n-1 {
		p ^= 1
	}

	return p
}

func Parity(n uint64) uint8 {
	return precomputedParity[n & (1<<16 - 1)] ^ 
		precomputedParity[(n >> 16) & (1<<16 - 1)] ^
		precomputedParity[(n >> 32) & (1<<16 - 1)] ^
		precomputedParity[(n >> 48) & (1<<16 - 1)]
}
