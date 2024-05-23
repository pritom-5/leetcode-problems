package linkedlist

func (ll *linkedList) ReverseLinkedList() {
	var prev, curr, next *linkedListNode
	curr = ll.head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	ll.head = prev
}

func mergeLinkedList(ll_1, ll_2 *linkedList) *linkedList {
	var ll_1_node, ll_2_node *linkedListNode = ll_1.head, ll_2.head
	new_ll_slice := make([]int, 0)

	for ll_1_node != nil && ll_2_node != nil {
		if ll_1_node.value >= ll_2_node.value {
			new_ll_slice = append(new_ll_slice, ll_2_node.value)
			ll_2_node = ll_2_node.next

		} else {
			new_ll_slice = append(new_ll_slice, ll_1_node.value)
			ll_1_node = ll_1_node.next
		}
	}

	for ll_1_node != nil {
		new_ll_slice = append(new_ll_slice, ll_1_node.value)
		ll_1_node = ll_1_node.next
	}
	for ll_2_node != nil {
		new_ll_slice = append(new_ll_slice, ll_2_node.value)
		ll_2_node = ll_2_node.next
	}

	return createLinkedListFromSlice(new_ll_slice)
}

func IsPalindrome(nums []int) bool {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left] != nums[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func IsPalindromeLinkedList(linked_list *linkedList) bool {
	linked_list_value_slice := make([]int, 0)

	curr := linked_list.head

	for curr != nil {
		linked_list_value_slice = append(linked_list_value_slice, curr.value)
		curr = curr.next
	}

	return IsPalindrome(linked_list_value_slice)
}

func (linked_list *linkedList) removeElements(val int) {
	curr := linked_list.head

	if curr == nil {
		return
	}

	for curr.next != nil {
		if curr.next.value == val {
			curr.next = curr.next.next
			linked_list.size -= 1
			continue
		}

		curr = curr.next
	}

	if curr.value == val {
		linked_list.head = curr.next
		linked_list.size -= 1
	}
}

func (linked_list *linkedList) removeDuplicates() {
	curr := linked_list.head

	if curr == nil || curr.next == nil {
		return
	}

	for curr != nil {
		if curr.next != nil && curr.value == curr.next.value {
			curr.next = curr.next.next
			linked_list.size -= 1
			continue
		}
		curr = curr.next
	}
}

func (linked_list *linkedList) middleNode() int {
	if linked_list.head == nil || linked_list.head.next == nil {
		return linked_list.head.value
	}

	var slow, fast *linkedListNode = linked_list.head, linked_list.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.value
}

func (linked_list *linkedList) HasCycle() bool {
	if linked_list.head == nil || linked_list.head.next == nil {
		return false
	}
	slow, fast := linked_list.head, linked_list.head.next

	for fast.next != nil && fast.next.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}
	return false
}
