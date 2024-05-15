package linkedlist


type linkedListNode struct {
	value int
	next *linkedListNode
}

type linkedList struct {
	head *linkedListNode
	size int
}

func (ll *linkedList) AddNode (val int) {
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

func (ll *linkedList) ReverseLinkedList () {
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
