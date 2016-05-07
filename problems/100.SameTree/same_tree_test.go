package sametree

import "testing"

func newTree(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func TestIsSame(t *testing.T) {
	tests := []struct {
		p, q *TreeNode
		want bool
	}{
		{
			nil,
			nil,
			true,
		},
		{
			nil,
			newTree(1, nil, nil),
			false,
		},
		{
			newTree(2, nil, newTree(3, nil, nil)),
			newTree(2, newTree(3, nil, nil), nil),
			false,
		},
		{
			newTree(4, newTree(5, nil, nil), nil),
			newTree(4, newTree(5, nil, nil), nil),
			true,
		},
	}
	for _, test := range tests {
		if got := isSameTree(test.p, test.q); got != test.want {
			t.Errorf("isSameTree(%v, %v) = %v, want %v", test.p, test.q, got, test.want)
		}
	}
}
