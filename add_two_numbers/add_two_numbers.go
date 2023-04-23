package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return AddTwoNumbers(l1, l2)
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	newVal := l1.Val + l2.Val
	var head *ListNode
	carryOver := 0
	if newVal > 9 {
		head = &ListNode{Val: newVal - 10}
		carryOver = 1
	} else {
		head = &ListNode{Val: newVal}
	}
	l1 = l1.Next
	l2 = l2.Next
	tail := head

	zeroValue := &ListNode{}
	for l1 != nil || l2 != nil {
		if carryOver == 0 {
			if l1 == nil {
				tail.Next = l2
				return head
			}
			if l2 == nil {
				tail.Next = l1
				return head
			}
		}

		if l1 == nil {
			l1 = zeroValue
		}
		if l2 == nil {
			l2 = zeroValue
		}

		newVal = l1.Val + l2.Val + carryOver
		if newVal > 9 {
			tail.Next = &ListNode{Val: newVal - 10}
			carryOver = 1
		} else {
			tail.Next = &ListNode{Val: newVal}
			carryOver = 0
		}
		l1 = l1.Next
		l2 = l2.Next
		tail = tail.Next
	}

	if carryOver != 0 {
		tail.Next = &ListNode{Val: 1}
	}

	return head
}
