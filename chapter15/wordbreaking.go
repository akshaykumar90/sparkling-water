// problem 15.12

package chapter15

func WordBreaking(dict map[string]bool, s string) bool {
	suffix := make([]int, len(s))
	for i := range suffix {
		suffix[i] = -1
	}

	for i := len(s); i >= 0; i-- {
		for j := i + 1; j <= len(s); j++ {
			if dict[s[i:j]] && (j == len(s) || suffix[j] != -1) {
				suffix[i] = j
			}
		}
	}

	return suffix[0] != -1
}
