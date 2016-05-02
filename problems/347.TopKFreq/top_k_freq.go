package topkfreq

import "sort"

type elem struct{ n, freq int }
type byFreq []elem

func (b byFreq) Len() int      { return len(b) }
func (b byFreq) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byFreq) Less(i, j int) bool {
	// We want to place more frequet elements before less frequent ones.
	return b[i].freq > b[j].freq
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n] += 1
	}
	var elems []elem
	for n, f := range freq {
		elems = append(elems, elem{n, f})
	}
	sort.Sort(byFreq(elems))
	top := make([]int, k)
	for i := 0; i < k; i++ {
		top[i] = elems[i].n
	}
	return top
}
