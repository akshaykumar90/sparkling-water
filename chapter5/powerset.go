// problem 5.5

package chapter5

import "fmt"

func powerSetHelper(S []int, idx int, res []int) {
	if idx == len(S) {
		fmt.Println(res)
	} else {
		powerSetHelper(S, idx+1, res)
		res = append(res, S[idx])
		powerSetHelper(S, idx+1, res)
		res = res[:len(res)-1]
	}
}

func PowerSet(S []int) {
	powerSetHelper(S, 0, make([]int, 0, len(S)))
}
