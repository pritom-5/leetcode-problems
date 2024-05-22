package linkedlist

import (
	"reflect"
	"testing"
)


func TestAddNode(t *testing.T)  {
	linked_list := createLinkedListFromSlice([]int{10, 20})


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
	linked_list_01 := createLinkedListFromSlice([]int{10, 20, 30})

	linked_list_01.ReverseLinkedList02()

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

func TestMergeLinkedList(t *testing.T)  {
	ll_1 := createLinkedListFromSlice([]int {1, 2, 4})
	ll_2 := createLinkedListFromSlice([]int {1, 3, 4})

	ll_3 := createLinkedListFromSlice([]int {1, 1, 2, 3, 4, 4})

	merged_linked_list := mergeLinkedList02(ll_1, ll_2)

	if !reflect.DeepEqual(ll_3, merged_linked_list) {
		t.Errorf("merged_linked_list: %#v", merged_linked_list)
	}
}

func TestIsPalindrome(t *testing.T)  {
	test_data_items := []struct{
		id string
		nums []int
		exp bool
	}{
		{id: "1", nums: []int{1, 2, 3, 2, 1}, exp: true},
		{id: "2", nums: []int{1, 2, 2, 1}, exp: true},
		{id: "3", nums: []int{1, 3,  2, 1}, exp: false},
		{id: "4", nums: []int{}, exp: true},
	}

	for _, item := range test_data_items {
		t.Run(item.id, func(t *testing.T) {
			got := IsPalindrome(item.nums)
			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}

		})
	}
	
}
	func TestIsPalindromeLinkedList(t *testing.T)  {
		test_items := []struct{
			id string
			linked_list *linkedList
			exp bool
		}{
			{id: "1", linked_list: createLinkedListFromSlice([]int{1, 2, 2, 1}), exp: true},
			{id: "2", linked_list: createLinkedListFromSlice([]int{1, 2}), exp: false},
		}

	for _, item := range test_items {
		t.Run(item.id, func(t *testing.T) {
			got := IsPalindromeLinkedList(item.linked_list)

			if got != item.exp {
				t.Errorf("got: %t, exp: %t", got, item.exp)
			}

		})
	}
}

func TestRemoveElements (t *testing.T) {
	test_items := []struct{
		id string
		val int
		linked_list, exp *linkedList
	} {
		{id: "1", linked_list: createLinkedListFromSlice([]int{1,2,6,3,4,5,6}), val: 6,	exp: createLinkedListFromSlice([]int{1, 2, 3, 4, 5})},
		{id: "2", linked_list: createLinkedListFromSlice([]int{}), val: 1,	exp: createLinkedListFromSlice([]int{})},
		{id: "3", linked_list: createLinkedListFromSlice([]int{7,7,7,7}), val: 7,	exp: createLinkedListFromSlice([]int{})},
	}

	for _, item := range test_items {
		t.Run(item.id, func(t *testing.T) {
			item.linked_list.removeElements02(item.val)

			if !reflect.DeepEqual(item.linked_list, item.exp) {
				t.Errorf("exp: %#v, got: %#v", item.linked_list, item.exp)
			}
		})
	}
}


func TestRemoveDuplicates(t *testing.T)  {
	testing_items := []struct{
		id string
		linked_list, exp *linkedList
	} {
		{id: "1", linked_list: createLinkedListFromSlice([]int{1, 1, 2}), exp: createLinkedListFromSlice([]int{1, 2})},
		{id: "2", linked_list: createLinkedListFromSlice([]int{1,1,2,3,3}), exp: createLinkedListFromSlice([]int{1, 2, 3})},
		{id: "3", linked_list: createLinkedListFromSlice([]int{1, 1, 1}), exp: createLinkedListFromSlice([]int{1})},
	}

	for _, item := range testing_items {
		t.Run(item.id, func(t *testing.T) {
			item.linked_list.removeDuplicates02()
			if !reflect.DeepEqual(item.linked_list, item.exp) {
				t.Errorf("got: %#v, exp: %#v", item.linked_list, item.exp)
			}
		})
	}
}

func TestMiddleNode(t *testing.T)  {
	testing_data := []struct {
		id string
		linked_list *linkedList
		exp int
	} {
		{id: "1", linked_list: createLinkedListFromSlice([]int {1, 2, 3, 4, 5}), exp: 3},
		{id: "2", linked_list: createLinkedListFromSlice([]int {1, 2, 3, 4, 5, 6}), exp: 4},
		{id: "3", linked_list: createLinkedListFromSlice([]int {1}), exp: 1},
	}

	for _, item := range testing_data {
		t.Run(item.id, func(t *testing.T) {
			got := item.linked_list.middleNode()
			if got != item.exp {
				t.Errorf("got: %d, exp: %d", got, item.exp)
			}

		})
	}
	
}

func TestCreateCyclicLinkedListFromSlice(t *testing.T)  {
	cyclic_linked_list := createCyclicLinkedListFromSlice([]int{3, 2, 0, -4}, 1)

	one_got := cyclic_linked_list.head.value 
	two_got := cyclic_linked_list.head.next.value 
	three_got := cyclic_linked_list.head.next.next.value 
	four_got := cyclic_linked_list.head.next.next.next.value
	five_got := cyclic_linked_list.head.next.next.next.next.value

	if (one_got != 3) {
		t.Errorf("one_got: %d", one_got)
	}
	if (two_got != 2) {
		t.Errorf("two_got: %d", two_got)
	}
	if (three_got != 0) {
		t.Errorf("three_got: %d", three_got)
	}
	if (four_got != -4) {
		t.Errorf("four_got: %d", four_got)
	}
	if (five_got != 2) {
		t.Errorf("five_got: %d", five_got)
	}
}

func TestHasCycle(t *testing.T)  {
	cyclic_linked_list_1 := createCyclicLinkedListFromSlice([]int {3,2,0,-4}, 1)
	cyclic_linked_list_2 := createCyclicLinkedListFromSlice([]int {1}, -1)
	got_1 := cyclic_linked_list_1.hasCycle02()
	got_2 := cyclic_linked_list_2.hasCycle02()

	if got_1 != true {
		t.Errorf("got: %t", got_1)
	}
	if got_2 != false {
		t.Errorf("got: %t", got_2)
	}
	
}



