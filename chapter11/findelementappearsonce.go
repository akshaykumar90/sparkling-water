// Problem 11.17

package chapter11

func FindElementAppearsOnce(arr []int) int {
	var ones, twos int

	for _, a := range arr {
		nextOnes := (^a & ones) | (a & ^ones & ^twos)
		nextTwos := (^a & twos) | (a & ones)
		ones, twos = nextOnes, nextTwos
	}

	return ones
}
