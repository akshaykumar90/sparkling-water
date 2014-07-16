package chapter7

func OverlappingListsNoCycle(fst, snd *Element) *Element {
	getLastElement := func (head *Element) (int, *Element) {
		l := 0
		for ; head != nil; l, head = l+1, head.Next {}
		return l, head
	}
	
	l1, last1 := getLastElement(fst)
	l2, last2 := getLastElement(snd)

	if last1 == last2 {
		var slow, fast *Element
		var diff int
		if l1 < l2 {
			slow, fast = snd, fst
			diff = l2-l1
		} else {
			slow, fast = fst, snd
			diff = l1-l2
		}

		for ; diff > 0; diff-- {
			slow = slow.Next
		}

		for ; slow != fast; slow, fast = slow.Next, fast.Next {}

		return slow

	} else {
		return nil
	}
}
