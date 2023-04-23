package longest_palindromic_substr

func longestPalindrome(s string) string {
	return LongestPalindrome(s)
}

func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	rs := []rune(s)
	if len(rs) == 1 {
		return s
	}
	if len(rs) == 2 {
		if rs[0] == rs[1] {
			return s
		} else {
			return string(rs[0])
		}
	}

	maxLength := 0
	maxLeft := 0
	for i := 1; i < len(rs)-1; i++ {
		currLen, left := expandPalindrome(i-1, i+1, rs, 1)
		if currLen > maxLength {
			maxLength = currLen
			maxLeft = left
		}
	}

	for i := 0; i < len(rs)-1; i++ {
		currLen, left := expandPalindrome(i, i+1, rs, 0)
		if currLen > maxLength {
			maxLength = currLen
			maxLeft = left
		}
	}

	return string(rs[maxLeft:(maxLeft + maxLength)])
}

// expandPalindrome returns the length of the palindrome and its leftmost index
func expandPalindrome(left int, right int, rs []rune, currLength int) (int, int) {
	for {
		if rs[left] != rs[right] {
			return currLength, left + 1
		}
		currLength = currLength + 2

		left = left - 1
		right = right + 1
		if left < 0 || right > len(rs)-1 {
			return currLength, left + 1
		}
	}
}
