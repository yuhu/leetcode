package reverseint

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{123, 321},
		{-123, -321},
		{100, 1},
		{-100, -1},
		{1534236469, 0},
	}
	for _, test := range tests {
		if got := reverse(test.x); got != test.want {
			t.Errorf("reverse(%d) = %d, want %d", test.x, got, test.want)
		}
	}
}
