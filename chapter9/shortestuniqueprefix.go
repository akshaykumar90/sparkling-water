// Problem 9.15

package chapter9

type TrieNode struct {
	r        rune
	children []*TrieNode
}

func (n *TrieNode) Insert(runes []rune) {
	if len(runes) > 0 {
		var child *TrieNode
		for _, v := range n.children {
			if v.r == runes[0] {
				child = v
				break
			}
		}
		if child == nil {
			child = &TrieNode{r: runes[0]}
			n.children = append(n.children, child)
		}
		child.Insert(runes[1:])
	}
}

func (n *TrieNode) ShortestPrefix(runes []rune, i int) int {
	if i < len(runes) {
		for _, v := range n.children {
			if v.r == runes[i] {
				return v.ShortestPrefix(runes, i+1)
			}
		}
	}
	return i
}

func ShortestUniquePrefix(s string, D []string) string {
	var root TrieNode

	for _, word := range D {
		root.Insert([]rune(word))
	}

	return s[:root.ShortestPrefix([]rune(s), 0)+1]
}
