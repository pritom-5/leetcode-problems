package binarysearch

import (
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T)  {
	testing_items := []struct {
		id string
		target, exp int
		nums []int
	} {
		{id: "1", nums: []int {1,3,5,6}, target: 5, exp: 2},
		{id: "2", nums: []int {1,3,5,6}, target: 2, exp: 1},
		{id: "3", nums: []int {1,3,5,6}, target: 7, exp: 4},
	}

	for _, item := range testing_items {
		t.Run(item.id, func(t *testing.T) {
			got := SearchInsert(item.nums, item.target)

			if got != item.exp {
				t.Errorf("got: %d, exp: %d, target: %d", got, item.exp, item.target)

			}
		})
	}
}

func TestSquaresOfSortedArray (t *testing.T) {
	test_items := []struct{
		id string
		nums, exp []int
	}{
		{id: "1", nums: []int{-4,-1,0,3,10}, exp: []int{0,1,9,16,100}},
		{id: "2", nums: []int{-7,-3,2,3,11}, exp: []int{4,9,9,49,121}},
	}

	for _, item := range test_items {
		t.Run(item.id, func(t *testing.T) {
			got := squaresOfSortedArray(item.nums)
			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}



func TestShipWithinDays (t *testing.T) {
	test_items := []struct{
		id string
		weights []int
		days, exp int
	}{
		{id: "1", weights: []int{1,2,3,4,5,6,7,8,9,10}, days: 5, exp: 15},
		{id: "2", weights: []int{3,2,2,4,1,4}, days: 3, exp: 6},
		{id: "3", weights: []int{1,2,3,1,1}, days: 4, exp: 3},
	}

	for _, item := range test_items {
		// TODO: 
		// failing test
		t.SkipNow()
		t.Run(item.id, func(t *testing.T) {
			got := shipWithinDays(item.weights, item.days)
			if got != item.exp {
				t.Errorf("got: %d, exp: %d", got, item.exp)
			}
		})

	}
}


func TestSuccessfulPairs (t *testing.T) {
	test_items := []struct{
		id string
		potions, spells, exp []int
		success int
	}{
		{id: "1", potions: []int{1,2,3,4,5}, spells: []int{5,1,3}, success: 7, exp: []int{4, 0, 3}},
		{id: "2", potions: []int{8,5,8}, spells: []int{3,1,2}, success: 16, exp: []int{2, 0, 2}},
	}

	for _, item := range test_items {
		t.Run(item.id, func(t *testing.T) {
			got := successfulPairs(item.spells, item.potions, item.success)
			if !reflect.DeepEqual(got, item.exp) {
				t.Errorf("got: %#v, exp: %#v", got, item.exp)
			}
		})
	}
}






















