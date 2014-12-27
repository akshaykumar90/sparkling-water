// Problem 9.15

package chapter9

type TrieNode struct {
	r        rune
	children map[rune]*TrieNode
}

func (n *TrieNode) Insert(word string) {
	curr := n

	for _, r := range word {
		if _, ok := curr.children[r]; !ok {
			curr.children[r] = &TrieNode{r, make(map[rune]*TrieNode)}
		}
		curr = curr.children[r]
	}
}

func (n *TrieNode) ShortestPrefix(s string) int {
	curr := n

	for i, r := range s {
		if _, ok := curr.children[r]; !ok {
			return i
		}
		curr = curr.children[r]
	}

	return -1
}

func ShortestUniquePrefix(s string, D []string) string {
	root := &TrieNode{children: make(map[rune]*TrieNode)}

	for _, word := range D {
		root.Insert(word)
	}

	return s[:root.ShortestPrefix(s)+1]
}
