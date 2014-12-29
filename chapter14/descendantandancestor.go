// Problem 14.13

package chapter14

func DescendantAndAncestor(r, s, m *TreeNode) bool {
	fst, snd := r, s
	var foundFst, foundSnd bool

	for fst != nil || snd != nil {
		if fst != nil {
			if fst == m {
				foundFst = true
			} else if fst == s {
				return foundFst
			}

			if s.Value < fst.Value {
				fst = fst.Left
			} else {
				fst = fst.Right
			}
		}
		if snd != nil {
			if snd == m {
				foundSnd = true
			} else if snd == r {
				return foundSnd
			}

			if r.Value < snd.Value {
				snd = snd.Left
			} else {
				snd = snd.Right
			}
		}
	}

	return false
}
