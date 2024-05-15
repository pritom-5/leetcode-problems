package stack

import (
	"fmt"
	"strconv"
)

func IsValid(s string) bool {

	samples := map[string]struct{}{
		"()": {},
		"[]": {},
		"{}": {},
	}

	s_arr := []rune(s)

	paren_stack := make([]rune, 0)

	for i := len(s_arr) - 1; i >= 0; i-- {
		curr := s_arr[i]

		if len(paren_stack) == 0 {
			paren_stack = append(paren_stack, curr)
			continue
		}

		last_popped_paren := paren_stack[len(paren_stack)-1]
		tmp_combined_paren := string(curr) + string(last_popped_paren)

		if _, ok := samples[tmp_combined_paren]; ok {
			paren_stack = paren_stack[:len(paren_stack)-1]
			continue
		}

		paren_stack = append(paren_stack, curr)

		fmt.Printf("paren_stack: %#v\n", paren_stack)
	}

	return len(paren_stack) == 0
}

func getSumOfArray(arr []string) int {
	total := 0

	for _, item := range arr {
		item_int, err := strconv.Atoi(item)
		if err != nil {
			fmt.Printf("item: %s", item)
			continue
		}
		total += item_int
	}
	return total
}

func BaseballGame(ops []string) int {
	stack := make([]string, 0)

	for _, op := range ops {
		switch op {
		case "C":
			stack = stack[:len(stack)-1]

		case "D":
			last_int_item, err := strconv.Atoi(stack[len(stack)-1])
			if err != nil {
				continue
			}
			doubled := last_int_item * 2
			stack = append(stack, strconv.Itoa(doubled))

		case "+":
			last_int_item, err := strconv.Atoi(stack[len(stack)-1])
			if err != nil {
				continue
			}
			second_last_int_item, err := strconv.Atoi(stack[len(stack)-2])
			if err != nil {
				continue
			}
			stack = append(stack, strconv.Itoa(last_int_item+second_last_int_item))

		default:
			stack = append(stack, op)
		}
	}

	return getSumOfArray(stack)
}
