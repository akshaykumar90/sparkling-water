// problem 7.1

package chapter7

type Element struct {
	Value int
	Next  *Element
}

func MergeSortedLists(fst *Element, snd *Element) *Element {
	root_parent := Element{}
	cur := &root_parent
	for fst != nil && snd != nil {
		if fst.Value <= snd.Value {
			cur.Next = fst
			fst = fst.Next
		} else {
			cur.Next = snd
			snd = snd.Next
		}
		cur = cur.Next
	}
	if fst != nil {
		cur.Next = fst
	} else {
		cur.Next = snd
	}
	return root_parent.Next
}
