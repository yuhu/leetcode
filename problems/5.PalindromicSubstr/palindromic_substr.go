package palindromicsubstr

func palindrome(s string, left int, right int) string {
	for s[left] == s[right] {
		left -= 1
		right += 1
		if left < 0 || right >= len(s) {
			break
		}
	}
	return s[left+1 : right]
}

func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		pal := palindrome(s, i, i)
		if len(pal) > len(longest) {
			longest = pal
		}
		if i < len(s)-1 {
			pal = palindrome(s, i, i+1)
			if len(pal) > len(longest) {
				longest = pal
			}
		}
	}
	return longest
}
