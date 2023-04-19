package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	return LongestCommonPrefix(strs)
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	rStrs := make([][]rune, len(strs))
	for i, str := range strs {
		rStrs[i] = []rune(str)
	}

	var res []rune
	stringsCount := len(rStrs)
	firstElemRuneCount := len(rStrs[0])
	for i := 0; i < firstElemRuneCount; i++ {
		currRune := rStrs[0][i]
		for j := 1; j < stringsCount; j++ {
			if len(rStrs[j]) <= i {
				return string(res)
			}
			if currRune != rStrs[j][i] {
				return string(res)
			}
		}
		res = append(res, currRune)
	}
	return string(res)
}
