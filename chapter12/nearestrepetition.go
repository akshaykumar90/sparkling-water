// Problem 12.3

package chapter12

import "math"

func NearestRepetition(S []string) int {
	hash := make(map[string]int)

	closest := math.MaxInt32

	for i, s := range S {
		if p, ok := hash[s]; ok {
			dist := i - p
			if dist < closest {
				closest = dist
			}
		}
		hash[s] = i
	}

	return closest
}
