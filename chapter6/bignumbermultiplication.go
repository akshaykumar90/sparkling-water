// problem 6.9

package chapter6

func BigNumberMultiplication(fst, snd string) int {
	p := 0

	sign := 1
	if fst[0] == '-' {
		sign *= -1
		fst = fst[1:]
	}
	if snd[0] == '-' {
		sign *= -1
		snd = snd[1:]
	}

	for _, f := range fst {
		fd := int(f - '0')
		pp := 0
		for _, s := range snd {
			sd := int(s - '0')
			pp = pp*10 + fd*sd
		}
		p = p*10 + pp
	}

	return p * sign
}
