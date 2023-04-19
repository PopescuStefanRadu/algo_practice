package two_sum

func twoSum(nums []int, target int) []int {
	return TwoSum(nums, target)
}

func TwoSum(nums []int, target int) []int {
	valueToPos := make(map[int]int, len(nums))
	for i, val := range nums {
		valueToPos[val] = i
	}

	for pos, val := range nums {
		requiredSummand := target - val
		secondOperandPos, found := valueToPos[requiredSummand]
		if secondOperandPos == pos {
			continue
		}
		if found {
			return []int{pos, secondOperandPos}
		}
	}
	return nil
}
