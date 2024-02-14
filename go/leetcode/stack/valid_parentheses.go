package stack

func IsValidParentheses(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	countMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, v := range s {

		if _, ok := countMap[v]; ok {
			stack = append(stack, v)

		} else if len(stack) == 0 || countMap[stack[len(stack)-1]] != v {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0

}
