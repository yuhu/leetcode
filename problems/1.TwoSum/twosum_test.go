package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
	}
	for _, test := range tests {
		if got := twoSum(test.nums, test.target); !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", test.nums, test.target, got, test.want)
		}
	}
}
