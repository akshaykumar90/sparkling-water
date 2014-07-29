// problem 12.7

package chapter12

import "sort"

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

func Anagrams(words []string) [][]string {
	hash := make(map[string][]string)

	res := make([][]string, 0)

	for _, w := range words {
		hash[sorted(w)] = append(hash[sorted(w)], w)
	}

	for _, v := range hash {
		res = append(res, v)
	}

	return res
}
