package palindromicsubstr

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"",
			"",
		},
		{
			"a",
			"a",
		},
		{
			"bb",
			"bb",
		},
		{
			"ccbabd",
			"bab",
		},
		{
			"ccabbadd",
			"abba",
		},
	}
	for _, test := range tests {
		if got := longestPalindrome(test.s); got != test.want {
			t.Errorf("longestPalindrome(%q) = %q, want %q", test.s, got, test.want)
		}
	}
}
