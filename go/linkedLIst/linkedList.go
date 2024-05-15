package linkedlist

type linkedListNode struct {
	value int
	next  *linkedListNode
}

type linkedList struct {
	head *linkedListNode
	size int
}

func (ll *linkedList) AddNode(val int) {
	ll.size += 1

	new_node := &linkedListNode{value: val, next: nil}

	if ll.head == nil {
		ll.head = new_node
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}

	current.next = new_node
}

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

func createLinkedListFromSlice(values []int) *linkedList {
	linked_list := &linkedList{}
	for _, val := range values {
		linked_list.AddNode(val)
	}
	return linked_list
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

func (linked_list *linkedList) RemoveElements(val int) {
	var curr, prev *linkedListNode
	curr = linked_list.head

	for curr != nil {
		if curr.value == val {
			linked_list.size -= 1
			prev.next = curr.next
		}
		prev = curr
		curr = curr.next
	}



}
