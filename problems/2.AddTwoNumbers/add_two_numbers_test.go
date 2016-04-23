package addtwonumbers

import (
	"strconv"
	"strings"
	"testing"
)

func sameList(lhs, rhs *ListNode) bool {
	if lhs == nil && rhs == nil {
		return true
	}
	if lhs == nil || rhs == nil {
		return false
	}
	return lhs.Val == rhs.Val && sameList(lhs.Next, rhs.Next)
}

func (l *ListNode) String() string {
	var vals []string
	for {
		if l == nil {
			break
		}
		vals = append(vals, strconv.Itoa(l.Val))
		l = l.Next
	}
	return strings.Join(vals, " -> ")
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   newList(2, 4, 3),
			l2:   newList(5, 6, 4),
			want: newList(7, 0, 8),
		},
	}
	for _, test := range tests {
		if got := addTwoNumbers(test.l1, test.l2); !sameList(got, test.want) {
			t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", test.l1, test.l2, got, test.want)
		}
	}
}
