// Problem 12.11

package chapter12

func SearchFrequentItems(arr []int, k int) []int {
	freq := make(map[int]int)

	for _, x := range arr {
		freq[x]++
		if len(freq) > k {
			for a := range freq {
				freq[a]--
				if freq[a] == 0 {
					delete(freq, a)
				}
			}
		}
	}

	for a := range freq {
		freq[a] = 0
	}

	for _, x := range arr {
		if _, ok := freq[x]; ok {
			freq[x]++
		}
	}

	n := len(arr)
	result := make([]int, 0)

	for a, f := range freq {
		if f >= n/k {
			result = append(result, a)
		}
	}

	return result
}
