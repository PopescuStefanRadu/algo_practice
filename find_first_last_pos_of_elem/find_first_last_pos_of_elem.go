package find_first_last_pos_of_elem

import "sort"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	pos := sort.SearchInts(nums, target)

	if pos >= len(nums) || nums[pos] != target {
		return []int{-1, -1}
	}

	pos2 := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})

	if pos2 != 0 {
		pos2--
	}

	return []int{pos, pos2}
}
