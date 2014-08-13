// problem 6.21

package chapter6

func remove(ss []rune) []rune {
	j := 0
	for _, r := range ss {
		if r != 'b' {
			ss[j] = r
			j += 1
		}
	}
	return ss[:j]
}

func replace(ss []rune) []rune {
	cnt := 0
	for _, r := range ss {
		if r == 'a' {
			cnt++
		}
	}

	sz := len(ss) / 2
	newSz := sz + cnt

	for i := sz - 1; i >= 0 && cnt > 0; i-- {
		if ss[i] == 'a' {
			ss[i+cnt] = 'd'
			ss[i+cnt-1] = 'd'
			cnt--
		} else {
			ss[i+cnt] = ss[i]
		}
	}
	return ss[:newSz]
}

func ReplaceAndRemove(str string) string {
	strSlice := []rune(str)

	strSlice = remove(strSlice)

	biggerSlice := make([]rune, 2*len(strSlice))
	copy(biggerSlice, strSlice)
	strSlice = biggerSlice

	strSlice = replace(strSlice)

	return string(strSlice)
}
