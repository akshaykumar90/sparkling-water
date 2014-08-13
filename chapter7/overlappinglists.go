// problem 7.5

package chapter7

func isSameCycle(fst, snd *Element) bool {
	curr := snd.Next
	for curr != snd {
		if curr == fst {
			return true
		}
		curr = curr.Next
	}
	return false
}

func OverlappingLists(fst, snd *Element) *Element {
	fstCycle := CheckCycle(fst)
	sndCycle := CheckCycle(snd)

	switch {
	case fstCycle == nil && sndCycle == nil:
		return OverlappingListsNoCycle(fst, snd)
	case fstCycle == nil && sndCycle != nil:
		fallthrough
	case fstCycle != nil && sndCycle == nil:
		return nil
	default:
		if isSameCycle(fstCycle, sndCycle) {
			return fstCycle
		} else {
			return nil
		}
	}
}
