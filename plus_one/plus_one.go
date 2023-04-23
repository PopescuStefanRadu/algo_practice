package plus_one

func plusOne(digits []int) []int {
	return PlusOne(digits)
}

func PlusOne(digits []int) []int {
	remainder := 1
	for i := len(digits) - 1; remainder != 0 && i >= 0; i-- {
		sum := digits[i] + remainder
		if sum > 9 {
			digits[i] = sum % 10
		} else {
			digits[i] = sum
			remainder = 0
		}
	}
	if remainder != 0 {
		digits = append([]int{remainder}, digits...)
	}
	return digits
}
