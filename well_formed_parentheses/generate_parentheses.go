package well_formed_parentheses

func generateParenthesis(n int) []string {
	return GenerateParenthesis(n)
}

func GenerateParenthesis(n int) []string {
	combinations := make([]string, 0)
	GenerateParenthesisRec(&combinations, "", n, 0, 0)
	return combinations
}

func GenerateParenthesisRec(allCombinations *[]string, combination string, total, open, closed int) {
	if open == total && closed == total {
		*allCombinations = append(*allCombinations, combination)
		return
	}

	if open < total {
		GenerateParenthesisRec(allCombinations, combination+"(", total, open+1, closed)
	}

	if closed < open {
		GenerateParenthesisRec(allCombinations, combination+")", total, open, closed+1)
	}
}
