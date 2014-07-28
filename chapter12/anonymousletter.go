// problem 12.9

package chapter12

func AnonymousLetter(L string, M string) bool {
	dict := make(map[rune]int)

	for _, r := range L {
		dict[r]++
	}

	for _, r := range M {
		if dict[r] > 0 {
			dict[r]--
			if dict[r] == 0 {
				delete(dict, r)
				if len(dict) == 0 {
					return true
				}
			}
		}
	}

	return len(dict) == 0
}
