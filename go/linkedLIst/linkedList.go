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
