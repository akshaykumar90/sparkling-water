// problem 12.8

package chapter12

func CanStringBePalindrome(s string) bool {
	hash := make(map[rune]bool)
	for _, r := range s {
		hash[r] = !hash[r]
	}

	cnt := 0
	for _, v := range hash {
		if v == true {
			cnt += 1
		}
	}

	return cnt <= 1
}
