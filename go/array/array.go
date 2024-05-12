package array

func ContainsDuplicate(nums []int) bool  {
	nums_map := make(map[int]int)

	for _, num := range nums {
		_, ok := nums_map[num]

		if ok {
			return true
		} 

		nums_map[num] = 1
	}

	return false
}

func isAllItemsOfArrayZero (arr [26]int) bool {
	for _, item := range arr {
		if item != 0 {
			return false
		}
	}
	return true
}

func ValidAnagram (s, t string) bool {
	if (len(s) != len(t)) {
		return false
	}

	s_arr := []rune(s)
	t_arr := []rune(t)
	var letters_arr [26]int


	for i, s_letter := range s_arr {
		letters_arr[s_letter - 'a'] += 1
		letters_arr[t_arr[i] - 'a'] -= 1
	}

	return isAllItemsOfArrayZero(letters_arr)

}

func ConcatenationOfArray (nums []int) []int {
	new_nums := make([]int, len(nums) * 2)

	for i, num := range nums {
		new_nums[i] = num
		new_nums[len(nums) + i] = num
	}

	return new_nums
}

func GreatestRightSide (arr []int) []int {
	output_arr := make([]int, len(arr))
	greatest_item := arr[len(arr) - 1]

	output_arr[len(arr) - 1] = -1 

	for i := len(arr) - 2; i >= 0; i-- {
		output_arr[i] = greatest_item
		num := arr[i]
		if (num > greatest_item) {
			greatest_item = num
		}
	}

	return output_arr

}

















