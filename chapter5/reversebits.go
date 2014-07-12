package chapter5

func reverse(n uint16) uint16 {
	num := uint64(n)
	var i, j uint = 0, 15

	for i < j {
		num = SwapBits(num, i, j)
		i++
		j--
	}

	return uint16(num)
}

var precomputedReversed = buildReversedTable()

func buildReversedTable() []uint16 {
	table := make([]uint16, 1<<16)

	for i := 0; i < 1<<16; i++ {
		table[i] = reverse(uint16(i))
	}

	return table
}

func ReverseBits(n uint64) uint64 {
	bitmask := uint64(0xFFFF)
	return uint64(precomputedReversed[n&bitmask])<<48 |
		uint64(precomputedReversed[n>>16&bitmask])<<32 |
		uint64(precomputedReversed[n>>32&bitmask])<<16 |
		uint64(precomputedReversed[n>>48&bitmask])
}
