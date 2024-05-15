package linkedlist

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T)  {
	linked_list := linkedList{}
	linked_list.AddNode(10)
	linked_list.AddNode(20)

	value_1 := linked_list.head.value
	value_2 := linked_list.head.next.value

	if value_1 != 10 {
		t.Errorf("value_1: %d, exp: %d", value_1, 10)
	}
	if value_2 != 20 {
		t.Errorf("value_2: %d, exp: %d", value_2, 20)
	}
}

func TestReverseLinkedList(t *testing.T)  {
	linked_list_01 := linkedList{}
	fmt.Print(linked_list_01)
	linked_list_01.AddNode(10)
	linked_list_01.AddNode(20)
	linked_list_01.AddNode(30)

	linked_list_01.ReverseLinkedList()

	if linked_list_01.head.value != 30 {
		t.Errorf("got: %d, exp: %d", linked_list_01.head.value, 30)
	}
	if linked_list_01.head.next.value != 20 {
		t.Errorf("got: %d, exp: %d", linked_list_01.head.next.value, 20)
	}
	if linked_list_01.head.next.next.value != 10 {
		t.Errorf("got: %d, exp: %d", linked_list_01.head.next.next.value, 10)
	}
	
}
