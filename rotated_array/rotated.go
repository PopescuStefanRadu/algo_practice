package rotated_array

// TODO understand and implement better search, see leetcode submission and sort.Search

func Search(nums []int, target int) int {
	position, isRotated := FindRotationPosition(nums, 0, len(nums)-1)
	if !isRotated {
		search, _ := BinarySearch(nums, target, 0, len(nums)-1)
		return search
	}

	if target > nums[len(nums)-1] {
		search, _ := BinarySearch(nums, target, 0, position-1)
		return search
	}
	search, _ := BinarySearch(nums, target, position, len(nums)-1)
	return search
}

func FindRotationPosition(nums []int, left, right int) (int, bool) {
	if !checkIfRotated(nums, left, right) {
		return 0, false
	}

	if right-left <= 1 {
		if nums[left] < nums[right] {
			return left, true
		} else {
			return right, true
		}
	}

	mid := (left + right) / 2

	if pos, rot := FindRotationPosition(nums, left, mid); rot {
		return pos, true
	}
	if pos, rot := FindRotationPosition(nums, mid, right); rot {
		return pos, true
	}

	return (left + right) / 2, true
}

func checkIfRotated(nums []int, left, right int) bool {
	return nums[left] > nums[right]
}

func BinarySearch(nums []int, target, left, right int) (int, bool) {
	if target < nums[left] || target > nums[right] {
		return -1, false
	}
	if left == right {
		if nums[left] == target {
			return left, true
		}
		return -1, false
	}
	midPos := (left + right) / 2
	mid := nums[midPos]
	if target == mid {
		return midPos, true
	}
	if target < mid {
		return BinarySearch(nums, target, left, midPos)
	}
	return BinarySearch(nums, target, midPos+1, right)
}
