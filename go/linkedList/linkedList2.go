package linkedlist

func (linked_list *linkedList) ReverseLinkedList02() {
	var prev, curr, next *linkedListNode

	curr = linked_list.head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	linked_list.head = prev
}

func mergeLinkedList02(linked_list_01, linked_list_02 *linkedList) *linkedList {
	merged_linked_list := &linkedList{}
	var node_01, node_02 *linkedListNode = linked_list_01.head, linked_list_02.head

	addNodeToMergedLinkedList := func(value int) {
		new_node := &linkedListNode{value: value, next: nil}
		merged_linked_list.size += 1

		curr := merged_linked_list.head

		if merged_linked_list.head == nil {
			merged_linked_list.head = new_node
			return
		}

		for curr != nil && curr.next != nil {
			curr = curr.next
		}

		curr.next = new_node

	}

	for node_01 != nil && node_02 != nil {
		if node_01.value <= node_02.value {
			addNodeToMergedLinkedList(node_01.value)
			node_01 = node_01.next
		} else {
			addNodeToMergedLinkedList(node_02.value)
			node_02 = node_02.next
		}
	}

	for node_01 != nil {
		addNodeToMergedLinkedList(node_01.value)
		node_01 = node_01.next
	}

	for node_02 != nil {
		addNodeToMergedLinkedList(node_02.value)
		node_02 = node_02.next
	}

	return merged_linked_list
}

func (linked_list *linkedList) hasCycle02() bool {
	var slow, fast *linkedListNode = linked_list.head, linked_list.head.next

	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}
	return false
}

func (linked_list *linkedList) middleNode02() int {
	var fast, slow *linkedListNode = linked_list.head, linked_list.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.value
}

func (linked_list *linkedList) removeDuplicates02() {
	curr := linked_list.head

	for curr != nil && curr.next != nil {
		if curr.value == curr.next.value {
			curr.next = curr.next.next
			linked_list.size -= 1
			continue
		}
		curr = curr.next
	}
}

func (linked_list *linkedList) removeElements02(value int) {
	curr := linked_list.head

	if curr == nil {
		return
	}

	for curr != nil && curr.next != nil {
		if curr.next.value == value {
			curr.next = curr.next.next
			linked_list.size -= 1
			continue
		}
		curr = curr.next
	}

	if linked_list.head.value == value {
		linked_list.head = linked_list.head.next
		linked_list.size -= 1
	}
}
