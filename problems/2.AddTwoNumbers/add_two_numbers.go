package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func newList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	tail := head
	for i := 1; i < len(vals); i++ {
		node := &ListNode{Val: vals[i]}
		tail.Next = node
		tail = node
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	var vals []int
	carry := 0
	for l1 != nil || l2 != nil {
		num := carry
		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}
		if num >= 10 {
			num -= 10
			carry = 1
		} else {
			carry = 0
		}
		vals = append(vals, num)
	}
	if carry > 0 {
		vals = append(vals, carry)
	}
	return newList(vals...)
}
