package remove_nth_node_from_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = reverse(head)
	return reverse(removeNthFromStart(head, n))
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newListHead *ListNode

	for head != nil {
		next := head.Next
		head.Next = newListHead
		newListHead = head
		head = next
	}
	return newListHead
}

func removeNthFromStart(head *ListNode, n int) *ListNode {
	if n == 1 {
		return head.Next
	}

	current := head
	var prev *ListNode
	for i := 1; i < n; i++ {
		if current == nil {
			break
		}
		prev = current
		current = current.Next
	}
	if current == nil || prev == nil {
		return head
	}
	prev.Next = current.Next
	return head
}
