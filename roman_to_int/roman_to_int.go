package roman_to_int

var romanRuneToInt map[rune]int = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	return RomanToInt(s)
}

func RomanToInt(s string) int {

	s2 := make([]rune, 0)
	for _, val := range s {
		s2 = append(s2, val)
	}

	num := 0
	var prevDistinctLetterVal *int = nil
	for i := len(s2) - 1; i >= 0; i-- {
		val := s2[i]
		intEquivalent := romanRuneToInt[val]

		if prevDistinctLetterVal == nil {
			prevDistinctLetterVal = &intEquivalent
		}

		if intEquivalent < *prevDistinctLetterVal {
			num -= intEquivalent
		} else {
			num += intEquivalent
			prevDistinctLetterVal = &intEquivalent
		}
	}
	return num
}
