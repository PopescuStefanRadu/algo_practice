package letter_combinations_phone_num

var letterMappings = map[rune][]string{
	'1': nil,
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	return LetterCombinations(digits)
}

func LetterCombinations(digits string) []string {
	rs := []rune(digits)

	if digits == "" {
		return nil
	}
	result := make([]string, 0)
	letterCombinationsRec(rs, "", &result)
	return result
}

func letterCombinationsRec(digits []rune, currentLetterCombination string, result *[]string) {
	if len(digits) == 0 {
		*result = append(*result, currentLetterCombination)
		return
	}

	currentDigit := digits[0]
	currentLettersToAdd := letterMappings[currentDigit]
	remainingDigitsToConsume := digits[1:]
	if currentLettersToAdd == nil {
		letterCombinationsRec(remainingDigitsToConsume, currentLetterCombination, result)
		return
	}

	for _, r := range currentLettersToAdd {
		letterCombinationsRec(remainingDigitsToConsume, currentLetterCombination+r, result)
	}
}
