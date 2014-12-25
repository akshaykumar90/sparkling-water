// Problem 5.10

package chapter5

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	if a == 0 {
		return b
	}

	if a&1 == 0 && b&1 == 0 {
		return GCD(a>>1, b>>1) << 1
	} else if a&1 == 0 && b&1 != 0 {
		return GCD(a>>1, b)
	} else if a&1 == 1 && b&1 == 0 {
		return GCD(a, b>>1)
	} else if a > b {
		return GCD(a-b, b)
	} else {
		return GCD(a, b-a)
	}
}
