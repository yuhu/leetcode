package hindex2

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		want      int
	}{
		{
			[]int{0, 1, 3, 5, 6},
			3,
		},
		{
			nil,
			0,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{1},
			1,
		},
	}
	for _, test := range tests {
		if got := hIndex(test.citations); got != test.want {
			t.Errorf("hIndex(%v) = %d, want %d", test.citations, got, test.want)
		}
	}
}
