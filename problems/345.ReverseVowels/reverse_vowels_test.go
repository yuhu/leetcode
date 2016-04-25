package reversevowels

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"a", "a"},
		{"hello", "holle"},
		{"leetcode", "leotcede"},
	}
	for _, test := range tests {
		if got := reverseVowels(test.s); got != test.want {
			t.Errorf("reverseVowels(%q) = %q, want %q", test.s, got, test.want)
		}
	}
}
