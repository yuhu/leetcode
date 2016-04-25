package reversestr

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
	}
	for _, test := range tests {
		if got := reverseString(test.s); got != test.want {
			t.Errorf("reverseString(%q) = %q, want %q", test.s, got, test.want)
		}
	}
}
