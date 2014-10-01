// Problem 5.4

package chapter5

func ClosestIntSameBits(x uint64) uint64 {
	for i := uint(0); i < 63; i++ {
		if ((x>>i)&1)^((x>>(i+1))&1) == 1 {
			return x ^ (1<<i | 1<<(i+1))
		}
	}
	return x
}
