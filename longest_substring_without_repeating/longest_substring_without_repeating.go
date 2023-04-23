package longest_substring_without_repeating

func lengthOfLongestSubstring(s string) int {
	return LengthOfLongestSubstring(s)
}

func LengthOfLongestSubstring(s string) int {
	rs := []rune(s)
	occurrences := map[rune]int{}

	maxLength := 0
	currLen := 0
	for i := 0; i < len(rs); {
		r := rs[i]
		if pos, found := occurrences[r]; found {
			if currLen > maxLength {
				maxLength = currLen
			}
			currLen = 0
			i = pos + 1
			occurrences = map[rune]int{}
			// TODO should use a for to delete keys up to pos instead and reuse same map value
			continue
		}
		currLen++
		occurrences[r] = i
		i++
	}
	if currLen > maxLength {
		maxLength = currLen
	}

	return maxLength
}
