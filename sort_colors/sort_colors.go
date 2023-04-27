package sort_colors

func sortColors(nums []int) {
	incidence := make(map[int]int, 3)
	for _, v := range nums {
		incidence[v] = incidence[v] + 1
	}

	pos := 0
	for i := 0; i < 3; i++ {
		for v, found := incidence[i]; found && v != 0; v, found = incidence[i] {
			nums[pos] = i
			incidence[i] = incidence[i] - 1
			pos++
		}
	}
}
