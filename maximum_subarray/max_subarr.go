package maximum_subarray

func maxSubArray(nums []int) int {
	// TODO resolve
	return MaxSubArray(nums)
}

// -2, 1,  -3, 4, -1, 2, 1, -5, 4

func MaxSubArray(nums []int) int {
	maxSubArr := nums[0]
	currSum := 0

	for _, num := range nums {
		if currSum < 0 {
			currSum = 0
		}
		currSum += num
		if currSum > maxSubArr {
			maxSubArr = currSum
		}
	}
	return maxSubArr
}
