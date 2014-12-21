// Problem 6.23

package chapter6

func matchSingle(r, ch rune) bool {
	return r == '.' || r == ch
}

func regularExpressionHelper(r, s string) bool {
	if len(r) == 0 {
		return true
	}

	if r == "$" {
		return len(s) == 0
	}

	if len(r) >= 2 && r[1] == '*' {
		for i, c := range s {
			if matchSingle(rune(r[0]), c) && regularExpressionHelper(r[2:], s[i+1:]) {
				return true
			}
		}

		return regularExpressionHelper(r[2:], s)
	}

	return len(s) > 0 && matchSingle(rune(r[0]), rune(s[0])) && regularExpressionHelper(r[1:], s[1:])
}

func RegularExpression(r, s string) bool {
	if r[0] == '^' {
		return regularExpressionHelper(r[1:], s)
	}

	for i := 0; i <= len(s); i++ {
		if regularExpressionHelper(r, s[i:]) {
			return true
		}
	}

	return false
}
