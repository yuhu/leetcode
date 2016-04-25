package intbreak

import "testing"

func TestIntegerBreak(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 6},
		{10, 36},
	}
	for _, test := range tests {
		if got := integerBreak(test.n); got != test.want {
			t.Errorf("integerBreak(%d) = %d, want %d", test.n, got, test.want)
		}
	}
}
