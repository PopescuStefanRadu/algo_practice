package permutations

func permute(nums []int) [][]int {
	return Permute(nums)
}

func Permute(nums []int) [][]int {
	result := make([][]int, 0)
	return PermuteRec(nums, nil, &result)
}

func PermuteRec(nums []int, currentPermutation []int, result *[][]int) [][]int {
	if len(nums) == 0 {
		*result = append(*result, currentPermutation)
	}

	for i, v := range nums {
		newPermutation := make([]int, len(currentPermutation), len(currentPermutation)+1)
		copy(newPermutation, currentPermutation)
		newPermutation = append(newPermutation, v)
		newNums := make([]int, i)
		copy(newNums, nums[:i])
		newNums = append(newNums, nums[i+1:]...)
		PermuteRec(newNums, newPermutation, result)
	}
	return *result
}
