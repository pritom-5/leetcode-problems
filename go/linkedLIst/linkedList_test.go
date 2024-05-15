package linkedlist

import "testing"

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

