package linkedlist

func (head *linkedListNode) remove (val int) *linkedListNode {
	if head == nil {
		return head
	}

	var prev, curr *linkedListNode = head, head.next

	for curr != nil {
		if curr.value == val {
			prev.next = curr.next
		}
		prev = curr
		curr = curr.next
	}

	if head.value == val {
		head = head.next
	}
	return head
}
