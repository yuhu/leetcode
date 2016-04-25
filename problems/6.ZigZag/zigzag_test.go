package zigzag

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		want    string
	}{
		{
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		{
			"AB",
			1,
			"AB",
		},
	}
	for _, test := range tests {
		if got := convert(test.s, test.numRows); got != test.want {
			t.Errorf("convert(%q, %d) = %q, want %q", test.s, test.numRows, got, test.want)
		}
	}
}
