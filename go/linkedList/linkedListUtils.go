package linkedlist

func createLinkedListFromSlice(values []int) *linkedList {
	linked_list := &linkedList{}
	for _, val := range values {
		linked_list.AddNode(val)
	}
	return linked_list
}


func createCyclicLinkedListFromSlice(values []int, pos int) *linkedList {
	if len(values) == 0 {
		return &linkedList{}
	}

	var cyclic_pos *linkedListNode

	linked_list:= &linkedList{head: &linkedListNode{value: values[0], next: nil}, size: 1}
	curr := linked_list.head

	for i, val := range values {
		if i == 0 {
			continue
		}

		new_node := &linkedListNode{value: val, next: nil}

		if (i == pos) {
			cyclic_pos = new_node
		}

		curr.next = new_node
		curr = curr.next
		linked_list.size += 1
	}

	curr.next = cyclic_pos

	return linked_list
}


func (linked_list *linkedList) sliceFromLinkedList () []int {
	slice := make([]int, 0)

	curr := linked_list.head

	for curr != nil {
		slice = append(slice, curr.value)
		curr = curr.next
	}

	return slice
}
