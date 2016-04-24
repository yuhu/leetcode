package longestsubstr

func lengthOfLongestSubstring(s string) int {
	charIdx := make(map[rune]int)
	maxLen, lastIdx := 0, -1
	for i, c := range s {
		if idx, ok := charIdx[c]; ok {
			if idx > lastIdx {
				lastIdx = idx
			}
		}
		if currLen := i - lastIdx; currLen > maxLen {
			maxLen = currLen
		}
		charIdx[c] = i
	}
	return maxLen
}
