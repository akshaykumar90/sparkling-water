// problem 12.15

package chapter12

type CoverValue struct {
	start  int
	length int
}

func SubseqCover(A, Q []string) []string {
	bag := make(map[string]int)
	for i, q := range Q {
		bag[q] = i
	}

	arr := make([]CoverValue, len(Q))

	for p, w := range A {
		if idx, ok := bag[w]; ok {
			if idx == 0 {
				arr[idx] = CoverValue{p, 1}
			} else if arr[idx-1].length > 0 {
				currCoverLength := p - arr[idx-1].start + 1
				if arr[idx].length == 0 || currCoverLength < arr[idx].length {
					arr[idx] = CoverValue{arr[idx-1].start, currCoverLength}
				}
			}
		}
	}

	last := arr[len(Q)-1]
	if last.length > 0 {
		return A[last.start : last.start+last.length]
	} else {
		return nil
	}
}
