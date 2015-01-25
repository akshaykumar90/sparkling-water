// Problem 11.16

package chapter11

func FindMissingAndDuplicate(arr []int) (int, int) {
	xor := 0
	for i, a := range arr {
		xor ^= i ^ a
	}

	differBit := xor & ^(xor - 1)

	x := 0
	for i, a := range arr {
		if i&differBit > 0 {
			x ^= i
		}
		if a&differBit > 0 {
			x ^= a
		}
	}

	for _, a := range arr {
		if x == a {
			return x ^ xor, x
		}
	}

	return x, x ^ xor
}
