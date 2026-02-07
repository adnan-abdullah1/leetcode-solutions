func removeOuterParentheses(strs string) string {

	stack := []rune{}
	ans := []rune{}
    
	for _, v := range strs {
		if v == '(' {
			if len(stack) == 0 {
				stack = append(stack, v)
			} else {
				stack = append(stack, v)
				ans = append(ans, v)
			}
		} else {
			if len(stack) == 1 {
				stack = []rune{}
			} else {
				ans = append(ans, ')')
				stack = stack[0 : len(stack)-1]
			}
		}
	}

    return string(ans)
}