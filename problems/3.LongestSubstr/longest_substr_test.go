package longestsubstr

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
		{
			"abba",
			2,
		},
	}
	for _, test := range tests {
		if got := lengthOfLongestSubstring(test.s); got != test.want {
			t.Errorf("lengthOfLongestSubstring(%q) = %v, want %v", test.s, got, test.want)
		}
	}
}
