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

func TestMergeLinkedList(t *testing.T)  {
	ll_1 := createLinkedListFromSlice([]int {1, 2, 4})
	ll_2 := createLinkedListFromSlice([]int {1, 3, 4})

	ll_3 := createLinkedListFromSlice([]int {1, 1, 2, 3, 4, 4})

	merged_linked_list := mergeLinkedList(ll_1, ll_2)

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


