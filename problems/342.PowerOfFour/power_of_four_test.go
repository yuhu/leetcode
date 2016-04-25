package poweroffour

import "testing"

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{4, true},
		{8, false},
		{15, false},
		{16, true},
		{-1, false},
		{-16, false},
	}
	for _, test := range tests {
		if got := isPowerOfFour(test.x); got != test.want {
			t.Errorf("isPowerOfFour(%d) = %v, want %v", test.x, got, test.want)
		}
	}
}
