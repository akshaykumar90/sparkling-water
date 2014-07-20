// problem 9.14
package chapter9

func LowestCommonAncestorHash(fst, snd *TreeNodeWithParent) *TreeNodeWithParent {
	visited := make(map[*TreeNodeWithParent]bool)

	for fst != nil || snd != nil {
		if fst != nil {
			if visited[fst] {
				return fst
			}

			visited[fst] = true
			fst = fst.Parent
		}

		if snd != nil {
			if visited[snd] {
				return snd
			}

			visited[snd] = true
			snd = snd.Parent
		}
	}

	return nil
}
