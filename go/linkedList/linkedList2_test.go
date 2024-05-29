package linkedlist

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T)  {
	list_1 := createLinkedListFromSlice([]int{7, 7, 7, 7})

	mod_list_1_head := list_1.head.remove(7)
	mod_list_1_slice := mod_list_1_head.next.sliceFromLinkedList()
	exp_1 := []int{}

	if !reflect.DeepEqual(mod_list_1_slice, exp_1) {
		t.Errorf("got: %#v, exp: %#v", mod_list_1_slice, exp_1)
	}
	
	
}
