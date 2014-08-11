// problem 6.3

package chapter6

import "math"

func RobotBattery(hs []int) int {
	minH := math.MaxInt32
	maxDiff := 0

	for _, h := range hs {
		if h < minH {
			minH = h
		}
		if h-minH > maxDiff {
			maxDiff = h - minH
		}
	}

	return maxDiff
}
