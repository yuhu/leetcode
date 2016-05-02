package palindromenum

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{0, true},
		{1, true},
		{10, false},
		{11, true},
		{202, true},
		{-313, false},
	}
	for _, test := range tests {
		if got := isPalindrome(test.x); got != test.want {
			t.Errorf("isPalindrome(%d) = %v, want %v", test.x, got, test.want)
		}
	}
}
