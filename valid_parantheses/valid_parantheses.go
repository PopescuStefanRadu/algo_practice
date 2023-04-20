package valid_parantheses

var reverseParenthesis = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

func isValid(s string) bool {
	return IsValidParentheses(s)
}

func IsValidParentheses(s string) bool {
	runes := []rune(s)

	var stack []rune
	for _, r := range runes {
		if r == '{' || r == '[' || r == '(' {
			stack = append(stack, r)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		popped := stack[len(stack)-1]
		if reverseParenthesis[popped] != r {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
