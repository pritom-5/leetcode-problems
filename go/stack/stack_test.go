package stack

import "testing"

func TestIsValid(t *testing.T) {

	testing_data := []struct {
		id  int
		s   string
		exp bool
	}{
		{id: 1, s: "()", exp: true},
		{id: 2, s: "()[]{}", exp: true},
		{id: 3, s: "(]", exp: false},
	}

	for _, item := range testing_data {
		got := IsValid(item.s)
		if got != item.exp {
			t.Errorf("s: %s, got: %t, exp: %t", item.s, got, item.exp)
		}
	}
}

func TestBaseballGame(t *testing.T) {

	testing_data := []struct {
		id, exp int
		ops     []string
	}{
		{id: 1, ops: []string{"5", "2", "C", "D", "+"}, exp: 30},
		{id: 2, ops: []string{"5", "-2", "4", "C", "D", "9", "+", "+"}, exp: 27},
		{id: 3, ops: []string{"1", "C"}, exp: 0},
	}

	for _, item := range testing_data {
		got := BaseballGame(item.ops)
		if got != item.exp {
			t.Errorf("got: %d, exp: %d", got, item.exp)
		}
	}
}
