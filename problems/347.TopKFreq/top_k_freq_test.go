package topkfreq

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
	}
	for _, test := range tests {
		if got := topKFrequent(test.nums, test.k); !reflect.DeepEqual(got, test.want) {
			t.Errorf("topKFrequent(%v, %d) = %v, want %v", test.nums, test.k, got, test.want)
		}
	}
}
