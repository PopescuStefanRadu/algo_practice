package jump_game

func canJump(nums []int) bool {
	return canJumpFrom(nums, 0, make(map[int]bool))
}

func canJumpFrom(nums []int, currentPos int, memory map[int]bool) bool {
	if currentPos >= len(nums)-1 {
		return true
	}
	if m, exists := memory[currentPos]; exists {
		return m
	}

	jumpDistance := nums[currentPos]
	if jumpDistance == 0 {
		memory[currentPos] = false
		return false
	}

	for jumpDistance > 0 {
		if canJumpFrom(nums, currentPos+jumpDistance, memory) {
			memory[currentPos] = true
			return true
		}
		jumpDistance--
	}
	memory[currentPos] = false
	return false
}
