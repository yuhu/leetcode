package hindex

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	for i, c := range citations {
		h := len(citations) - i
		if c >= h {
			return h
		}
	}
	return 0
}
