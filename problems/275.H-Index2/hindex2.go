package hindex2

import "sort"

func hIndex(citations []int) int {
	h := sort.Search(len(citations), func(i int) bool {
		return citations[i] >= len(citations)-i
	})
	return len(citations) - h
}
